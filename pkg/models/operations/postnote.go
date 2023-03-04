package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostNoteRequestBody struct {
	AccessToken string            `json:"accessToken"`
	Note        shared.NoteCreate `json:"note"`
}

type PostNoteRequest struct {
	Request *PostNoteRequestBody `request:"mediaType=application/json"`
}

type PostNote200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostNoteResponse struct {
	ContentType                      string
	StatusCode                       int
	RawResponse                      *http.Response
	PostNote200ApplicationJSONObject *PostNote200ApplicationJSON
}
