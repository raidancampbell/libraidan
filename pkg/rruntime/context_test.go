package rruntime

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSerializeContext(t *testing.T) {
	foobar := context.WithValue(context.Background(), "foo", "bar")
	assert.Equal(t, "bar", foobar.Value("foo"))
	serialized, err := SerializeCtx(foobar)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)
}

func TestDeserializeContext_SimpleRoundTrip(t *testing.T) {
	foobar := context.WithValue(context.Background(), "foo", "bar")
	assert.Equal(t, "bar", foobar.Value("foo"))
	serialized, err := SerializeCtx(foobar)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)

	ctx, _, _ := DeserializeCtx(serialized)
	assert.Equal(t, "bar", ctx.Value("foo"))
}

func TestDeserializeContext_SimpleRoundTripTodo(t *testing.T) {
	foobar := context.WithValue(context.TODO(), "foo", "bar")
	assert.Equal(t, "bar", foobar.Value("foo"))
	serialized, err := SerializeCtx(foobar)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)

	ctx, _, _ := DeserializeCtx(serialized)
	assert.Equal(t, "bar", ctx.Value("foo"))
}

func TestDeserializeContext_Simpleint(t *testing.T) {
	foobar := context.WithValue(context.Background(), 1, 2)
	assert.Equal(t, 2, foobar.Value(1))
	serialized, err := SerializeCtx(foobar)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)

	ctx, _, _ := DeserializeCtx(serialized)
	assert.Equal(t, 2, ctx.Value(1))
}

func TestDeserializeContext_SimpleNestedRoundTrip(t *testing.T) {
	foobar := context.WithValue(context.Background(), "foo", "bar")
	foobar = context.WithValue(foobar, "baz", "qux")
	assert.Equal(t, "bar", foobar.Value("foo"))
	serialized, err := SerializeCtx(foobar)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)

	ctx, _, _ := DeserializeCtx(serialized)
	assert.Equal(t, "bar", ctx.Value("foo"))
	assert.Equal(t, "qux", ctx.Value("baz"))
}

func TestDeserializeContext_SimpleDeepNestedRoundTrip(t *testing.T) {
	foobar := context.WithValue(context.Background(), "foo", "bar")
	foobar = context.WithValue(foobar, "baz", "qux")
	foobar = context.WithValue(foobar, "one", "two")
	assert.Equal(t, "bar", foobar.Value("foo"))
	serialized, err := SerializeCtx(foobar)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)

	ctx, _, _ := DeserializeCtx(serialized)
	assert.Equal(t, "bar", ctx.Value("foo"))
	assert.Equal(t, "qux", ctx.Value("baz"))
	assert.Equal(t, "two", ctx.Value("one"))
}

type k struct{}

func TestDeserializeContext_SimpleStruct(t *testing.T) {
	foobar := context.WithValue(context.Background(), k{}, "foo")

	serialized, err := SerializeCtx(foobar)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)
	ctx, _, _ := DeserializeCtx(serialized)
	assert.Equal(t, "foo", ctx.Value(k{}))
}

// unsupported, will fail but shouldn't panic
func TestSerialize_func(t *testing.T) {
	foobar := context.WithValue(context.Background(), k{}, func() string {
		return "foo"
	})
	assert.NotPanics(t, func() { _, _ = SerializeCtx(foobar) })
	_, err := SerializeCtx(foobar, SerializeOpts{IgnoreFunctions: false})
	assert.NotNil(t, err)

	assert.NotPanics(t, func() { _, _ = SerializeCtx(foobar) })
	serialized, err := SerializeCtx(foobar, SerializeOpts{IgnoreFunctions: true})
	assert.Nil(t, err)
	assert.NotNil(t, serialized)
}

func TestDeserialize_nestedSameKey(t *testing.T) {
	parentctx := context.WithValue(context.Background(), "name", "parent")
	childctx := context.WithValue(parentctx, "name", "child")
	assert.Equal(t, "child", childctx.Value("name"))
	assert.Equal(t, "parent", parentctx.Value("name"))

	serialized, err := SerializeCtx(childctx)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)
	ctx, _, _ := DeserializeCtx(serialized)

	assert.Equal(t, "child", ctx.Value("name"))
}

func TestDeserialize_nestedWithCancel(t *testing.T) {
	cancCtxP, canc := context.WithCancel(context.Background())
	parentctx := context.WithValue(cancCtxP, "name", "parent")
	childctx := context.WithValue(parentctx, "name", "child")
	cancCtxC, canc2 := context.WithCancel(childctx)
	assert.Equal(t, "child", cancCtxC.Value("name"))
	assert.Equal(t, "parent", parentctx.Value("name"))
	assert.Equal(t, nil, cancCtxP.Value("name"))
	assert.NotNil(t, canc)
	assert.NotNil(t, canc2)

	serialized, err := SerializeCtx(cancCtxC)
	assert.Nil(t, err)
	assert.NotNil(t, serialized)

	ctx, cancFunc, err := DeserializeCtx(serialized)
	assert.Nil(t, err)
	assert.NotNil(t, cancFunc)

	assert.Equal(t, "child", ctx.Value("name"))
}

func TestDeserialize_nestedWithDeadline(t *testing.T) {
	deadlineCtx, canc := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	parentctx := context.WithValue(deadlineCtx, "name", "parent")
	childctx := context.WithValue(parentctx, "name", "child")
	cancCtxC, canc2 := context.WithCancel(childctx)
	assert.Equal(t, "child", cancCtxC.Value("name"))
	assert.Equal(t, "parent", parentctx.Value("name"))
	assert.Equal(t, nil, deadlineCtx.Value("name"))
	assert.NotNil(t, canc)
	assert.NotNil(t, canc2)

	serialized, err := SerializeCtx(cancCtxC, SerializeOpts{
		RetainDeadline: true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, serialized)
	ctx, cancFunc, _ := DeserializeCtx(serialized)
	assert.NotNil(t, cancFunc)

	assert.Equal(t, "child", ctx.Value("name"))
	deadline, ok := ctx.Deadline()
	assert.True(t, ok)

	remainingMs := time.Until(deadline).Milliseconds()
	assert.LessOrEqual(t, remainingMs, int64(5000))
	assert.GreaterOrEqual(t, remainingMs, int64(1))
}
