package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type PutNoteRequestBody struct {
	AccessToken string            `json:"accessToken"`
	ID          string            `json:"id"`
	Note        shared.NoteUpdate `json:"note"`
}

type PutNoteRequest struct {
	Request *PutNoteRequestBody `request:"mediaType=application/json"`
}

type PutNote200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutNoteResponse struct {
	ContentType                     string
	StatusCode                      int
	PutNote200ApplicationJSONObject *PutNote200ApplicationJSON
}
