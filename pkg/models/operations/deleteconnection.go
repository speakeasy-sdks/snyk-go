package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type DeleteConnectionRequestBody struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

type DeleteConnectionSecurity struct {
	VesselAPIToken shared.SchemeVesselAPIToken `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteConnectionRequest struct {
	Request  *DeleteConnectionRequestBody `request:"mediaType=application/json"`
	Security DeleteConnectionSecurity
}

type DeleteConnectionResponse struct {
	ContentType                              string
	StatusCode                               int
	DeleteConnection200ApplicationJSONString *string
}
