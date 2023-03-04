package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostWebhookRequestBody struct {
	AccessToken *string                       `json:"accessToken,omitempty"`
	Webhook     *shared.WebhookMetadataCreate `json:"webhook,omitempty"`
}

type PostWebhookRequest struct {
	Request *PostWebhookRequestBody `request:"mediaType=application/json"`
}

type PostWebhook200ApplicationJSON struct {
	Webhook *shared.WebhookMetadata `json:"webhook,omitempty"`
}

type PostWebhookResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PostWebhook200ApplicationJSONObject *PostWebhook200ApplicationJSON
}
