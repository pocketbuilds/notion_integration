package notion

import (
	"github.com/spf13/cast"
)

const PropertyTypeNumber = "number"

func init() {
	RegisterProperty(&PropertyNumber{})
}

type PropertyNumber struct {
	Id     string  `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Number float64 `json:"number"`
}

// GetId implements Property.
func (p *PropertyNumber) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyNumber) GetType() string {
	return PropertyTypeNumber
}

// GetValue implements Property.
func (p *PropertyNumber) GetValue() any {
	return p.Number
}

// SetValue implements Property.
func (p *PropertyNumber) SetValue(v any) {
	p.Number = cast.ToFloat64(v)
}

// Editable implements Property.
func (p *PropertyNumber) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigNumber{})
}

type PropertyConfigNumber struct {
	Id     string `json:"id,omitempty"`
	Type   string `json:"type"`
	Number struct {
		Format string `json:"format"`
	} `json:"number"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigNumber) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigNumber) GetType() string {
	return PropertyTypeNumber
}
