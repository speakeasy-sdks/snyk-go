package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PatchEventAttendee200ApplicationJSON struct {
	AccessToken   *string                     `json:"accessToken,omitempty"`
	EventAttendee *shared.EventAttendeeUpdate `json:"eventAttendee,omitempty"`
}

type PatchEventAttendeeResponse struct {
	ContentType                                string
	StatusCode                                 int
	RawResponse                                *http.Response
	PatchEventAttendee200ApplicationJSONObject *PatchEventAttendee200ApplicationJSON
}
