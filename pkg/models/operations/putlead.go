package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PutLeadRequestBody struct {
	AccessToken *string            `json:"accessToken,omitempty"`
	Lead        *shared.LeadUpdate `json:"lead,omitempty"`
}

type PutLeadRequest struct {
	Request *PutLeadRequestBody `request:"mediaType=application/json"`
}

type PutLead200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutLeadResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	PutLead200ApplicationJSONObject *PutLead200ApplicationJSON
}
