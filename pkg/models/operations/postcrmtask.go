// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostCrmTaskRequestBody struct {
	AccessToken string             `json:"accessToken"`
	Task        *shared.TaskCreate `json:"task,omitempty"`
}

// PostCrmTask200ApplicationJSON - OK
type PostCrmTask200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostCrmTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PostCrmTask200ApplicationJSONObject *PostCrmTask200ApplicationJSON
}
