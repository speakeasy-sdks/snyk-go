package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetTaskQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetTaskRequest struct {
	QueryParams GetTaskQueryParams
}

type GetTask200ApplicationJSON struct {
	Task *shared.Task `json:"task,omitempty"`
}

type GetTaskResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	GetTask200ApplicationJSONObject *GetTask200ApplicationJSON
}
