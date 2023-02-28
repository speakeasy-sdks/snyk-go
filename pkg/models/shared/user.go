package shared

import (
	"time"
)

type UserAssociations struct {
	EventIds []string `json:"eventIds"`
	NoteIds  []string `json:"noteIds"`
	TaskIds  []string `json:"taskIds"`
}

// User
// Users of the connected CRM platform. These are *not* the contacts or leads, but rather the user accounts from the connected CRM.
type User struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	Associations UserAssociations       `json:"associations"`
	CreatedTime  time.Time              `json:"createdTime"`
	Email        *string                `json:"email,omitempty"`
	FirstName    *string                `json:"firstName,omitempty"`
	ID           string                 `json:"id"`
	LastName     *string                `json:"lastName,omitempty"`
	ModifiedTime time.Time              `json:"modifiedTime"`
	NativeID     string                 `json:"nativeId"`
}
