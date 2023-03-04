package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PostAccountRequestBody struct {
	AccessToken string               `json:"accessToken"`
	Account     shared.AccountCreate `json:"account"`
}

type PostAccountRequest struct {
	Request *PostAccountRequestBody `request:"mediaType=application/json"`
}

type PostAccount200ApplicationJSON struct {
	ID *string `json:"id,omitempty"`
}

type PostAccountResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PostAccount200ApplicationJSONObject *PostAccount200ApplicationJSON
}
