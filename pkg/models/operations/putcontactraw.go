package operations

import (
	"net/http"
)

type PutContactRawRequest struct {
	Request []byte `request:"mediaType=application/xml"`
}

type PutContactRaw200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutContactRawResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	PutContactRaw200ApplicationJSONObject *PutContactRaw200ApplicationJSON
}
