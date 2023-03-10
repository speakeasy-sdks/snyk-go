package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostEventRequestBody struct {
	AccessToken string              `json:"accessToken"`
	Event       *shared.EventCreate `json:"event,omitempty"`
}

type PostEventRequest struct {
	Request *PostEventRequestBody `request:"mediaType=application/json"`
}

type PostEvent200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostEventResponse struct {
	ContentType                       string
	StatusCode                        int
	RawResponse                       *http.Response
	PostEvent200ApplicationJSONObject *PostEvent200ApplicationJSON
}
