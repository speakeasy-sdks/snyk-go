package shared

import (
	"time"
)

type DealAssociations struct {
	AccountIds []string `json:"accountIds"`
	ContactIds []string `json:"contactIds"`
	EventIds   []string `json:"eventIds,omitempty"`
	NoteIds    []string `json:"noteIds"`
	TaskIds    []string `json:"taskIds"`
}

// Deal
// (Alias: Opportunity) An object representing either the opportunity to sell a product to an Account, or in the case of an `isWon` deal, a product sold to an Account.
//
// *CRM Caveats*:
// - Pipedrive: Users must have "Deal probability" turned on for the given pipeline for probability to be configurable and returned.
type Deal struct {
	Additional      map[string]interface{} `json:"additional,omitempty"`
	Amount          *float64               `json:"amount,omitempty"`
	Associations    DealAssociations       `json:"associations"`
	CloseDate       time.Time              `json:"closeDate"`
	CreatedTime     time.Time              `json:"createdTime"`
	ExpectedRevenue *float64               `json:"expectedRevenue,omitempty"`
	ID              string                 `json:"id"`
	IsClosed        *bool                  `json:"isClosed,omitempty"`
	IsWon           *bool                  `json:"isWon,omitempty"`
	ModifiedTime    time.Time              `json:"modifiedTime"`
	Name            *string                `json:"name,omitempty"`
	NativeID        string                 `json:"nativeId"`
	Probability     *string                `json:"probability,omitempty"`
	Stage           *string                `json:"stage,omitempty"`
}
