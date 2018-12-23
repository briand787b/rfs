package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

type ctxKey int

const (
	skipCtxKey ctxKey = iota
	takeCtxKey
	mediaTypeCtxKey
)

// getSkip returns the value of the skip query string parameter
// It defaults to 0 if not provided. However, it does return an error if
// it exists but cannot be converted to an integer
func getSkip(r *http.Request) (int, error) {
	s := r.URL.Query().Get("skip")
	if s == "" {
		return 0, nil
	}

	sI, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Wrap(err, "could not convert skip parameter to int")
	}

	return sI, nil
}

// getTake returns the value of the take query string parameter
// It defaults to 100 if not provided. However, it does return an error if
// it exists but cannot be converted to an integer
func getTake(r *http.Request) (int, error) {
	t := r.URL.Query().Get("take")
	if t == "" {
		return 100, nil
	}

	tI, err := strconv.Atoi(t)
	if err != nil {
		return 0, errors.Wrap(err, "could not convert skip parameter to int")
	}

	return tI, nil
}

func skipTake(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, err := getSkip(r)
		if err != nil {
			render.Render(w, r, ErrNotFound) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		t, err := getTake(r)
		if err != nil {
			render.Render(w, r, ErrNotFound) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		ctx := context.WithValue(r.Context(), skipCtxKey, s)
		ctx = context.WithValue(ctx, takeCtxKey, t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
