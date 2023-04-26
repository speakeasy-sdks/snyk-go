// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CrmIntegrationIntegrationIDEnum string

const (
	CrmIntegrationIntegrationIDEnumSalesforce CrmIntegrationIntegrationIDEnum = "salesforce"
	CrmIntegrationIntegrationIDEnumHubspot    CrmIntegrationIntegrationIDEnum = "hubspot"
	CrmIntegrationIntegrationIDEnumPipedrive  CrmIntegrationIntegrationIDEnum = "pipedrive"
)

func (e CrmIntegrationIntegrationIDEnum) ToPointer() *CrmIntegrationIntegrationIDEnum {
	return &e
}

func (e *CrmIntegrationIntegrationIDEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "salesforce":
		fallthrough
	case "hubspot":
		fallthrough
	case "pipedrive":
		*e = CrmIntegrationIntegrationIDEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CrmIntegrationIntegrationIDEnum: %s", s)
	}
}

type CrmIntegration struct {
	// Base 64 data URI
	IconURL       *string                          `json:"iconURL,omitempty"`
	IntegrationID *CrmIntegrationIntegrationIDEnum `json:"integrationId,omitempty"`
	Name          *string                          `json:"name,omitempty"`
}
