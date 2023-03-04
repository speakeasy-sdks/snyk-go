package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PutTaskRequestBody struct {
	AccessToken string             `json:"accessToken"`
	ID          string             `json:"id"`
	Task        *shared.TaskUpdate `json:"task,omitempty"`
}

type PutTaskRequest struct {
	Request *PutTaskRequestBody `request:"mediaType=application/json"`
}

type PutTask200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutTaskResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	PutTask200ApplicationJSONObject *PutTask200ApplicationJSON
}
