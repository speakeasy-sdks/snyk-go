package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetLeadDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetLeadDetailsRequest struct {
	QueryParams GetLeadDetailsQueryParams
}

type GetLeadDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetLeadDetailsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetLeadDetails200ApplicationJSONObject *GetLeadDetails200ApplicationJSON
}
