package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetNotesQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

type GetNotesRequest struct {
	QueryParams GetNotesQueryParams
}

type GetNotes200ApplicationJSON struct {
	NextPageCursor *string       `json:"nextPageCursor,omitempty"`
	Notes          []shared.Note `json:"notes,omitempty"`
}

type GetNotesResponse struct {
	ContentType                      string
	StatusCode                       int
	GetNotes200ApplicationJSONObject *GetNotes200ApplicationJSON
}
