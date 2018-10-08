package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Serve runs the master API server, blocking until a termination
// signal is received
func Serve() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/testmodels", func(c chi.Router) {
		c.Get("/", handleTestModelGetAll)
		c.Post("/", handleTestModel)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
