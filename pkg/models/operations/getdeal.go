package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetDealQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetDealRequest struct {
	QueryParams GetDealQueryParams
}

type GetDeal200ApplicationJSON struct {
	Deal *shared.Deal `json:"deal,omitempty"`
}

type GetDealResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	GetDeal200ApplicationJSONObject *GetDeal200ApplicationJSON
}
