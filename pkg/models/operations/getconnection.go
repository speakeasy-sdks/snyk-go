package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetConnectionQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
}

type GetConnectionRequest struct {
	QueryParams GetConnectionQueryParams
}

type GetConnection200ApplicationJSON struct {
	Connection *shared.Connection `json:"connection,omitempty"`
}

type GetConnectionResponse struct {
	ContentType                           string
	StatusCode                            int
	GetConnection200ApplicationJSONObject *GetConnection200ApplicationJSON
}
