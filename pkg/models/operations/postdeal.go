// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostDealRequestBody struct {
	AccessToken string            `json:"accessToken"`
	Deal        shared.DealCreate `json:"deal"`
}

// PostDeal200ApplicationJSON - OK
type PostDeal200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostDealResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PostDeal200ApplicationJSONObject *PostDeal200ApplicationJSON
}
