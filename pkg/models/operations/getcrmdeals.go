package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetCrmDealsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetCrmDealsRequest struct {
	QueryParams GetCrmDealsQueryParams
}

type GetCrmDeals200ApplicationJSON struct {
	Deals          []shared.Deal `json:"deals,omitempty"`
	NextPageCursor *string       `json:"nextPageCursor,omitempty"`
}

type GetCrmDealsResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	GetCrmDeals200ApplicationJSONObject *GetCrmDeals200ApplicationJSON
}
