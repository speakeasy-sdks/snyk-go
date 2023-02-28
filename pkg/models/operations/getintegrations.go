package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetIntegrations200ApplicationJSON struct {
	Integrations []shared.CrmIntegration `json:"integrations,omitempty"`
}

type GetIntegrationsResponse struct {
	ContentType                             string
	StatusCode                              int
	GetIntegrations200ApplicationJSONObject *GetIntegrations200ApplicationJSON
}
