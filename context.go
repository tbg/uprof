package uprof

import "context"

type ctxKey struct{}

func WithProfile(ctx context.Context, p *Profile) context.Context {
	return context.WithValue(ctx, ctxKey{}, p)
}

func FromContext(ctx context.Context) *Profile {
	i := ctx.Value(ctxKey{})
	v, _ := i.(*Profile)
	return v
}
