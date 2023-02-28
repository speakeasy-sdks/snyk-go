package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

type GetContactQueryParams struct {
	AccessToken string  `queryParam:"style=form,explode=true,name=accessToken"`
	AllFields   *bool   `queryParam:"style=form,explode=true,name=allFields"`
	Email       *string `queryParam:"style=form,explode=true,name=email"`
	ID          *string `queryParam:"style=form,explode=true,name=id"`
}

type GetContactRequest struct {
	QueryParams GetContactQueryParams
}

type GetContact200ApplicationJSON struct {
	Contact *shared.Contact `json:"contact,omitempty"`
}

type GetContactResponse struct {
	ContentType                        string
	StatusCode                         int
	GetContact200ApplicationJSONObject *GetContact200ApplicationJSON
}
