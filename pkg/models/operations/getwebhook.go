package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetWebhookQueryParams struct {
	AccessToken *string `queryParam:"style=form,explode=true,name=accessToken"`
	WebhookID   *string `queryParam:"style=form,explode=true,name=webhookId"`
}

type GetWebhookRequest struct {
	QueryParams GetWebhookQueryParams
}

type GetWebhook200ApplicationJSON struct {
	Webhook *shared.WebhookMetadata `json:"webhook,omitempty"`
}

type GetWebhookResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	GetWebhook200ApplicationJSONObject *GetWebhook200ApplicationJSON
}
