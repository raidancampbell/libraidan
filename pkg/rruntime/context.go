package rruntime

import (
	"bytes"
	"context"
	"encoding/gob"
	"reflect"
	"time"
	"unsafe"
)

var (
	cancCtxTyp reflect.Type
	timeCtxTyp reflect.Type
)

func init() {
	// these are used as constants for type comparison.
	// unfortunately we need to access them via reflection, so an init function is required
	cancCtx, c := context.WithCancel(context.Background())
	c() // immediately cancel to prevent resource leaks
	cancCtxTyp = reflect.ValueOf(cancCtx).Elem().Type()
	timeCtx, c := context.WithDeadline(context.Background(), time.Time{})
	c()
	timeCtxTyp = reflect.ValueOf(timeCtx).Elem().Type()
}

// SerializeOpts defines options used during serialization
type SerializeOpts struct {
	// RetainCancel indicates whether the given context's cancel function (if any)
	// should be retained and re-inflated during deserialization.
	RetainCancel bool
	// RetainDeadline indicates whether the given context's deadline (if any)
	// should be retained and re-inflated during deserialization.
	RetainDeadline bool
	// IgnoreFunctions indicates whether functions stored in the given context's values
	// should be ignored.  Unignored functions will cause an error during serialization
	IgnoreFunctions bool
}

// Serialize serializes a given context.  See SerializeCtx
// deprecated: use SerializeCtx
var Serialize = SerializeCtx

// Deserialize deserializes the given output of Serialize.  See DeserializeCtx
// deprecated: use DeserializeCtx
var Deserialize = DeserializeCtx

// SerializeCtx serializes a given context's values, deadlines, and cancellations into a byte array
// If any function has been added to the context as a value, an error will be returned
// duplicate context keys aren't supported: the context stack is being frozen at the point of serialization.
// the deepest (furthest from base context) key is used.
// an optional parameter is used to specify serialization options.
// omitted options retain cancel/deadlines, but will not ignore functions
func SerializeCtx(ctx context.Context, opts ...SerializeOpts) ([]byte, error) {
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)

	s := contextData{
		Values:    make(map[interface{}]interface{}),
		HasCancel: false,
		Deadline:  time.Time{},
	}

	serialized := buildMap(ctx, s)

	// if options were passed
	if len(opts) > 0 {
		// override cancel/deadline
		if !opts[0].RetainCancel {
			serialized.HasCancel = false
		}
		if !opts[0].RetainDeadline {
			serialized.HasDeadline = false
		}
		// ignore functions to allow serialization to pass
		if opts[0].IgnoreFunctions {
			for key, val := range serialized.Values {
				if reflect.TypeOf(key).Kind() == reflect.Func || reflect.TypeOf(val).Kind() == reflect.Func {
					delete(serialized.Values, key)
				}
			}
		}
	}

	// Encoding the map
	err := e.Encode(serialized)
	return buf.Bytes(), err
}

// DeserializeCtx inflates the byte-array output of SerializeCtx into a context and optional CancelFunc
// The options specified during serialization dictate whether CancelFunc is non-nil
func DeserializeCtx(ser []byte) (context.Context, context.CancelFunc, error) {
	dec := gob.NewDecoder(bytes.NewReader(ser))
	data := contextData{}
	err := dec.Decode(&data)
	if err != nil {
		return context.Background(), func() {}, err
	}

	// make a new base context
	ctx := context.Background()

	// get back the values
	for key, val := range data.Values {
		ctx = context.WithValue(ctx, key, val)
	}

	// get back the cancel
	var c context.CancelFunc
	if data.HasCancel {
		ctx, c = context.WithCancel(ctx)
	}

	// get back the deadline
	if data.HasDeadline {
		ctx, c = context.WithDeadline(ctx, data.Deadline)
	}

	return ctx, c, nil
}

type contextData struct {
	Values      map[interface{}]interface{}
	HasCancel   bool      // we'll just make a cancel func on the other side
	HasDeadline bool      // time.Time's default value is in-band, so we need a flag to state whether it's valid
	Deadline    time.Time // only useful if HasDeadline is set
}

func buildMap(ctx context.Context, s contextData) contextData {
	rs := reflect.ValueOf(ctx).Elem()
	if rs.Type() == reflect.ValueOf(context.Background()).Elem().Type() {
		// base case: if the current context is an emptyCtx, we're done.
		return s
	}

	rf := rs.FieldByName("key")
	if rf.IsValid() { // if there's a key, it's a valueCtx
		// make the key field read+write
		rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()

		if rf.CanInterface() { // panic-protection
			key := rf.Interface()

			// grab the val field, make it read+write
			rv := rs.FieldByName("val")
			rv = reflect.NewAt(rv.Type(), unsafe.Pointer(rv.UnsafeAddr())).Elem()
			if rv.CanInterface() {
				val := rv.Interface()
				// only add the key if it doesn't exist.  nested contexts can have the same keys
				// but the concept is lost after serialization: you can't drop things off the stack
				// to the same layer as pre-serialization
				// we're recursing up the stack, so the first instance of the key is the one we want
				if _, exists := s.Values[key]; !exists {
					s.Values[key] = val
					// register them for serialization
					gob.Register(key)
					gob.Register(val)
				}
			}
		}
	} else { // it's either a cancelCtx or timerCtx
		if rs.Type() == cancCtxTyp {
			s.HasCancel = true
		}
		if rs.Type() == timeCtxTyp {
			// if there's multiple deadlines in a context, choose the earliest
			deadline := rs.FieldByName("deadline")
			deadline = reflect.NewAt(deadline.Type(), unsafe.Pointer(deadline.UnsafeAddr())).Elem()
			deadlineTime := deadline.Convert(reflect.TypeOf(time.Time{})).Interface().(time.Time)
			if s.HasDeadline && deadlineTime.Before(s.Deadline) {
				s.Deadline = deadlineTime
			} else {
				s.HasDeadline = true
				s.Deadline = deadlineTime
			}
		}

	}

	parent := rs.FieldByName("Context")
	if parent.IsValid() && !parent.IsNil() {
		// if there's a parent context, recurse
		return buildMap(parent.Interface().(context.Context), s)
	}
	// not possible, but the compiler requires it.
	// the parent context would be empty, and is caught in the beginning
	return s
}
