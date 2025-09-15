package notion

import (
	"github.com/spf13/cast"
)

const PropertyTypeSelect = "select"

func init() {
	RegisterProperty(&PropertySelect{})
}

type PropertySelect struct {
	Id     string       `json:"id,omitempty"`
	Type   string       `json:"type,omitempty"`
	Select *SelectValue `json:"select"`
}

type SelectValue struct {
	Name string `json:"name"`
}

// GetId implements Property.
func (p *PropertySelect) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertySelect) GetType() string {
	return PropertyTypeSelect
}

// GetValue implements Property.
func (p *PropertySelect) GetValue() any {
	if p.Select != nil {
		return p.Select.Name
	}
	return nil
}

// SetValue implements Property.
func (p *PropertySelect) SetValue(v any) {
	if str := cast.ToString(v); str != "" {
		p.Select = &SelectValue{
			Name: str,
		}
	} else {
		p.Select = nil
	}
}

// Editable implements Property.
func (p *PropertySelect) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigSelect{})
}

type PropertyConfigSelect struct {
	Id      string `json:"id,omitempty"`
	Type    string `json:"type"`
	Options []struct {
		Id    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Color string `json:"color,omitempty"`
	} `json:"options"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigSelect) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigSelect) GetType() string {
	return PropertyTypeSelect
}
