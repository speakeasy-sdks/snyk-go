// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetTasksRequest struct {
	// The token for the customer's CRM account. This was generated when they connected their account.
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	// Returns all fields including non-unifiable and custom fields under the "additional" property in the response
	AllFields *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor    *string `queryParam:"style=form,explode=true,name=cursor"`
}

// GetTasks200ApplicationJSON - OK
type GetTasks200ApplicationJSON struct {
	NextPageCursor *string       `json:"nextPageCursor,omitempty"`
	Tasks          []shared.Task `json:"tasks,omitempty"`
}

type GetTasksResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetTasks200ApplicationJSONObject *GetTasks200ApplicationJSON
}
