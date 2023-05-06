package runsafe

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getCtx(t *testing.T) {
	recovered, err := receivesAndDrops(context.WithValue(context.Background(), "foo", "bar"))
	assert.Nil(t, err)
	if !assert.NotNil(t, recovered) {
		return
	}
	assert.Equal(t, "quux", recovered.Value("baz"))
}

func receivesAndDrops(ctx context.Context) (context.Context, error) {
	_ = ctx.Value("foo")
	return lacksAndCreates()
}

func lacksAndCreates() (context.Context, error) {
	return intermediate(context.WithValue(context.Background(), "baz", "quux"))
}

//go:noinline
func intermediate(ctx context.Context) (context.Context, error) {
	_ = ctx.Value("foo")
	return redHerring(-1, 42)
}

func redHerring(a, b int) (context.Context, error) {
	return RecoverCtx()
}
