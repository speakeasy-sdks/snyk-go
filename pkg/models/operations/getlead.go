package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetLeadQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetLeadRequest struct {
	QueryParams GetLeadQueryParams
}

type GetLead200ApplicationJSON struct {
	Lead *shared.Lead `json:"lead,omitempty"`
}

type GetLeadResponse struct {
	ContentType                     string
	StatusCode                      int
	GetLead200ApplicationJSONObject *GetLead200ApplicationJSON
}
