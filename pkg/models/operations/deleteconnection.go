// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteConnectionSecurity struct {
	VesselAPIToken string `security:"scheme,type=apiKey,subtype=header,name=vessel-api-token"`
}

type DeleteConnectionRequestBody struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

type DeleteConnectionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DeleteConnection200ApplicationJSONString *string
}
