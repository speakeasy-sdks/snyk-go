// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

// PatchEventAttendee200ApplicationJSON - OK
type PatchEventAttendee200ApplicationJSON struct {
	AccessToken   *string                     `json:"accessToken,omitempty"`
	EventAttendee *shared.EventAttendeeUpdate `json:"eventAttendee,omitempty"`
}

type PatchEventAttendeeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PatchEventAttendee200ApplicationJSONObject *PatchEventAttendee200ApplicationJSON
}
