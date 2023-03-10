package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PutContactApplicationJSON struct {
	AccessToken string               `json:"accessToken"`
	Contact     shared.ContactUpdate `json:"contact"`
	ID          string               `json:"id"`
}

type PutContactJSONRequest struct {
	Request *PutContactApplicationJSON `request:"mediaType=application/json"`
}

type PutContactJSON200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutContactJSONResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	PutContactJSON200ApplicationJSONObject *PutContactJSON200ApplicationJSON
}
