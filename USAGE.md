<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/snyk-go"
    "github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
    "github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New()

    req := operations.DeleteConnectionRequest{
        Security: operations.DeleteConnectionSecurity{
            VesselAPIToken: shared.SchemeVesselAPIToken{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        Request: &operations.DeleteConnectionRequestBody{
            AccessToken: "unde",
        },
    }

    ctx := context.Background()
    res, err := s.DeleteConnection(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteConnection200ApplicationJSONString != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->