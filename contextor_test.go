package contextor_test

import (
	"context"
	"testing"

	"github.com/metrumresearchgroup/wrapt"

	"github.com/metrumresearchgroup/contextor"
)

func TestRoundTrip(tt *testing.T) {
	t := wrapt.WrapT(tt)

	ctx := context.Background()

	type value string

	v := value("string")

	sr := contextor.New[value]("label")

	ctx, err := sr.Set(ctx, v)
	t.R.NoError(err)

	var v2 value

	v2, err = sr.Get(ctx)
	t.R.NoError(err)

	t.R.Equal(v2, v)
}

func TestProvidedKeyRoundTrip(tt *testing.T) {
	t := wrapt.WrapT(tt)

	ctx := context.Background()

	type value string

	v := value("string")

	type externalKeyType string // theoretically in another packge
	var PublicKeyValue externalKeyType = "public-key-value"
	sr := contextor.NewProvidedKey[value](PublicKeyValue)

	ctx, err := sr.Set(ctx, v)
	t.R.NoError(err)

	var v2 value

	v2, err = sr.Get(ctx)
	t.R.NoError(err)

	t.R.Equal(v2, v)
}
