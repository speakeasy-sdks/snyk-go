# Go SDK for Synk API

<div align="center">
   <p>The Snyk API v1 has the ability to test a package for issues as they are defined by Snyk, and to provide Snyk security automation according to your own workflows, unconstrained by security processes in Snyk products. Customers and partners can perform functions including</p>
   <a href="https://resend.com/docs/api-reference/concepts"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/snyk-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/snyk-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/snyk-go/releases"><img src="https://img.shields.io/github/v/release/resendlabs/resend-go?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/snyk-go
```
<!-- End SDK Installation -->

## Authentication

To use the Snyk API, you must first get your API token from Snyk. Go to the General Account Settings in your Snyk account. In the `KEY` field, click to show and then select and copy your API token. Refer to the screenshot that follows.
To use the Snyk API, supply the token in an Authorization header, preceded by `Token`: `Authorization`: `Token API_KEY`

For more information about authentication, refer to the overview in the API reference docs. For additional information see Revoking and regenerating Snyk API tokens.

![snyk](https://user-images.githubusercontent.com/68016351/222008880-db6536c7-1652-4edb-9e9b-611666316f21.png)

## SDK Example Usage

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

<!-- Start SDK Available Operations -->
## SDK Available Operations

### Snyk SDK

* `DeleteConnection` - Delete connection
* `DeleteWebhook` - Remove Webhook
* `GetAccount` - Get Account
* `GetAccountDetails` - Get Account Details
* `GetConnection` - Get Connection
* `GetConnections` - Get All Connections
* `GetContact` - Get Contact
* `GetContactDetails` - Get Contact Details
* `GetContacts` - Get All Contacts
* `GetCrmAccounts` - Get All Accounts
* `GetCrmDeals` - Get All Deals
* `GetCrmUsers` - Get All Users
* `GetDeal` - Get Deal
* `GetDealDetails` - Get Deal Details
* `GetEvent` - Get Event
* `GetEventAttendee` - Get Event Attendee
* `GetEventAttendeeDetails` - Get Event Attendee Details
* `GetEventAttendees` - Get All Event Attendees
* `GetEventDetails` - Get Event Details
* `GetEvents` - Get All Events
* `GetIntegrations` - Get CRM Integrations
* `GetLead` - Get Lead
* `GetLeadDetails` - Get Lead Details
* `GetLeads` - Get all Leads
* `GetNote` - Get Note
* `GetNoteDetails` - Get Note Details
* `GetNotes` - Get All Notes
* `GetTask` - Get Task
* `GetTaskDetails` - Get Task Details
* `GetTasks` - Get All Tasks
* `GetUser` - Get User
* `GetUserDetails` - Get User Details
* `GetWebhook` - Get Webhook
* `PatchEventAttendee` - Update Event Attendee
* `PostAccount` - Create Account
* `PostContact` - Create Contact
* `PostCrmTask` - Create Task
* `PostDeal` - Create Deal
* `PostEvent` - Create Event
* `PostEventAttendee` - Create Event Attendee
* `PostLead` - Create Lead
* `PostLinkExchange` - Exchange public token for access token
* `PostLinkToken` - Create link token
* `PostNote` - Create Note
* `PostWebhook` - Create Webhook
* `PutAccount` - Update Account
* `PutContactJSON` - Update Contact
* `PutContactRaw` - Update Contact
* `PutDeal` - Update Deal
* `PutEvent` - Update Event
* `PutLead` - Update Lead
* `PutNote` - Update Note
* `PutTask` - Update Task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
