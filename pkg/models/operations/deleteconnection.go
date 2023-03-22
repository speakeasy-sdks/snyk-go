// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type DeleteConnectionSecurity struct {
	VesselAPIToken shared.SchemeVesselAPIToken `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteConnectionRequestBody struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

type DeleteConnectionRequest struct {
	Request  *DeleteConnectionRequestBody `request:"mediaType=application/json"`
	Security DeleteConnectionSecurity
}

type DeleteConnectionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteConnection200ApplicationJSONString *string
}
