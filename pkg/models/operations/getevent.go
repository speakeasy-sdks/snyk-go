// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetEventRequest struct {
	// The token for the customer's CRM account. This was generated when they connected their account.
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	// Returns all fields including non-unifiable and custom fields under the "additional" property in the response
	AllFields *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID        string `queryParam:"style=form,explode=true,name=id"`
}

// GetEvent200ApplicationJSON - OK
type GetEvent200ApplicationJSON struct {
	// Events are a type of activity that has attendees and takes place at a certain point in time (i.e., has a start and end date).
	//
	// For the currently supported CRMs, these are the objects Events most closely maps to:
	// - Salesforce = Events
	// - HubSpot = Meetings
	// - Pipedrive = Activities
	Event *shared.Event `json:"event,omitempty"`
}

type GetEventResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetEvent200ApplicationJSONObject *GetEvent200ApplicationJSON
}
