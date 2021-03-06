package contextor

import (
	"context"
	"errors"
	"fmt"
)

var ErrWrongType = errors.New("wrong type")

// Contextor is the general interface to the context.
type Contextor[T any] interface {
	Get(ctx context.Context) (T, error)
	Set(context.Context, T) (context.Context, error)
}

// New creates a new Contextor for any type, using the label indicated
// to make it unique. (if necessary)
//
// You can leave the label empty if you only expect one copy of a particular
// type in the system.
func New[T any](label string) Contextor[T] {
	var empty T

	return &contextor[T, ctxKey]{key: ctxKey(fmt.Sprintf("%T-%s", empty, label))}
}

func NewProvidedKey[T any, K comparable](providedKey K) Contextor[T] {
	return &contextor[T, K]{key: providedKey}
}

// Set places a value on the context, creating a key from the label and
// type.
func (c *contextor[T, K]) Set(ctx context.Context, v T) (context.Context, error) {
	return context.WithValue(ctx, c.key, v), nil
}

// Get retrieves a value from context, creating the key on the fly
// to match the label and type.
func (c *contextor[T, K]) Get(ctx context.Context) (T, error) {
	var empty T
	var v T
	a := ctx.Value(c.key)
	v, ok := a.(T)

	if !ok {
		return empty, fmt.Errorf("target type %T expected; value from context was type %T: %w", v, a, ErrWrongType)
	}

	return v, nil
}

// inner types

type (
	ctxKey                         string
	contextor[T any, K comparable] struct{ key K }
)
