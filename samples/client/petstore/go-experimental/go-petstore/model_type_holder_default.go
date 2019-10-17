/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// TypeHolderDefault struct for TypeHolderDefault
type TypeHolderDefault struct {
	StringItem string `json:"string_item"`
	NumberItem float32 `json:"number_item"`
	IntegerItem int32 `json:"integer_item"`
	BoolItem bool `json:"bool_item"`
	ArrayItem []int32 `json:"array_item"`
}

// GetStringItem returns the StringItem field value
func (o *TypeHolderDefault) GetStringItem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StringItem
}

// SetStringItem sets field value
func (o *TypeHolderDefault) SetStringItem(v string) {
	o.StringItem = v
}

// GetNumberItem returns the NumberItem field value
func (o *TypeHolderDefault) GetNumberItem() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NumberItem
}

// SetNumberItem sets field value
func (o *TypeHolderDefault) SetNumberItem(v float32) {
	o.NumberItem = v
}

// GetIntegerItem returns the IntegerItem field value
func (o *TypeHolderDefault) GetIntegerItem() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntegerItem
}

// SetIntegerItem sets field value
func (o *TypeHolderDefault) SetIntegerItem(v int32) {
	o.IntegerItem = v
}

// GetBoolItem returns the BoolItem field value
func (o *TypeHolderDefault) GetBoolItem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BoolItem
}

// SetBoolItem sets field value
func (o *TypeHolderDefault) SetBoolItem(v bool) {
	o.BoolItem = v
}

// GetArrayItem returns the ArrayItem field value
func (o *TypeHolderDefault) GetArrayItem() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ArrayItem
}

// SetArrayItem sets field value
func (o *TypeHolderDefault) SetArrayItem(v []int32) {
	o.ArrayItem = v
}

type NullableTypeHolderDefault struct {
	Value TypeHolderDefault
	ExplicitNull bool
}

func (v NullableTypeHolderDefault) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableTypeHolderDefault) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

