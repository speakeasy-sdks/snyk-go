package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PutEventRequestBody struct {
	AccessToken string              `json:"accessToken"`
	Event       *shared.EventUpdate `json:"event,omitempty"`
	ID          string              `json:"id"`
}

type PutEventRequest struct {
	Request *PutEventRequestBody `request:"mediaType=application/json"`
}

type PutEvent200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutEventResponse struct {
	ContentType                      string
	StatusCode                       int
	RawResponse                      *http.Response
	PutEvent200ApplicationJSONObject *PutEvent200ApplicationJSON
}
