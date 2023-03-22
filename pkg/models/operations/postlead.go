// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostLeadRequestBody struct {
	AccessToken string            `json:"accessToken"`
	Lead        shared.LeadCreate `json:"lead"`
}

type PostLeadRequest struct {
	Request *PostLeadRequestBody `request:"mediaType=application/json"`
}

// PostLead200ApplicationJSON - OK
type PostLead200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostLeadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PostLead200ApplicationJSONObject *PostLead200ApplicationJSON
}
