// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type EventAssociations struct {
	AccountIds       []string `json:"accountIds"`
	ContactIds       []string `json:"contactIds"`
	DealIds          []string `json:"dealIds"`
	EventAttendeeIds []string `json:"eventAttendeeIds"`
	LeadIds          []string `json:"leadIds"`
	// The Id of the User that owns this event
	OwnerUserID string `json:"ownerUserId"`
}

// Event - Events are a type of activity that has attendees and takes place at a certain point in time (i.e., has a start and end date).
//
// For the currently supported CRMs, these are the objects Events most closely maps to:
// - Salesforce = Events
// - HubSpot = Meetings
// - Pipedrive = Activities
type Event struct {
	// Returned when `allFields` is set in the query params. Includes all data, untransformed, we recieved from the downstream CRM
	Additional   map[string]interface{} `json:"additional,omitempty"`
	Associations EventAssociations      `json:"associations"`
	// The date that the event was created
	CreatedTime   time.Time  `json:"createdTime"`
	Description   *string    `json:"description,omitempty"`
	EndDateTime   *time.Time `json:"endDateTime,omitempty"`
	ID            string     `json:"id"`
	IsAllDayEvent *bool      `json:"isAllDayEvent,omitempty"`
	// Where this event takes place. Can be virtual
	Location *string `json:"location,omitempty"`
	// The date the event was last modified
	ModifiedTime time.Time `json:"modifiedTime"`
	// Id in the downstream CRM
	NativeID      string     `json:"nativeId"`
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	Subject       *string    `json:"subject,omitempty"`
	Type          *string    `json:"type,omitempty"`
}
