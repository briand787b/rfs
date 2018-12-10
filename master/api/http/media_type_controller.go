package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/briand787b/rfs/core/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/briand787b/rfs/core/models"
)

type mediaTypeController struct {
	l   log.Logger
	mts models.MediaTypeStore
}

func newMediaTypeController(mts models.MediaTypeStore, l log.Logger) *mediaTypeController {
	return &mediaTypeController{
		l:   l,
		mts: mts,
	}
}

func (mtc mediaTypeController) mediaTypeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mtID := chi.URLParam(r, "media_type_id")
		if mtID == "" {
			render.Render(w, r, ErrNotFound) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		mtIDInt, err := strconv.Atoi(mtID)
		if err != nil {
			render.Render(w, r, ErrNotFound) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		mt, err := mtc.mts.GetByID(mtIDInt)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), mediaTypeCtxKey, mt)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (mtc mediaTypeController) handleMediaTypeGetByID(w http.ResponseWriter, r *http.Request) {
	mt, ok := r.Context().Value(mediaTypeCtxKey).(*models.MediaType)
	if !ok {
		render.Render(w, r, ErrNotFound)
		return
	}

	render.Render(w, r, NewMediaTypeResponse(mt))
}
