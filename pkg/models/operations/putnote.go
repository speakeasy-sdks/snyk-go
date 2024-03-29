// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PutNoteRequestBody struct {
	AccessToken string `json:"accessToken"`
	ID          string `json:"id"`
	// Update an existing Note. Please note that updating associations is not currently supported.
	Note shared.NoteUpdate `json:"note"`
}

// PutNote200ApplicationJSON - OK
type PutNote200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutNoteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PutNote200ApplicationJSONObject *PutNote200ApplicationJSON
}
