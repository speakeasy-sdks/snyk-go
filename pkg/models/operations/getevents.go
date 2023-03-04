package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetEventsQueryParams struct {
	AccessToken *string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetEventsRequest struct {
	QueryParams GetEventsQueryParams
}

type GetEvents200ApplicationJSON struct {
	Events         []shared.Event `json:"events,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetEventsResponse struct {
	ContentType                       string
	StatusCode                        int
	RawResponse                       *http.Response
	GetEvents200ApplicationJSONObject *GetEvents200ApplicationJSON
}
