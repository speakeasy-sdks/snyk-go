package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetUserDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetUserDetailsRequest struct {
	QueryParams GetUserDetailsQueryParams
}

type GetUserDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetUserDetailsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetUserDetails200ApplicationJSONObject *GetUserDetails200ApplicationJSON
}
