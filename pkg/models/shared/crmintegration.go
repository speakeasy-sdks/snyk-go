package shared

type CrmIntegrationIntegrationIDEnum string

const (
	CrmIntegrationIntegrationIDEnumSalesforce CrmIntegrationIntegrationIDEnum = "salesforce"
	CrmIntegrationIntegrationIDEnumHubspot    CrmIntegrationIntegrationIDEnum = "hubspot"
	CrmIntegrationIntegrationIDEnumPipedrive  CrmIntegrationIntegrationIDEnum = "pipedrive"
)

type CrmIntegration struct {
	IconURL       *string                          `json:"iconURL,omitempty"`
	IntegrationID *CrmIntegrationIntegrationIDEnum `json:"integrationId,omitempty"`
	Name          *string                          `json:"name,omitempty"`
}
