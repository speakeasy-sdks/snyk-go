package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetEventAttendeeDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetEventAttendeeDetailsRequest struct {
	QueryParams GetEventAttendeeDetailsQueryParams
}

type GetEventAttendeeDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetEventAttendeeDetailsResponse struct {
	ContentType                                     string
	StatusCode                                      int
	GetEventAttendeeDetails200ApplicationJSONObject *GetEventAttendeeDetails200ApplicationJSON
}
