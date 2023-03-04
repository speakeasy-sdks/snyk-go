package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetContactsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetContactsRequest struct {
	QueryParams GetContactsQueryParams
}

type GetContacts200ApplicationJSON struct {
	Contacts       []shared.Contact `json:"contacts,omitempty"`
	NextPageCursor *string          `json:"nextPageCursor,omitempty"`
}

type GetContactsResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	GetContacts200ApplicationJSONObject *GetContacts200ApplicationJSON
}
