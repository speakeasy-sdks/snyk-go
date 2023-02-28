package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type PostLinkTokenSecurity struct {
	VesselAPIToken shared.SchemeVesselAPIToken `security:"scheme,type=apiKey,subtype=header"`
}

type PostLinkTokenRequest struct {
	Security PostLinkTokenSecurity
}

type PostLinkToken200ApplicationJSON struct {
	LinkToken *string `json:"linkToken,omitempty"`
}

type PostLinkTokenResponse struct {
	ContentType                           string
	StatusCode                            int
	PostLinkToken200ApplicationJSONObject *PostLinkToken200ApplicationJSON
}
