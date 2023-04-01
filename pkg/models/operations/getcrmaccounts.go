// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetCrmAccountsRequest struct {
	// The token for the customer's CRM account. This was generated when they connected their account.
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	// Returns all fields including non-unifiable and custom fields under the "additional" property in the response
	AllFields *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor    *string `queryParam:"style=form,explode=true,name=cursor"`
}

// GetCrmAccounts200ApplicationJSON - OK
type GetCrmAccounts200ApplicationJSON struct {
	Accounts       []shared.Account `json:"accounts,omitempty"`
	NextPageCursor *string          `json:"nextPageCursor,omitempty"`
}

type GetCrmAccountsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetCrmAccounts200ApplicationJSONObject *GetCrmAccounts200ApplicationJSON
}
