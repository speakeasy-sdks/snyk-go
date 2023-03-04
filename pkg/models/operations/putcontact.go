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

type PutContactRequests struct {
	ApplicationXML []byte                     `request:"mediaType=application/xml"`
	Object         *PutContactApplicationJSON `request:"mediaType=application/json"`
}

type PutContactRequest struct {
	Request *PutContactRequests
}

type PutContact200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutContactResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	PutContact200ApplicationJSONObject *PutContact200ApplicationJSON
}
