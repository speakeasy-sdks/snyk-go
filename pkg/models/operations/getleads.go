package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetLeadsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetLeadsRequest struct {
	QueryParams GetLeadsQueryParams
}

type GetLeads200ApplicationJSON struct {
	Leads          []shared.Lead `json:"leads,omitempty"`
	NextPageCursor *string       `json:"nextPageCursor,omitempty"`
}

type GetLeadsResponse struct {
	ContentType                      string
	StatusCode                       int
	GetLeads200ApplicationJSONObject *GetLeads200ApplicationJSON
}
