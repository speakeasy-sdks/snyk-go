package operations

import (
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
	"net/http"
)

type GetConnections200ApplicationJSON struct {
	Connections []shared.Connection `json:"connections,omitempty"`
}

type GetConnectionsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetConnections200ApplicationJSONObject *GetConnections200ApplicationJSON
}
