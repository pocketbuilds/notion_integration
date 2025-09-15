package notion

import (
	"github.com/spf13/cast"
)

const PropertyTypeCheckbox = "checkbox"

func init() {
	RegisterProperty(&PropertyCheckbox{})
}

type PropertyCheckbox struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Checkbox bool   `json:"checkbox"`
}

// GetId implements Property.
func (p *PropertyCheckbox) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyCheckbox) GetType() string {
	return PropertyTypeCheckbox
}

// GetValue implements Property.
func (p *PropertyCheckbox) GetValue() any {
	return p.Checkbox
}

// SetValue implements Property.
func (p *PropertyCheckbox) SetValue(v any) {
	p.Checkbox = cast.ToBool(v)
}

// Editable implements Property.
func (p *PropertyCheckbox) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigCheckbox{})
}

type PropertyConfigCheckbox struct {
	Id       string   `json:"id,omitempty"`
	Type     string   `json:"type"`
	Checkbox struct{} `json:"checkbox"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigCheckbox) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigCheckbox) GetType() string {
	return PropertyTypeCheckbox
}
