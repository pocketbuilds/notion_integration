package notion

import (
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/spf13/cast"
)

const PropertyTypePhoneNumber = "phone_number"

func init() {
	RegisterProperty(&PropertyPhoneNumber{})
}

type PropertyPhoneNumber struct {
	Id          string  `json:"id,omitempty"`
	Type        string  `json:"type,omitempty"`
	PhoneNumber *string `json:"phone_number"`
}

// GetId implements Property.
func (p *PropertyPhoneNumber) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyPhoneNumber) GetType() string {
	return PropertyTypePhoneNumber
}

// GetValue implements Property.
func (p *PropertyPhoneNumber) GetValue() any {
	if p.PhoneNumber == nil {
		return ""
	}
	return *p.PhoneNumber
}

// SetValue implements Property.
func (p *PropertyPhoneNumber) SetValue(v any) {
	if str := cast.ToString(v); str == "" {
		p.PhoneNumber = nil
	} else {
		p.PhoneNumber = types.Pointer(str)
	}
}

// Editable implements Property.
func (p *PropertyPhoneNumber) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigPhoneNumber{})
}

type PropertyConfigPhoneNumber struct {
	Id          string   `json:"id,omitempty"`
	Type        string   `json:"type"`
	PhoneNumber struct{} `json:"phone_number"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigPhoneNumber) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigPhoneNumber) GetType() string {
	return PropertyTypePhoneNumber
}
