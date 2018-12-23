package http

import (
	"net/http"
	"strconv"

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
