package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetEventAttendeeQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetEventAttendeeRequest struct {
	QueryParams GetEventAttendeeQueryParams
}

type GetEventAttendee200ApplicationJSON struct {
	EventAttendee *shared.EventAttendee `json:"eventAttendee,omitempty"`
}

type GetEventAttendeeResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	GetEventAttendee200ApplicationJSONObject *GetEventAttendee200ApplicationJSON
}
