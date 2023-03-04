package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetTasksQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetTasksRequest struct {
	QueryParams GetTasksQueryParams
}

type GetTasks200ApplicationJSON struct {
	NextPageCursor *string       `json:"nextPageCursor,omitempty"`
	Tasks          []shared.Task `json:"tasks,omitempty"`
}

type GetTasksResponse struct {
	ContentType                      string
	StatusCode                       int
	RawResponse                      *http.Response
	GetTasks200ApplicationJSONObject *GetTasks200ApplicationJSON
}
