package nats

import "context"

type ctxKey string

const (
	ctxKeySubject ctxKey = "subject"
)

// WithSubject returns a new context with the subject set to the provided value.
// This subject will be used when sending or receiving messages and overrides the default.
func WithSubject(ctx context.Context, subject string) context.Context {
	return context.WithValue(ctx, ctxKeySubject, subject)
}
