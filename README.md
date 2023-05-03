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

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Snyk SDK](docs/snyk/README.md)

* [DeleteConnection](docs/snyk/README.md#deleteconnection) - Delete connection
* [DeleteWebhook](docs/snyk/README.md#deletewebhook) - Remove Webhook
* [GetAccount](docs/snyk/README.md#getaccount) - Get Account
* [GetAccountDetails](docs/snyk/README.md#getaccountdetails) - Get Account Details
* [GetConnection](docs/snyk/README.md#getconnection) - Get Connection
* [GetConnections](docs/snyk/README.md#getconnections) - Get All Connections
* [GetContact](docs/snyk/README.md#getcontact) - Get Contact
* [GetContactDetails](docs/snyk/README.md#getcontactdetails) - Get Contact Details
* [GetContacts](docs/snyk/README.md#getcontacts) - Get All Contacts
* [GetCrmAccounts](docs/snyk/README.md#getcrmaccounts) - Get All Accounts
* [GetCrmDeals](docs/snyk/README.md#getcrmdeals) - Get All Deals
* [GetCrmUsers](docs/snyk/README.md#getcrmusers) - Get All Users
* [GetDeal](docs/snyk/README.md#getdeal) - Get Deal
* [GetDealDetails](docs/snyk/README.md#getdealdetails) - Get Deal Details
* [GetEvent](docs/snyk/README.md#getevent) - Get Event
* [GetEventAttendee](docs/snyk/README.md#geteventattendee) - Get Event Attendee
* [GetEventAttendeeDetails](docs/snyk/README.md#geteventattendeedetails) - Get Event Attendee Details
* [GetEventAttendees](docs/snyk/README.md#geteventattendees) - Get All Event Attendees
* [GetEventDetails](docs/snyk/README.md#geteventdetails) - Get Event Details
* [GetEvents](docs/snyk/README.md#getevents) - Get All Events
* [GetIntegrations](docs/snyk/README.md#getintegrations) - Get CRM Integrations
* [GetLead](docs/snyk/README.md#getlead) - Get Lead
* [GetLeadDetails](docs/snyk/README.md#getleaddetails) - Get Lead Details
* [GetLeads](docs/snyk/README.md#getleads) - Get all Leads
* [GetNote](docs/snyk/README.md#getnote) - Get Note
* [GetNoteDetails](docs/snyk/README.md#getnotedetails) - Get Note Details
* [GetNotes](docs/snyk/README.md#getnotes) - Get All Notes
* [GetTask](docs/snyk/README.md#gettask) - Get Task
* [GetTaskDetails](docs/snyk/README.md#gettaskdetails) - Get Task Details
* [GetTasks](docs/snyk/README.md#gettasks) - Get All Tasks
* [GetUser](docs/snyk/README.md#getuser) - Get User
* [GetUserDetails](docs/snyk/README.md#getuserdetails) - Get User Details
* [GetWebhook](docs/snyk/README.md#getwebhook) - Get Webhook
* [PatchEventAttendee](docs/snyk/README.md#patcheventattendee) - Update Event Attendee
* [PostAccount](docs/snyk/README.md#postaccount) - Create Account
* [PostContact](docs/snyk/README.md#postcontact) - Create Contact
* [PostCrmTask](docs/snyk/README.md#postcrmtask) - Create Task
* [PostDeal](docs/snyk/README.md#postdeal) - Create Deal
* [PostEvent](docs/snyk/README.md#postevent) - Create Event
* [PostEventAttendee](docs/snyk/README.md#posteventattendee) - Create Event Attendee
* [PostLead](docs/snyk/README.md#postlead) - Create Lead
* [PostLinkExchange](docs/snyk/README.md#postlinkexchange) - Exchange public token for access token
* [PostLinkToken](docs/snyk/README.md#postlinktoken) - Create link token
* [PostNote](docs/snyk/README.md#postnote) - Create Note
* [PostWebhook](docs/snyk/README.md#postwebhook) - Create Webhook
* [PutAccount](docs/snyk/README.md#putaccount) - Update Account
* [PutContactJSON](docs/snyk/README.md#putcontactjson) - Update Contact
* [PutContactRaw](docs/snyk/README.md#putcontactraw) - Update Contact
* [PutDeal](docs/snyk/README.md#putdeal) - Update Deal
* [PutEvent](docs/snyk/README.md#putevent) - Update Event
* [PutLead](docs/snyk/README.md#putlead) - Update Lead
* [PutNote](docs/snyk/README.md#putnote) - Update Note
* [PutTask](docs/snyk/README.md#puttask) - Update Task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
