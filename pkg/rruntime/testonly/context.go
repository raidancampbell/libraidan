package testonly

import "context"

type key struct{}

func AddKeyToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, key{}, val)
}

func ReadFromContext(ctx context.Context) interface{} {
	return ctx.Value(key{})
}
