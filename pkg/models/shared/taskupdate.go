// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TaskUpdate struct {
	// Create any property in the downstream CRM, such as a custom or native property. For a list of all properties see /details.
	Additional map[string]interface{} `json:"additional,omitempty"`
	Body       *string                `json:"body,omitempty"`
	DueDate    *string                `json:"dueDate,omitempty"`
	// If both `isDone` and `status` are included, `isDone` will take precedence and automatically set `status` to either the default or a 'closed' status depending on the `isDone` value.
	IsDone *bool `json:"isDone,omitempty"`
	// not supported by Pipedrive
	Priority *string `json:"priority,omitempty"`
	// not supported by Pipedrive
	Status  *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	UserID  *string `json:"userId,omitempty"`
}
