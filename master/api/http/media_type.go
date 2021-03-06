package http

import (
	"net/http"

	"github.com/briand787b/rfs/core/models"

	"github.com/pkg/errors"
)

// MediaTypeRequest is the request payload for Article data model.
//
// NOTE: It's good practice to have well defined request and response payloads
// so you can manage the specific inputs and outputs for clients, and also gives
// you the opportunity to transform data on input or output, for example
// on request, we'd like to protect certain fields and on output perhaps
// we'd like to include a computed field based on other values that aren't
// in the data model. Also, check out this awesome blog post on struct composition:
// http://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
type MediaTypeRequest struct {
	*models.MediaType
}

func (mtr *MediaTypeRequest) Bind(r *http.Request) error {
	// a.Article is nil if no Article fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if mtr.MediaType == nil {
		return errors.New("missing required MediaTypeRequest fields")
	}

	// a.Subfield is nil if no Userpayload fields are sent in the request. In this app
	// this won't cause a panic, but checks in this Bind method may be required if
	// a.Subfield or futher nested fields like a.Subfield.Name are accessed elsewhere.

	// just a post-process after a decode..
	// a.ProtectedID = ""                                 // unset the protected ID
	// a.Article.Title = strings.ToLower(a.Article.Title) // as an example, we down-case
	return nil
}

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
	NextSkip int `json:"next_skip,omitempty"`
}

func NewMediaTypeResponseList(mts []models.MediaType, skip, take int) *MediaTypeResponseList {
	return &MediaTypeResponseList{
		MediaTypes: mts,
		Skip:       skip,
		Take:       take,
	}
}

func (mtrl *MediaTypeResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	if len(mtrl.MediaTypes) >= mtrl.Take {
		mtrl.NextSkip = mtrl.Skip + mtrl.Take
	}

	return nil
}
