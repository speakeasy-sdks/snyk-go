package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetAccountDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetAccountDetailsRequest struct {
	QueryParams GetAccountDetailsQueryParams
}

type GetAccountDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetAccountDetailsResponse struct {
	ContentType                               string
	StatusCode                                int
	GetAccountDetails200ApplicationJSONObject *GetAccountDetails200ApplicationJSON
}
