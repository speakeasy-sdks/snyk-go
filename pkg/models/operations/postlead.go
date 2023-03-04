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

type PostLead200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostLeadResponse struct {
	ContentType                      string
	StatusCode                       int
	RawResponse                      *http.Response
	PostLead200ApplicationJSONObject *PostLead200ApplicationJSON
}
