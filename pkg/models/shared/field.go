package shared

type FieldOptions struct {
	Key  *string     `json:"key,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

type FieldTypeEnum string

const (
	FieldTypeEnumString       FieldTypeEnum = "string"
	FieldTypeEnumNumber       FieldTypeEnum = "number"
	FieldTypeEnumDatetime     FieldTypeEnum = "datetime"
	FieldTypeEnumDate         FieldTypeEnum = "date"
	FieldTypeEnumBoolean      FieldTypeEnum = "boolean"
	FieldTypeEnumReference    FieldTypeEnum = "reference"
	FieldTypeEnumPhone        FieldTypeEnum = "phone"
	FieldTypeEnumURL          FieldTypeEnum = "url"
	FieldTypeEnumID           FieldTypeEnum = "id"
	FieldTypeEnumEmail        FieldTypeEnum = "email"
	FieldTypeEnumPercent      FieldTypeEnum = "percent"
	FieldTypeEnumSingleselect FieldTypeEnum = "singleselect"
	FieldTypeEnumMultiselect  FieldTypeEnum = "multiselect"
	FieldTypeEnumAddress      FieldTypeEnum = "address"
	FieldTypeEnumDaterange    FieldTypeEnum = "daterange"
	FieldTypeEnumDecimal      FieldTypeEnum = "decimal"
	FieldTypeEnumTime         FieldTypeEnum = "time"
	FieldTypeEnumObject       FieldTypeEnum = "object"
)

// Field
// (Alias: property) A field is a key-value pair on a CRM Object that provides information about that object.
type Field struct {
	Creatable bool           `json:"creatable"`
	Custom    bool           `json:"custom"`
	IsArray   *bool          `json:"isArray,omitempty"`
	Key       string         `json:"key"`
	Name      string         `json:"name"`
	Options   []FieldOptions `json:"options,omitempty"`
	Type      FieldTypeEnum  `json:"type"`
	Universal bool           `json:"universal"`
	Updatable bool           `json:"updatable"`
}
