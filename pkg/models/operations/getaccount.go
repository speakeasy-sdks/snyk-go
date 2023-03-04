package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetAccountQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetAccountRequest struct {
	QueryParams GetAccountQueryParams
}

type GetAccount200ApplicationJSON struct {
	Account *shared.Account `json:"account,omitempty"`
}

type GetAccountResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	GetAccount200ApplicationJSONObject *GetAccount200ApplicationJSON
}
