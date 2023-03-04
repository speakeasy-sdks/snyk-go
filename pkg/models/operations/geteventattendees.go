package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetEventAttendeesQueryParams struct {
	AccessToken *string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetEventAttendeesRequest struct {
	QueryParams GetEventAttendeesQueryParams
}

type GetEventAttendees200ApplicationJSON struct {
	EventAttendees []shared.EventAttendee `json:"eventAttendees,omitempty"`
	NextPageCursor *string                `json:"nextPageCursor,omitempty"`
}

type GetEventAttendeesResponse struct {
	ContentType                               string
	StatusCode                                int
	RawResponse                               *http.Response
	GetEventAttendees200ApplicationJSONObject *GetEventAttendees200ApplicationJSON
}
