package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetIntegrations200ApplicationJSON struct {
	Integrations []shared.CrmIntegration `json:"integrations,omitempty"`
}

type GetIntegrationsResponse struct {
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
	GetIntegrations200ApplicationJSONObject *GetIntegrations200ApplicationJSON
}
