package http

import (
	"context"
	"net/http"

	"github.com/briand787b/rfs/core/rfslog"
	"github.com/go-chi/render"
)

type Middleware struct {
	l rfslog.Logger
}

func (m *Middleware) skipTake(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, err := getSkip(r)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(m.l, err)) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		t, err := getTake(r)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(m.l, err)) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		ctx := context.WithValue(r.Context(), skipCtxKey, s)
		ctx = context.WithValue(ctx, takeCtxKey, t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
