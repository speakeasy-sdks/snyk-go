// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CrmIntegrationIntegrationID string

const (
	CrmIntegrationIntegrationIDSalesforce CrmIntegrationIntegrationID = "salesforce"
	CrmIntegrationIntegrationIDHubspot    CrmIntegrationIntegrationID = "hubspot"
	CrmIntegrationIntegrationIDPipedrive  CrmIntegrationIntegrationID = "pipedrive"
)

func (e CrmIntegrationIntegrationID) ToPointer() *CrmIntegrationIntegrationID {
	return &e
}

func (e *CrmIntegrationIntegrationID) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "salesforce":
		fallthrough
	case "hubspot":
		fallthrough
	case "pipedrive":
		*e = CrmIntegrationIntegrationID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CrmIntegrationIntegrationID: %v", v)
	}
}

type CrmIntegration struct {
	// Base 64 data URI
	IconURL       *string                      `json:"iconURL,omitempty"`
	IntegrationID *CrmIntegrationIntegrationID `json:"integrationId,omitempty"`
	Name          *string                      `json:"name,omitempty"`
}
