package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type PostEventAttendee200ApplicationJSON struct {
	AccessToken   *string                     `json:"accessToken,omitempty"`
	EventAttendee *shared.EventAttendeeCreate `json:"eventAttendee,omitempty"`
}

type PostEventAttendeeResponse struct {
	ContentType                               string
	StatusCode                                int
	PostEventAttendee200ApplicationJSONObject *PostEventAttendee200ApplicationJSON
}
