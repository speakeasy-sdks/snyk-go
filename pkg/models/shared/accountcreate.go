// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountCreate struct {
	// Create any property in the downstream CRM, such as a custom or native property. For a list of all properties see /details.
	Additional map[string]interface{} `json:"additional,omitempty"`
	Address    *Address               `json:"address,omitempty"`
	// Not supported by Pipedrive
	AnnualRevenue *float64 `json:"annualRevenue,omitempty"`
	// Not supported by Pipedrive
	Description *string `json:"description,omitempty"`
	// Not supported by Pipedrive
	Industry *string `json:"industry,omitempty"`
	Name     string  `json:"name"`
	// Not supported by Pipedrive
	NumberOfEmployees *float64 `json:"numberOfEmployees,omitempty"`
	// The primary phone | Not supported by Pipedrive
	Phone *string `json:"phone,omitempty"`
	// Not supported by Pipedrive
	Website *string `json:"website,omitempty"`
}
