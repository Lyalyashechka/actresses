package context

import (
	"context"
	"fmt"
	"math/rand"
)

type requestIDKey struct{}

func GenerateRequestID() string {
	return fmt.Sprintf("%016x", rand.Int())
}

func GetOrGenerateRequestID(ctx context.Context) string {
	reqID := ctx.Value(requestIDKey{})
	if reqID == nil {
		return GenerateRequestID()
	}

	return reqID.(string)
}

func SetRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, reqID)
}
