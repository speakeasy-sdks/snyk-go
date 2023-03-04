package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetCrmUsersQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetCrmUsersRequest struct {
	QueryParams GetCrmUsersQueryParams
}

type GetCrmUsers200ApplicationJSON struct {
	NextPageCursor *string       `json:"nextPageCursor,omitempty"`
	Users          []shared.User `json:"users,omitempty"`
}

type GetCrmUsersResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	GetCrmUsers200ApplicationJSONObject *GetCrmUsers200ApplicationJSON
}
