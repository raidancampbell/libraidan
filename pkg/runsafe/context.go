// Package runsafe contains things that leverage the unsafe builtin package.  Its contents should be treated as experimental and unstable.
package runsafe

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

// (stuff)(package name).(function name)(type descriptor address, type value address(stuff)
var pattern = regexp.MustCompile(`^.+[a-zA-Z][a-zA-Z0-9\-_]*\.[a-zA-Z][a-zA-Z0-9\-_]*\((?P<type_itab>0x[0-9a-f]+), (?P<type_value>0x[0-9a-f]+).+`)

// RecoverCtx returns (from the bottom up) the first context that's encountered in the callstack
// it stops the current goroutine to build a stacktrace, walks up the stack to find contexts, and returns the first one
// if no context is encountered, an empty context and an UnrecoverableContext Error are returned
// this is not guaranteed to work: many things can go wrong, chief among them is that inlined functions elide their parameter memory addresses
// see https://dave.cheney.net/2019/12/08/dynamically-scoped-variables-in-go for a more thorough explanation on how this works
func RecoverCtx() (context.Context, error) {
	return emptyItab(context.Background())
}

//go:noinline
// type descriptors addresses are seemingly constant for a given goroutine
// We leverage this to identify parameter addresses that are a context.Context
// Specifically, we must build up each of the concrete context.Context implementations
// so that we have the full legal set of context descriptors.
func emptyItab(_ context.Context) (context.Context, error) {
	return valueItab(context.WithValue(context.Background(), "", ""))
}

//go:noinline
func valueItab(_ context.Context) (context.Context, error) {
	ctx, c := context.WithCancel(context.Background())
	defer c()
	return cancelItab(ctx)
}

//go:noinline
func cancelItab(_ context.Context) (context.Context, error) {
	ctx, c := context.WithDeadline(context.Background(), time.Now())
	defer c()
	return timerItab(ctx)
}

//go:noinline
func timerItab(_ context.Context) (context.Context, error) {
	return doGetCtx()
}

// doGetCtx contains the actual logic for context recovery.
// The inflated callstack is required to protect against false positives
func doGetCtx() (context.Context, error) {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false) // get the current callstack as a string
	sc := bufio.NewScanner(bytes.NewReader(buf[:n]))
	var (
		deadlineType, cancelType, valueType, emptyType uintptr // hold the type descriptor pointers for each of the context implementations
		stackMatch                                     int     // used to count our way up the stack, as the stack is constant the lowest few levels and we need to leverage that
	)
	for sc.Scan() { // for each line (walking up the stack from here)
		// if the line doesn't match, skip.
		matches := pattern.FindStringSubmatch(sc.Text())
		if matches == nil {
			continue
		}
		// if this is the first iteration, then it's just our function. skip it.
		if stackMatch == 0 && strings.Contains(sc.Text(), "doGetCtx") {
			continue
		}

		stackMatch++

		// grab the two memory addresses (itab and type value)
		var p1, p2 uintptr
		_, err1 := fmt.Sscanf(matches[1], "%v", &p1)
		_, err2 := fmt.Sscanf(matches[2], "%v", &p2)
		if err1 != nil || err2 != nil {
			continue
		}

		// build up the legal values for each implementation of context
		// the stackMatch must match the known location in the stack.
		// Otherwise we might return a malformed context
		if stackMatch == 1 && strings.Contains(sc.Text(), "timerItab") {
			deadlineType = p1
		} else if stackMatch == 2 && strings.Contains(sc.Text(), "cancelItab") {
			cancelType = p1
		} else if stackMatch == 3 && strings.Contains(sc.Text(), "valueItab") {
			valueType = p1
		} else if stackMatch == 4 && strings.Contains(sc.Text(), "emptyItab") {
			emptyType = p1
		} else if p1 != emptyType && p1 != valueType && p1 != cancelType && p1 != deadlineType {
			// if we're in the caller's code, and the first parameter isn't a known context implementation, then skip this stack frame
			continue
		}

		if stackMatch <= 4 { // we're still building the legal context implementations
			continue
		}

		// at this point we're done building the legal context implementations, and this matched one.
		// rebuild a context from the addresses, and return
		idata := [2]uintptr{p1, p2}
		return *(*context.Context)(unsafe.Pointer(&idata)), nil
	}
	// no context was found.  Return a non-nil context to be polite, but also return an error.
	return context.Background(), UnrecoverableContext{}
}

// UnrecoverableContext is an error indicating that the context could not be dynamically recovered
type UnrecoverableContext struct{}

func (UnrecoverableContext) Error() string {
	return "unable to recover context"
}
