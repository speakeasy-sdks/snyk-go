package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetContactDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetContactDetailsRequest struct {
	QueryParams GetContactDetailsQueryParams
}

type GetContactDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetContactDetailsResponse struct {
	ContentType                               string
	StatusCode                                int
	GetContactDetails200ApplicationJSONObject *GetContactDetails200ApplicationJSON
}
