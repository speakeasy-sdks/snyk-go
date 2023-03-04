package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetUserQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetUserRequest struct {
	QueryParams GetUserQueryParams
}

type GetUser200ApplicationJSON struct {
	User *shared.User `json:"user,omitempty"`
}

type GetUserResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	GetUser200ApplicationJSONObject *GetUser200ApplicationJSON
}
