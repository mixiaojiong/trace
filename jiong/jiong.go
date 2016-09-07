package jiong

import (
	"context"
	"net/http"
)

type key string

const RequestIDKey key = "X-Request-ID"

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, RequestIDKey, req.Header.Get("X-Request-ID"))
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(RequestIDKey).(string)
}

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	ctx = newContextWithRequestID(ctx, req)
	h(ctx, rw, req)
}

type ContextAdapter struct {
	Ctx     context.Context
	Handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.Handler.ServeHTTPContext(ca.Ctx, rw, req)
}

func Handler(ctx context.Context, handler ContextHandlerFunc) *ContextAdapter {
	return &ContextAdapter{
		Ctx:     ctx,
		Handler: handler,
	}
}
