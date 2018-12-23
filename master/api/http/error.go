package http

import (
	"fmt"
	"net/http"

	"github.com/briand787b/rfs/core/rfslog"
	"github.com/pkg/errors"

	"github.com/go-chi/render"
)

// Error response payloads & renderers
var (
// ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
// ErrInvalidRequest = &ErrResponse{HTTPStatusCode: 400, StatusText: "Bad Request", l: rfslog.NewZapSugaredLogger()}
)

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging

	l      rfslog.Logger
	severe bool
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	if e.severe {
		fmt.Printf("ERROR IS SEVERE")
		e.l.Errorw("[SEVERE ERROR]",
			"error", e.Err,
		)
	} else {
		fmt.Println("ERROR IS NOT SEVERE")
		e.l.ShortError(e.Err)
	}

	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrNotFound(l rfslog.Logger) render.Renderer {
	return &ErrResponse{
		Err:            errors.New("Resource not found"),
		HTTPStatusCode: 404,
		StatusText:     "Resource not found",
		ErrorText:      "Resource not found",
		l:              l,
	}
}

func ErrInternalServer(l rfslog.Logger, err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error",
		ErrorText:      err.Error(),
		l:              l,
		severe:         true,
	}
}

func ErrInvalidRequest(l rfslog.Logger, err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
		l:              l,
	}
}

func ErrRender(l rfslog.Logger, err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
		l:              l,
	}
}
