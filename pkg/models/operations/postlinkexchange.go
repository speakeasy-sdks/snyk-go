// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostLinkExchangeSecurity struct {
	VesselAPIToken shared.SchemeVesselAPIToken `security:"scheme,type=apiKey,subtype=header"`
}

type PostLinkExchangeRequestBody struct {
	PublicToken string `json:"publicToken"`
}

type PostLinkExchangeRequest struct {
	Request  *PostLinkExchangeRequestBody `request:"mediaType=application/json"`
	Security PostLinkExchangeSecurity
}

type PostLinkExchange200ApplicationJSONIntegrationIDEnum string

const (
	PostLinkExchange200ApplicationJSONIntegrationIDEnumSalesforce PostLinkExchange200ApplicationJSONIntegrationIDEnum = "salesforce"
	PostLinkExchange200ApplicationJSONIntegrationIDEnumHubspot    PostLinkExchange200ApplicationJSONIntegrationIDEnum = "hubspot"
	PostLinkExchange200ApplicationJSONIntegrationIDEnumPipedrive  PostLinkExchange200ApplicationJSONIntegrationIDEnum = "pipedrive"
)

func (e *PostLinkExchange200ApplicationJSONIntegrationIDEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "salesforce":
		fallthrough
	case "hubspot":
		fallthrough
	case "pipedrive":
		*e = PostLinkExchange200ApplicationJSONIntegrationIDEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostLinkExchange200ApplicationJSONIntegrationIDEnum: %s", s)
	}
}

// PostLinkExchange200ApplicationJSON - Exchange for access token
type PostLinkExchange200ApplicationJSON struct {
	AccessToken   string                                              `json:"accessToken"`
	ConnectionID  string                                              `json:"connectionId"`
	IntegrationID PostLinkExchange200ApplicationJSONIntegrationIDEnum `json:"integrationId"`
	NativeOrgID   string                                              `json:"nativeOrgId"`
	NativeOrgURL  string                                              `json:"nativeOrgURL"`
}

type PostLinkExchangeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Exchange for access token
	PostLinkExchange200ApplicationJSONObject *PostLinkExchange200ApplicationJSON
}
