package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetTaskDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetTaskDetailsRequest struct {
	QueryParams GetTaskDetailsQueryParams
}

type GetTaskDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetTaskDetailsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetTaskDetails200ApplicationJSONObject *GetTaskDetails200ApplicationJSON
}
