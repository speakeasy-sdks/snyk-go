package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetCrmAccountsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetCrmAccountsRequest struct {
	QueryParams GetCrmAccountsQueryParams
}

type GetCrmAccounts200ApplicationJSON struct {
	Accounts       []shared.Account `json:"accounts,omitempty"`
	NextPageCursor *string          `json:"nextPageCursor,omitempty"`
}

type GetCrmAccountsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetCrmAccounts200ApplicationJSONObject *GetCrmAccounts200ApplicationJSON
}
