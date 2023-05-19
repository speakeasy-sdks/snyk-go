// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type FieldOptionsNameType string

const (
	FieldOptionsNameTypeStr    FieldOptionsNameType = "str"
	FieldOptionsNameTypeNumber FieldOptionsNameType = "number"
)

type FieldOptionsName struct {
	Str    *string
	Number *float64

	Type FieldOptionsNameType
}

func CreateFieldOptionsNameStr(str string) FieldOptionsName {
	typ := FieldOptionsNameTypeStr

	return FieldOptionsName{
		Str:  &str,
		Type: typ,
	}
}

func CreateFieldOptionsNameNumber(number float64) FieldOptionsName {
	typ := FieldOptionsNameTypeNumber

	return FieldOptionsName{
		Number: &number,
		Type:   typ,
	}
}

func (u *FieldOptionsName) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = FieldOptionsNameTypeStr
		return nil
	}

	number := new(float64)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&number); err == nil {
		u.Number = number
		u.Type = FieldOptionsNameTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u FieldOptionsName) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.Number != nil {
		return json.Marshal(u.Number)
	}

	return nil, nil
}

type FieldOptions struct {
	Key  *string           `json:"key,omitempty"`
	Name *FieldOptionsName `json:"name,omitempty"`
}

type FieldType string

const (
	FieldTypeString       FieldType = "string"
	FieldTypeNumber       FieldType = "number"
	FieldTypeDatetime     FieldType = "datetime"
	FieldTypeDate         FieldType = "date"
	FieldTypeBoolean      FieldType = "boolean"
	FieldTypeReference    FieldType = "reference"
	FieldTypePhone        FieldType = "phone"
	FieldTypeURL          FieldType = "url"
	FieldTypeID           FieldType = "id"
	FieldTypeEmail        FieldType = "email"
	FieldTypePercent      FieldType = "percent"
	FieldTypeSingleselect FieldType = "singleselect"
	FieldTypeMultiselect  FieldType = "multiselect"
	FieldTypeAddress      FieldType = "address"
	FieldTypeDaterange    FieldType = "daterange"
	FieldTypeDecimal      FieldType = "decimal"
	FieldTypeTime         FieldType = "time"
	FieldTypeObject       FieldType = "object"
)

func (e FieldType) ToPointer() *FieldType {
	return &e
}

func (e *FieldType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "string":
		fallthrough
	case "number":
		fallthrough
	case "datetime":
		fallthrough
	case "date":
		fallthrough
	case "boolean":
		fallthrough
	case "reference":
		fallthrough
	case "phone":
		fallthrough
	case "url":
		fallthrough
	case "id":
		fallthrough
	case "email":
		fallthrough
	case "percent":
		fallthrough
	case "singleselect":
		fallthrough
	case "multiselect":
		fallthrough
	case "address":
		fallthrough
	case "daterange":
		fallthrough
	case "decimal":
		fallthrough
	case "time":
		fallthrough
	case "object":
		*e = FieldType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FieldType: %v", v)
	}
}

// Field - (Alias: property) A field is a key-value pair on a CRM Object that provides information about that object.
type Field struct {
	// If this field can be used when creating the object
	Creatable bool `json:"creatable"`
	// If this field is a custom field
	Custom bool `json:"custom"`
	// If this field is an array
	IsArray *bool `json:"isArray,omitempty"`
	// The key in the CRM object (ex: annualRevenue, numberOfEmployees, etc)
	Key string `json:"key"`
	// The display name of this field (ex: number of employees, annual revenue, etc)
	Name string `json:"name"`
	// Possible options for this field
	Options []FieldOptions `json:"options,omitempty"`
	Type    FieldType      `json:"type"`
	// If this is a field we've unified across CRMs
	Universal bool `json:"universal"`
	// If this field can be updated
	Updatable bool `json:"updatable"`
}
