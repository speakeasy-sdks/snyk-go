package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type PostLinkExchangeRequestBody struct {
	PublicToken string `json:"publicToken"`
}

type PostLinkExchangeSecurity struct {
	VesselAPIToken shared.SchemeVesselAPIToken `security:"scheme,type=apiKey,subtype=header"`
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

type PostLinkExchange200ApplicationJSON struct {
	AccessToken   string                                              `json:"accessToken"`
	ConnectionID  string                                              `json:"connectionId"`
	IntegrationID PostLinkExchange200ApplicationJSONIntegrationIDEnum `json:"integrationId"`
	NativeOrgID   string                                              `json:"nativeOrgId"`
	NativeOrgURL  string                                              `json:"nativeOrgURL"`
}

type PostLinkExchangeResponse struct {
	ContentType                              string
	StatusCode                               int
	PostLinkExchange200ApplicationJSONObject *PostLinkExchange200ApplicationJSON
}
