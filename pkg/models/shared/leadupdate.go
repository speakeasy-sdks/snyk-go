// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LeadUpdate struct {
	// Name of the Account associated with this lead. Not supported for Pipedrive
	Account *string `json:"account,omitempty"`
	// Create any property in the downstream CRM, such as a custom or native property. For a list of all properties see /details.
	Additional map[string]interface{} `json:"additional,omitempty"`
	// Not supported for Pipedrive
	Email *string `json:"email,omitempty"`
	// Not supported for Pipedrive
	FirstName *string `json:"firstName,omitempty"`
	JobTitle  *string `json:"jobTitle,omitempty"`
	// Not supported for Pipedrive
	LastName string `json:"lastName"`
	// The primary mobile phone number for this lead. Possibly the same as the phone number. Not supported for Pipedrive
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// The primary phone. Not supported for Pipedrive
	Phone *string `json:"phone,omitempty"`
}
