package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/briand787b/rfs/core/models"
)

type mediaTypeController struct {
	mts models.MediaTypeStore
}

func newMediaTypeController(mts models.MediaTypeStore) *mediaTypeController {
	return &mediaTypeController{
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
	fmt.Println("IN MEDIA TYPE CONTROLLER")
	fmt.Printf("MEDIA TYPE STORE: %v\n", mtc.mts)
	mt, ok := r.Context().Value(mediaTypeCtxKey).(*models.MediaType)
	if !ok {
		http.Error(w, "could not find media_type_id from url", http.StatusBadRequest)
		return
	}

	bs, err := json.Marshal(mt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error marshalling MediaType to JSON", http.StatusInternalServerError)
	}

	w.Write(bs)
}
