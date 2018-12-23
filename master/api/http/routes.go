package http

import (
	"net/http"
	"reflect"
	"runtime"

	"github.com/pkg/errors"

	"github.com/briand787b/rfs/core/models"
	"github.com/briand787b/rfs/core/postgres"
	"github.com/briand787b/rfs/core/rfslog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// Authentication notes: when using jwt, add a field called "x-session-id" to log sessions, must be unique per jwt creation

// Serve runs the master API server, blocking until a termination
// signal is received
func Serve(lvl string) error {
	l, err := rfslog.NewZapSugaredLogger(lvl)
	if err != nil {
		return errors.Wrap(err, "could not instantiate new zap sugared logger")
	}

	// initialize auth
	setSecret()

	// initialize databases
	db := postgres.GetExtFull(l)

	// initialize stores
	mts := models.NewMediaTypePGStore(l, db)

	// initialize controllers
	mtc := newMediaTypeController(mts, l)

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// v1.0 router
	r.Route("/v1.0", func(r chi.Router) {
		// unauthenticated routes
		r.Post("/login", handleLogin)

		r.Route("/media_types", func(r chi.Router) {
			r.With(skipTake).Get("/", mtc.handleGetAllMediaTypes)
			r.Route("/{media_type_id}", func(r chi.Router) {
				r.Use(mtc.mediaTypeCtx)
				r.Get("/", mtc.handleMediaTypeGetByID)
			})
		})

		// authenticated routes
		r.Group(func(r chi.Router) {
			// Seek, verify and validate JWT tokens
			r.Use(jwtauth.Verifier(tokenAuth))

			// Handle valid / invalid tokens. In this example, we use
			// the provided authenticator middleware, but you can write your
			// own very easily, look at the Authenticator method in jwtauth.go
			// and tweak it, its not scary.
			r.Use(jwtauth.Authenticator)
			r.Route("/testmodels", func(c chi.Router) {
				c.Get("/", handleTestModelGetAll)
				c.Post("/", handleTestModel)
			})
		})

	})

	// v2.0 router
	r.Route("/v2.0", func(r chi.Router) {

	})

	walkFn := func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		l.Infow("[API ROUTE]",
			"METHOD", method,
			"ROUTE", route,
			"HANDLER", funcName,
		)
		return nil
	}

	if err := chi.Walk(r, walkFn); err != nil {
		return errors.Wrap(err, "could not print API routes")
	}

	return http.ListenAndServe(":8080", r)
}
