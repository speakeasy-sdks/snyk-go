package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostCrmTaskRequestBody struct {
	AccessToken string             `json:"accessToken"`
	Task        *shared.TaskCreate `json:"task,omitempty"`
}

type PostCrmTaskRequest struct {
	Request *PostCrmTaskRequestBody `request:"mediaType=application/json"`
}

type PostCrmTask200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostCrmTaskResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PostCrmTask200ApplicationJSONObject *PostCrmTask200ApplicationJSON
}
