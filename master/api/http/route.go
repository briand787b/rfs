package http

import (
	"log"
	"net/http"
	"os"

	"github.com/briand787b/rfs/core/models"
	"github.com/briand787b/rfs/core/postgres"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Serve runs the master API server, blocking until a termination
// signal is received
func Serve() {
	// initialize databases
	db := postgres.GetExtFull(os.Stdout)

	// initialize stores
	mts := models.NewMediaTypePGStore(db)

	// initialize controllers
	mtc := newMediaTypeController(mts)

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/testmodels", func(c chi.Router) {
		c.Get("/", handleTestModelGetAll)
		c.Post("/", handleTestModel)
	})

	r.Route("/media_types", func(r chi.Router) {
		r.Route("/{media_type_id}", func(r chi.Router) {
			r.Use(mtc.mediaTypeCtx)
			r.Get("/", mtc.handleMediaTypeGetByID)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
