# Snyk SDK

## Overview

Vessel's CRM APIs requires an access token to be used together with your Vessel API token. Ensure the following headers are provided when making API calls:

Key | Value | Description
---------|----------|----------
 vessel-api-token | `<VESSEL_API_TOKEN>` | The API token provided by us
 
Additionally, in the query or body parameters of each request depending on whether it is a GET or POST, make sure to include the `accessToken` for the connection you want to access.

### Available Operations

* [DeleteConnection](#deleteconnection) - Delete connection
* [DeleteWebhook](#deletewebhook) - Remove Webhook
* [GetAccount](#getaccount) - Get Account
* [GetAccountDetails](#getaccountdetails) - Get Account Details
* [GetConnection](#getconnection) - Get Connection
* [GetConnections](#getconnections) - Get All Connections
* [GetContact](#getcontact) - Get Contact
* [GetContactDetails](#getcontactdetails) - Get Contact Details
* [GetContacts](#getcontacts) - Get All Contacts
* [GetCrmAccounts](#getcrmaccounts) - Get All Accounts
* [GetCrmDeals](#getcrmdeals) - Get All Deals
* [GetCrmUsers](#getcrmusers) - Get All Users
* [GetDeal](#getdeal) - Get Deal
* [GetDealDetails](#getdealdetails) - Get Deal Details
* [GetEvent](#getevent) - Get Event
* [GetEventAttendee](#geteventattendee) - Get Event Attendee
* [GetEventAttendeeDetails](#geteventattendeedetails) - Get Event Attendee Details
* [GetEventAttendees](#geteventattendees) - Get All Event Attendees
* [GetEventDetails](#geteventdetails) - Get Event Details
* [GetEvents](#getevents) - Get All Events
* [GetIntegrations](#getintegrations) - Get CRM Integrations
* [GetLead](#getlead) - Get Lead
* [GetLeadDetails](#getleaddetails) - Get Lead Details
* [GetLeads](#getleads) - Get all Leads
* [GetNote](#getnote) - Get Note
* [GetNoteDetails](#getnotedetails) - Get Note Details
* [GetNotes](#getnotes) - Get All Notes
* [GetTask](#gettask) - Get Task
* [GetTaskDetails](#gettaskdetails) - Get Task Details
* [GetTasks](#gettasks) - Get All Tasks
* [GetUser](#getuser) - Get User
* [GetUserDetails](#getuserdetails) - Get User Details
* [GetWebhook](#getwebhook) - Get Webhook
* [PatchEventAttendee](#patcheventattendee) - Update Event Attendee
* [PostAccount](#postaccount) - Create Account
* [PostContact](#postcontact) - Create Contact
* [PostCrmTask](#postcrmtask) - Create Task
* [PostDeal](#postdeal) - Create Deal
* [PostEvent](#postevent) - Create Event
* [PostEventAttendee](#posteventattendee) - Create Event Attendee
* [PostLead](#postlead) - Create Lead
* [PostLinkExchange](#postlinkexchange) - Exchange public token for access token
* [PostLinkToken](#postlinktoken) - Create link token
* [PostNote](#postnote) - Create Note
* [PostWebhook](#postwebhook) - Create Webhook
* [PutAccount](#putaccount) - Update Account
* [PutContactJSON](#putcontactjson) - Update Contact
* [PutContactRaw](#putcontactraw) - Update Contact
* [PutDeal](#putdeal) - Update Deal
* [PutEvent](#putevent) - Update Event
* [PutLead](#putlead) - Update Lead
* [PutNote](#putnote) - Update Note
* [PutTask](#puttask) - Update Task

## DeleteConnection

Remove a connection for a given `accessToken`. Removing a connection disconnects the user's CRM so they'll need to re-authenticate should they want to re-connect their CRM.

For CRMs that support it (such as Pipedrive), removing the `accessToken` will also remove the Vessel app from their CRM.

### Example Usage

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
    res, err := s.Snyk.DeleteConnection(ctx, operations.DeleteConnectionRequestBody{
        AccessToken: snyk.String("provident"),
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

## DeleteWebhook

Removes a webhook for a given connection and id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.DeleteWebhook(ctx, operations.DeleteWebhookRequestBody{
        AccessToken: snyk.String("distinctio"),
        WebhookID: snyk.String("quibusdam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetAccount

Retrieve a single Account by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetAccount(ctx, operations.GetAccountRequest{
        AccessToken: "unde",
        AllFields: snyk.Bool(false),
        ID: "d8d69a67-4e0f-4467-8c87-96ed151a05df",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAccount200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetAccountDetails

Get details about all accounts. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetAccountDetails(ctx, operations.GetAccountDetailsRequest{
        AccessToken: "quo",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("odit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAccountDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetConnection

Get info about a connection for a given accessToken.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetConnection(ctx, operations.GetConnectionRequest{
        AccessToken: "at",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetConnection200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetConnections

List all established connections for a workspace

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetConnections(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetConnections200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetContact

Retrieve a Contact by either Id or email. When both email and Id are included, Id will take priority.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetContact(ctx, operations.GetContactRequest{
        AccessToken: "at",
        AllFields: snyk.Bool(false),
        Email: snyk.String("Jaren.Schmidt52@hotmail.com"),
        ID: snyk.String("ca1ba928-fc81-4674-acb7-39205929396f"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContact200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetContactDetails

Get details about all contacts. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetContactDetails(ctx, operations.GetContactDetailsRequest{
        AccessToken: "saepe",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("fuga"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContactDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetContacts

Retrieve all Contacts

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetContacts(ctx, operations.GetContactsRequest{
        AccessToken: "in",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("corporis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContacts200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetCrmAccounts

Retrieve all Accounts

*CRM Caveats*:
- Pipedrive: dealIds + contactIds not supported when querying for all accounts

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetCrmAccounts(ctx, operations.GetCrmAccountsRequest{
        AccessToken: "iste",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("iure"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCrmAccounts200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetCrmDeals

Retrieve all Deals

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetCrmDeals(ctx, operations.GetCrmDealsRequest{
        AccessToken: "saepe",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("quidem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCrmDeals200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetCrmUsers

Retrieve all Users

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetCrmUsers(ctx, operations.GetCrmUsersRequest{
        AccessToken: "architecto",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("ipsa"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCrmUsers200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetDeal

Retrieve a single Deal by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetDeal(ctx, operations.GetDealRequest{
        AccessToken: "reiciendis",
        AllFields: snyk.Bool(false),
        ID: "aaa2352c-5955-4907-aff1-a3a2fa946773",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDeal200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetDealDetails

Get details about all deals or a particular deal. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetDealDetails(ctx, operations.GetDealDetailsRequest{
        AccessToken: "error",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("quia"),
        ID: snyk.String("51aa52c3-f5ad-4019-9a1f-fe78f097b007"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDealDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetEvent

Retrieve a single Event by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetEvent(ctx, operations.GetEventRequest{
        AccessToken: "ut",
        AllFields: snyk.Bool(false),
        ID: "f15471b5-e6e1-43b9-9d48-8e1e91e450ad",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEvent200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetEventAttendee

Retrieve a single Event Attendee by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetEventAttendee(ctx, operations.GetEventAttendeeRequest{
        AccessToken: "explicabo",
        AllFields: snyk.Bool(false),
        ID: "abd44269-802d-4502-a94b-b4f63c969e9a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEventAttendee200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetEventAttendeeDetails

Get details about all event attendees. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetEventAttendeeDetails(ctx, operations.GetEventAttendeeDetailsRequest{
        AccessToken: "dolor",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("debitis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEventAttendeeDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetEventAttendees

Retrieve all Attendees for all Events

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetEventAttendees(ctx, operations.GetEventAttendeesRequest{
        AccessToken: snyk.String("a"),
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("dolorum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEventAttendees200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetEventDetails

Get details about all events. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetEventDetails(ctx, operations.GetEventDetailsRequest{
        AccessToken: "in",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("in"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEventDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetEvents

Retrieve all Events

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetEvents(ctx, operations.GetEventsRequest{
        AccessToken: snyk.String("illum"),
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("maiores"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEvents200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetIntegrations

Return all of the CRM integrations supported by Vessel.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetIntegrations(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetIntegrations200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetLead

Retrieve a single Lead by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetLead(ctx, operations.GetLeadRequest{
        AccessToken: "rerum",
        AllFields: snyk.Bool(false),
        ID: "14cd66ae-395e-4fb9-ba88-f3a66997074b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetLead200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetLeadDetails

Get details about all leads. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetLeadDetails(ctx, operations.GetLeadDetailsRequest{
        AccessToken: "id",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("labore"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetLeadDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetLeads

Retrieve all leads.

*CRM Caveats*:
- Pipedrive: Only `jobTitle` is returned when querying for all leads

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetLeads(ctx, operations.GetLeadsRequest{
        AccessToken: "labore",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("suscipit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetLeads200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetNote

Retrieve a single Note by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetNote(ctx, operations.GetNoteRequest{
        AccessToken: "natus",
        AllFields: snyk.Bool(false),
        ID: "b6e21419-5989-40af-a563-e2516fe4c8b7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetNote200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetNoteDetails

Get details about all notes. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetNoteDetails(ctx, operations.GetNoteDetailsRequest{
        AccessToken: "architecto",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("architecto"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetNoteDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetNotes

Retrieve all Notes

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetNotes(ctx, operations.GetNotesRequest{
        AccessToken: "repudiandae",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("ullam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetNotes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetTask

Retrieve a single Task by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetTask(ctx, operations.GetTaskRequest{
        AccessToken: "expedita",
        AllFields: snyk.Bool(false),
        ID: "7fd2ed02-8921-4cdd-8692-601fb576b0d5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTask200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetTaskDetails

Get details about all tasks. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetTaskDetails(ctx, operations.GetTaskDetailsRequest{
        AccessToken: "voluptatibus",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("perferendis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTaskDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetTasks

Retrieve all Tasks

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetTasks(ctx, operations.GetTasksRequest{
        AccessToken: "fugiat",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("amet"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTasks200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetUser

Retrieve a single User by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetUser(ctx, operations.GetUserRequest{
        AccessToken: "aut",
        AllFields: snyk.Bool(false),
        ID: "c5fbb258-7053-4202-873d-5fe9b90c2890",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUser200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetUserDetails

Get details about all users. 

Details include the type, possible values, and other meta data about the fields.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetUserDetails(ctx, operations.GetUserDetailsRequest{
        AccessToken: "occaecati",
        AllFields: snyk.Bool(false),
        Cursor: snyk.String("rerum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUserDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetWebhook

Retrieve information about a webhook for a given connection and id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.GetWebhook(ctx, operations.GetWebhookRequest{
        AccessToken: snyk.String("adipisci"),
        WebhookID: snyk.String("asperiores"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWebhook200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PatchEventAttendee

Update the status of an event attendee

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PatchEventAttendee(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PatchEventAttendee200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostAccount

Create a new Account

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostAccount(ctx, operations.PostAccountRequestBody{
        AccessToken: "earum",
        Account: shared.AccountCreate{
            Additional: map[string]interface{}{
                "iste": "dolorum",
                "deleniti": "pariatur",
            },
            Address: &shared.Address{
                City: snyk.String("Nettiecester"),
                Country: snyk.String("Venezuela"),
                PostalCode: snyk.String("53222"),
                State: snyk.String("qui"),
                Street: snyk.String("95744 Farrell Port"),
            },
            AnnualRevenue: snyk.Float64(2543.56),
            Description: snyk.String("veritatis"),
            Industry: snyk.String("ipsa"),
            Name: "Vera Kuhlman",
            NumberOfEmployees: snyk.Float64(6963.44),
            Phone: snyk.String("561.608.0764 x46568"),
            Website: snyk.String("distinctio"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAccount200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostContact

Create a new contact.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostContact(ctx, operations.PostContactRequestBody{
        AccessToken: "asperiores",
        Contact: shared.ContactCreate{
            Additional: map[string]interface{}{
                "ipsum": "voluptate",
                "id": "saepe",
            },
            Email: snyk.String("Brigitte75@gmail.com"),
            FirstName: snyk.String("Shannon"),
            JobTitle: snyk.String("Product Interactions Coordinator"),
            LastName: "Oberbrunner",
            MobilePhone: snyk.String("provident"),
            Phone: snyk.String("(856) 283-2478 x169"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostContact200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostCrmTask

Create a new Task.

*CRM Caveats*:
- Salesforce: You may only associate a Task with either a Lead or a Contact *and* either a Deal or an Account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostCrmTask(ctx, operations.PostCrmTaskRequestBody{
        AccessToken: "esse",
        Task: &shared.TaskCreate{
            AccountID: snyk.String("harum"),
            Additional: map[string]interface{}{
                "ipsum": "quisquam",
                "tenetur": "amet",
            },
            Body: snyk.String("tempore"),
            ContactID: snyk.String("accusamus"),
            DealID: snyk.String("numquam"),
            DueDate: snyk.String("enim"),
            IsDone: snyk.Bool(false),
            LeadID: snyk.String("dolorem"),
            Priority: snyk.String("sapiente"),
            Status: snyk.String("totam"),
            Subject: snyk.String("nihil"),
            UserID: snyk.String("sit"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostCrmTask200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostDeal

Create a new Deal

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostDeal(ctx, operations.PostDealRequestBody{
        AccessToken: "expedita",
        Deal: shared.DealCreate{
            AccountID: snyk.String("neque"),
            Additional: map[string]interface{}{
                "vel": "libero",
            },
            Amount: snyk.Float64(3741.7),
            CloseDate: "deserunt",
            Name: "Tracy Gottlieb",
            Probability: snyk.String("maxime"),
            Stage: snyk.String("pariatur"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostDeal200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostEvent

Create a new Event.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostEvent(ctx, operations.PostEventRequestBody{
        AccessToken: "soluta",
        Event: &shared.EventCreate{
            AccountID: snyk.String("dicta"),
            Additional: map[string]interface{}{
                "totam": "incidunt",
                "aspernatur": "dolores",
                "distinctio": "facilis",
            },
            ContactID: snyk.String("aliquid"),
            DealID: snyk.String("quam"),
            Description: snyk.String("molestias"),
            EndDateTime: "temporibus",
            IsAllDayEvent: snyk.Bool(false),
            LeadID: snyk.String("qui"),
            Location: snyk.String("neque"),
            OwnerUserID: snyk.String("fugit"),
            StartDateTime: "magni",
            Subject: snyk.String("odio"),
            Type: snyk.String("sunt"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostEvent200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostEventAttendee

Add an Attendee to an Event

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostEventAttendee(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostEventAttendee200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostLead

Create a new Lead

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostLead(ctx, operations.PostLeadRequestBody{
        AccessToken: "ullam",
        Lead: shared.LeadCreate{
            Account: snyk.String("nam"),
            Additional: map[string]interface{}{
                "voluptatem": "cumque",
                "soluta": "nobis",
                "et": "saepe",
                "ipsum": "veritatis",
            },
            Email: snyk.String("Katrine.Reynolds@hotmail.com"),
            FirstName: snyk.String("Vidal"),
            JobTitle: snyk.String("Dynamic Response Specialist"),
            LastName: "Fisher",
            MobilePhone: snyk.String("dolorum"),
            Phone: snyk.String("(205) 906-8792"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostLead200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostLinkExchange

Exchanges the public token for an access token used to interact with the account. Store the access token in a secure location.

### Example Usage

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
    res, err := s.Snyk.PostLinkExchange(ctx, operations.PostLinkExchangeRequestBody{
        PublicToken: "facilis",
    }, operations.PostLinkExchangeSecurity{
        VesselAPIToken: "YOUR_API_KEY_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostLinkExchange200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostLinkToken

Generates a link token to be used during the auhtentication flow. This token is passed to either the Vessel Link Component or `useVesselLink` hook.

### Example Usage

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
    res, err := s.Snyk.PostLinkToken(ctx, operations.PostLinkTokenSecurity{
        VesselAPIToken: "YOUR_API_KEY_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostLinkToken200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostNote

Create a new Note.

*CRM Caveats*:
- Salesforce: You may only associate a Note with one entity.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostNote(ctx, operations.PostNoteRequestBody{
        AccessToken: "cupiditate",
        Note: shared.NoteCreate{
            AccountID: snyk.String("qui"),
            Additional: map[string]interface{}{
                "laudantium": "odio",
            },
            ContactID: snyk.String("occaecati"),
            Content: "voluptatibus",
            DealID: snyk.String("quisquam"),
            LeadID: snyk.String("vero"),
            UserID: snyk.String("omnis"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostNote200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PostWebhook

Create a new webhook for a given connection

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PostWebhook(ctx, operations.PostWebhookRequestBody{
        AccessToken: snyk.String("quis"),
        Webhook: &shared.WebhookMetadataCreate{
            WebhookURL: snyk.String("ipsum"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostWebhook200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutAccount

Update an existing Account

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutAccount(ctx, operations.PutAccountRequestBody{
        AccessToken: "delectus",
        Account: shared.AccountUpdate{
            Additional: map[string]interface{}{
                "consectetur": "vero",
                "tenetur": "dignissimos",
            },
            Address: &shared.Address{
                City: snyk.String("Portsmouth"),
                Country: snyk.String("Sierra Leone"),
                PostalCode: snyk.String("67842"),
                State: snyk.String("quibusdam"),
                Street: snyk.String("167 Webster Grove"),
            },
            AnnualRevenue: snyk.Float64(8623.1),
            Description: snyk.String("fugit"),
            Industry: snyk.String("porro"),
            Name: snyk.String("Domingo Kris"),
            NumberOfEmployees: snyk.Float64(40.48),
            Phone: snyk.String("1-434-348-2145 x1290"),
            Website: snyk.String("ex"),
        },
        ID: "d9f5fce6-c556-4146-83e2-50fb008c42e1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutAccount200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutContactJSON

Update an existing Contact.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutContactJSON(ctx, operations.PutContactApplicationJSON{
        AccessToken: "non",
        Contact: shared.ContactUpdate{
            Additional: map[string]interface{}{
                "dolorum": "laborum",
            },
            Email: snyk.String("Dejon_Kemmer@hotmail.com"),
            FirstName: snyk.String("Keagan"),
            JobTitle: snyk.String("Internal Configuration Coordinator"),
            LastName: snyk.String("Reinger"),
            MobilePhone: snyk.String("quasi"),
            Phone: snyk.String("415-243-4244 x564"),
        },
        ID: "bd466d28-c10a-4b3c-9ca4-251904e523c7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutContactJSON200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutContactRaw

Update an existing Contact.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutContactRaw(ctx, []byte("recusandae"))
    if err != nil {
        log.Fatal(err)
    }

    if res.PutContactRaw200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutDeal

Update an existing Deal

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutDeal(ctx, operations.PutDealRequestBody{
        AccessToken: "aperiam",
        Deal: shared.DealUpdate{
            AccountID: snyk.String("distinctio"),
            Additional: map[string]interface{}{
                "dignissimos": "inventore",
                "nihil": "totam",
                "accusamus": "aliquam",
                "odio": "occaecati",
            },
            Amount: snyk.Float64(4145.67),
            CloseDate: "sapiente",
            Name: "Miss Rosie Krajcik",
            Probability: snyk.String("quas"),
            Stage: snyk.String("praesentium"),
        },
        ID: "282aa482-562f-4222-a981-7ee17cbe61e6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutDeal200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutEvent

Update an existing Event by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutEvent(ctx, operations.PutEventRequestBody{
        AccessToken: "harum",
        Event: &shared.EventUpdate{
            Additional: map[string]interface{}{
                "rerum": "occaecati",
                "minima": "distinctio",
            },
            Description: snyk.String("eligendi"),
            EndDateTime: snyk.String("sit"),
            IsAllDayEvent: snyk.Bool(false),
            Location: snyk.String("culpa"),
            OwnerUserID: snyk.String("tempore"),
            StartDateTime: snyk.String("adipisci"),
            Subject: snyk.String("cumque"),
            Type: snyk.String("consequuntur"),
        },
        ID: "0c4f3789-fd87-41f9-9dd2-efd121aa6f1e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutEvent200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutLead

Update an existing Lead by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutLead(ctx, operations.PutLeadRequestBody{
        AccessToken: snyk.String("vel"),
        Lead: &shared.LeadUpdate{
            Account: snyk.String("in"),
            Additional: map[string]interface{}{
                "libero": "illum",
                "soluta": "accusantium",
            },
            Email: snyk.String("Veronica.Carter@gmail.com"),
            FirstName: snyk.String("Fleta"),
            JobTitle: snyk.String("Regional Solutions Designer"),
            LastName: "Denesik",
            MobilePhone: snyk.String("quibusdam"),
            Phone: snyk.String("1-696-269-0804 x030"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutLead200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutNote

Update an existing Note by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutNote(ctx, operations.PutNoteRequestBody{
        AccessToken: "consectetur",
        ID: "39d08086-a184-4039-8c26-071f93f5f064",
        Note: shared.NoteUpdate{
            Additional: map[string]interface{}{
                "repellendus": "officia",
            },
            Content: snyk.String("maxime"),
            UserID: snyk.String("dignissimos"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutNote200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## PutTask

Update an existing Task by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/snyk-go"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/snyk-go/pkg/models/shared"
)

func main() {
    s := snyk.New(
        snyk.WithSecurity(shared.Security{
            VesselAPIToken: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Snyk.PutTask(ctx, operations.PutTaskRequestBody{
        AccessToken: "officia",
        ID: "f515cc41-3aa6-43aa-a8d6-7864dbb675fd",
        Task: &shared.TaskUpdate{
            Additional: map[string]interface{}{
                "recusandae": "aliquid",
                "aperiam": "cum",
            },
            Body: snyk.String("consectetur"),
            DueDate: snyk.String("in"),
            IsDone: snyk.Bool(false),
            Priority: snyk.String("exercitationem"),
            Status: snyk.String("earum"),
            Subject: snyk.String("facere"),
            UserID: snyk.String("numquam"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PutTask200ApplicationJSONObject != nil {
        // handle response
    }
}
```
