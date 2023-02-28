package shared

import (
	"time"
)

type WebhookMetadata struct {
	ConnectionID *string    `json:"connectionId,omitempty"`
	CreatedTime  *time.Time `json:"createdTime,omitempty"`
	WebhookID    *string    `json:"webhookId,omitempty"`
	WebhookURL   *string    `json:"webhookUrl,omitempty"`
}
