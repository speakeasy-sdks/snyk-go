package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetNoteDetailsQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetNoteDetailsRequest struct {
	QueryParams GetNoteDetailsQueryParams
}

type GetNoteDetails200ApplicationJSON struct {
	Fields         []shared.Field `json:"fields,omitempty"`
	NextPageCursor *string        `json:"nextPageCursor,omitempty"`
}

type GetNoteDetailsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetNoteDetails200ApplicationJSONObject *GetNoteDetails200ApplicationJSON
}
