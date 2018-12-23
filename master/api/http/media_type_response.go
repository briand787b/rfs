package http

import (
	"net/http"

	"github.com/briand787b/rfs/core/models"
)

// ArticleResponse is the response payload for the Article data model.
// See NOTE above in ArticleRequest as well.
//
// In the ArticleResponse object, first a Render() is called on itself,
// then the next field, and so on, all the way down the tree.
// Render is called in top-down order, like a http handler middleware chain.
type MediaTypeResponse struct {
	*models.MediaType

	// User *UserPayload `json:"user,omitempty"` -- Figure this out later

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func NewMediaTypeResponse(mt *models.MediaType) *MediaTypeResponse {
	resp := &MediaTypeResponse{MediaType: mt}

	// if resp.User == nil {
	// 	if user, _ := dbGetUser(resp.UserID); user != nil {
	// 		resp.User = NewUserPayloadResponse(user)
	// 	}
	// }

	return resp
}

func (rd *MediaTypeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}

type MediaTypeResponseList struct {
	MediaTypes []models.MediaType `json:"media_types"`

	Skip     int `json:"skip"`
	Take     int `json:"take"`
	NextSkip int `json:"next_skip"`
}

func NewMediaTypeResponseList(mts []models.MediaType, skip, take int) *MediaTypeResponseList {
	return &MediaTypeResponseList{
		MediaTypes: mts,
		Skip:       skip,
		Take:       take,
	}
}

func (mtrl *MediaTypeResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	mtrl.NextSkip = mtrl.Skip + mtrl.Take
	return nil
}
