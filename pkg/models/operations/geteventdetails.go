package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetEventDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetEventDetailsRequest struct {
	QueryParams GetEventDetailsQueryParams
}

type GetEventDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetEventDetailsResponse struct {
	ContentType                             string
	StatusCode                              int
	GetEventDetails200ApplicationJSONObject *GetEventDetails200ApplicationJSON
}
