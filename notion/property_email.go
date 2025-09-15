package notion

import (
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/spf13/cast"
)

const PropertyTypeEmail = "email"

func init() {
	RegisterProperty(&PropertyEmail{})
}

type PropertyEmail struct {
	Id    string  `json:"id,omitempty"`
	Type  string  `json:"type,omitempty"`
	Email *string `json:"email"`
}

// GetId implements Property.
func (p *PropertyEmail) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyEmail) GetType() string {
	return PropertyTypeEmail
}

// GetValue implements Property.
func (p *PropertyEmail) GetValue() any {
	if p.Email != nil {
		return *p.Email
	}
	return ""
}

// SetValue implements Property.
func (p *PropertyEmail) SetValue(v any) {
	if str := cast.ToString(v); str != "" {
		p.Email = types.Pointer(str)
	} else {
		p.Email = nil
	}
}

// Editable implements Property.
func (p *PropertyEmail) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigEmail{})
}

type PropertyConfigEmail struct {
	Id    string   `json:"id,omitempty"`
	Type  string   `json:"type"`
	Email struct{} `json:"email"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigEmail) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigEmail) GetType() string {
	return PropertyTypeEmail
}
