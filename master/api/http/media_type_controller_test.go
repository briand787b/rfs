package http

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/briand787b/rfs/core/rfslog"

	"github.com/briand787b/rfs/core/models"
)

func TestMediaTypeGetByID(t *testing.T) {
	tests := []struct {
		name     string
		ctxValue interface{}
		expCode  int
	}{
		{
			"existing media type is found and successfully retrieved",
			&models.MediaType{ID: 1, Name: "name"},
			200,
		},
	}

	for _, tt := range tests {
		tt := tt // prevent closure
		t.Run(tt.name, func(subT *testing.T) {
			// subT.Parallel()

			// mmts := &mockmodels.MediaTypeMockStore{GetByIDValue: tt.getValue, GetByIDError: tt.getError}
			l, _ := rfslog.NewZapSugaredLogger("debug")
			mtc := mediaTypeController{l: l}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/doesntmatter", nil)
			r = r.WithContext(context.WithValue(context.Background(), mediaTypeCtxKey, tt.ctxValue))
			mtc.handleGetByID(w, r)

			if w.Code != tt.expCode {
				subT.Fatalf("expected code to be %v, was %v\n", tt.expCode, w.Code)
			}
		})
	}

}
