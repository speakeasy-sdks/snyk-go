package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetDealDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
	ID          *string `queryParam:"style=form,explode=true,name=id"`
}

type GetDealDetailsRequest struct {
	QueryParams GetDealDetailsQueryParams
}

type GetDealDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetDealDetailsResponse struct {
	ContentType                            string
	StatusCode                             int
	GetDealDetails200ApplicationJSONObject *GetDealDetails200ApplicationJSON
}
