package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type PutDealRequestBody struct {
	AccessToken string            `json:"accessToken"`
	Deal        shared.DealUpdate `json:"deal"`
	ID          string            `json:"id"`
}

type PutDealRequest struct {
	Request *PutDealRequestBody `request:"mediaType=application/json"`
}

type PutDeal200ApplicationJSON struct {
	ID string `json:"id"`
}

type PutDealResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	PutDeal200ApplicationJSONObject *PutDeal200ApplicationJSON
}
