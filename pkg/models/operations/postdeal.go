package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type PostDealRequestBody struct {
	AccessToken string            `json:"accessToken"`
	Deal        shared.DealCreate `json:"deal"`
}

type PostDealRequest struct {
	Request *PostDealRequestBody `request:"mediaType=application/json"`
}

type PostDeal200ApplicationJSON struct {
	ID string `json:"id"`
}

type PostDealResponse struct {
	ContentType                      string
	StatusCode                       int
	PostDeal200ApplicationJSONObject *PostDeal200ApplicationJSON
}
