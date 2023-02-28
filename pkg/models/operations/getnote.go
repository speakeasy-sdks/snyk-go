package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetNoteQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetNoteRequest struct {
	QueryParams GetNoteQueryParams
}

type GetNote200ApplicationJSON struct {
	Note *shared.Note `json:"note,omitempty"`
}

type GetNoteResponse struct {
	ContentType                     string
	StatusCode                      int
	GetNote200ApplicationJSONObject *GetNote200ApplicationJSON
}
