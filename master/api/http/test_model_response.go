package http

import (
	"net/http"

	"github.com/briand787b/rfs/core/models"
	"github.com/go-chi/render"
)

type TestModelResponse struct {
	*models.TestModel
}

func NewTestModelResponse(tm *models.TestModel) render.Renderer {
	// potentially do stuff here if modifications are needed
	return &TestModelResponse{tm}
}

func (tmr *TestModelResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type TestModelListResponse []*TestModelResponse

func NewTestModelListResponse(testModels []models.TestModel) []render.Renderer {
	list := []render.Renderer{}
	for i := 0; i < len(testModels); i++ {
		list = append(list, NewTestModelResponse(&testModels[i]))
	}
	return list
}
