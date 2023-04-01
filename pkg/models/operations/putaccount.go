// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PutAccountRequestBody struct {
	AccessToken string               `json:"accessToken"`
	Account     shared.AccountUpdate `json:"account"`
	ID          string               `json:"id"`
}

// PutAccount200ApplicationJSON - OK
type PutAccount200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutAccountResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PutAccount200ApplicationJSONObject *PutAccount200ApplicationJSON
}
