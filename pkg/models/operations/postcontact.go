package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostContactRequestBody struct {
	AccessToken string               `json:"accessToken"`
	Contact     shared.ContactCreate `json:"contact"`
}

type PostContactRequest struct {
	Request *PostContactRequestBody `request:"mediaType=application/json"`
}

type PostContact200ApplicationJSON struct {
	ID *string `json:"id,omitempty"`
}

type PostContactResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PostContact200ApplicationJSONObject *PostContact200ApplicationJSON
}
