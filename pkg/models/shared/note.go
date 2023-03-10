package shared

import (
	"time"
)

type NoteAssociations struct {
	AccountIds  []string `json:"accountIds"`
	ContactIds  []string `json:"contactIds"`
	DealIds     []string `json:"dealIds"`
	LeadIds     []string `json:"leadIds"`
	OwnerUserID string   `json:"ownerUserId"`
}

// Note
// A Note attached to some CRM Object.
type Note struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	Associations NoteAssociations       `json:"associations"`
	Content      *string                `json:"content,omitempty"`
	CreatedTime  time.Time              `json:"createdTime"`
	ID           string                 `json:"id"`
	ModifiedTime time.Time              `json:"modifiedTime"`
	NativeID     string                 `json:"nativeId"`
}
