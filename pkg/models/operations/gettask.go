// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetTaskQueryParams struct {
	// The token for the customer's CRM account. This was generated when they connected their account.
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	// Returns all fields including non-unifiable and custom fields under the "additional" property in the response
	AllFields *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID        string `queryParam:"style=form,explode=true,name=id"`
}

type GetTaskRequest struct {
	QueryParams GetTaskQueryParams
}

// GetTask200ApplicationJSON - OK
type GetTask200ApplicationJSON struct {
	// A task attached to some CRM object
	Task *shared.Task `json:"task,omitempty"`
}

type GetTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetTask200ApplicationJSONObject *GetTask200ApplicationJSON
}
