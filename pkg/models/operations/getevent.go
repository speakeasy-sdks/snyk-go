package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetEventQueryParams struct {
	AccessToken string `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool  `queryParam:"style=form,explode=true,name=allFields"`
	ID          string `queryParam:"style=form,explode=true,name=id"`
}

type GetEventRequest struct {
	QueryParams GetEventQueryParams
}

type GetEvent200ApplicationJSON struct {
	Event *shared.Event `json:"event,omitempty"`
}

type GetEventResponse struct {
	ContentType                      string
	StatusCode                       int
	GetEvent200ApplicationJSONObject *GetEvent200ApplicationJSON
}
