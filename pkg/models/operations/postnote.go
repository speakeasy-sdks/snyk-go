package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
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
	PostNote200ApplicationJSONObject *PostNote200ApplicationJSON
}
