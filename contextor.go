package contextor

import (
	"context"
	"fmt"
	"strings"
)

type (
	Contextor[T any, K comparable] struct {
		key K
	}
	ctxKey string
)

// New creates a new Contextor for any type, using the label indicated
// to make it unique. (if necessary)
//
// You can leave the label empty if you only expect one copy of a particular
// type in the system.
func New[T any](label string) Contextor[T, ctxKey] {
	vals := []string{fmt.Sprintf("%T", *new(T))}
	if len(label) > 0 {
		vals = append(vals, label)
	}
	key := strings.Join(vals, "-")

	return Contextor[T, ctxKey]{key: ctxKey(key)}
}

func NewProvidedKey[T any, K comparable](providedKey K) Contextor[T, K] {
	return Contextor[T, K]{key: providedKey}
}

// Set places a value on the context, creating a key from the label and
// type.
func (c Contextor[T, K]) Set(ctx context.Context, v T) (context.Context, error) {
	return context.WithValue(ctx, c.key, v), nil
}

// Get retrieves a value from context, creating the key on the fly
// to match the label and type.
func (c Contextor[T, K]) Get(ctx context.Context) (T, error) {
	raw := ctx.Value(c.key)
	value, ok := raw.(T)
	if ok {
		return value, nil
	}

	return *new(T), &ErrWrongType[T]{actual: raw, wanted: value}
}

type ErrWrongType[T any] struct {
	actual any
	wanted T
}

func (e ErrWrongType[T]) Error() string {
	return fmt.Sprintf("wrong type for key: %T; wanted: %T", e.actual, e.wanted)
}
