<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New()

    ctx := context.Background()
    res, err := s.DeleteConnection(ctx, operations.DeleteConnectionRequestBody{
        AccessToken: snyk.String("corrupti"),
    }, operations.DeleteConnectionSecurity{
        VesselAPIToken: "YOUR_API_KEY_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteConnection200ApplicationJSONString != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->