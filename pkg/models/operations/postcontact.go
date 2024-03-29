// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostContactRequestBody struct {
	AccessToken string `json:"accessToken"`
	// Properties that a contact can be created with
	Contact shared.ContactCreate `json:"contact"`
}

// PostContact200ApplicationJSON - OK
type PostContact200ApplicationJSON struct {
	ID *string `json:"id,omitempty"`
}

type PostContactResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PostContact200ApplicationJSONObject *PostContact200ApplicationJSON
}
