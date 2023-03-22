// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetEventAttendeeDetailsQueryParams struct {
	// The token for the customer's CRM account. This was generated when they connected their account.
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	// Returns details about native fields.
	AllFields *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor    *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetEventAttendeeDetailsRequest struct {
	QueryParams GetEventAttendeeDetailsQueryParams
}

// GetEventAttendeeDetails200ApplicationJSON - OK
type GetEventAttendeeDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetEventAttendeeDetailsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetEventAttendeeDetails200ApplicationJSONObject *GetEventAttendeeDetails200ApplicationJSON
}
