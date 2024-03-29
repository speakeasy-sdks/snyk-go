// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// NoteUpdate - Update an existing Note. Please note that updating associations is not currently supported.
type NoteUpdate struct {
	// Create any property in the downstream CRM, such as a custom or native property. For a list of all properties see /details.
	Additional map[string]interface{} `json:"additional,omitempty"`
	Content    *string                `json:"content,omitempty"`
	// Only Admins can change this property
	UserID *string `json:"userId,omitempty"`
}
