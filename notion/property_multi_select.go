package notion

import "github.com/spf13/cast"

const PropertyTypeMultiSelect = "multi_select"

func init() {
	RegisterProperty(&PropertyMultiSelect{})
}

type PropertyMultiSelect struct {
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	MultiSelect []struct {
		Name string `json:"name"`
	} `json:"multi_select"`
}

// GetId implements Property.
func (p *PropertyMultiSelect) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyMultiSelect) GetType() string {
	return PropertyTypeMultiSelect
}

// GetValue implements Property.
func (p *PropertyMultiSelect) GetValue() any {
	v := make([]string, 0, len(p.MultiSelect))
	for _, s := range p.MultiSelect {
		v = append(v, s.Name)
	}
	return v
}

// SetValue implements Property.
func (p *PropertyMultiSelect) SetValue(v any) {
	p.MultiSelect = []struct {
		Name string "json:\"name\""
	}{}
	for _, s := range cast.ToStringSlice(v) {
		p.MultiSelect = append(p.MultiSelect, struct {
			Name string "json:\"name\""
		}{
			Name: s,
		})
	}
}

// Editable implements Property.
func (p *PropertyMultiSelect) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigMultiSelect{})
}

type PropertyConfigMultiSelect struct {
	Id      string `json:"id,omitempty"`
	Type    string `json:"type"`
	Options []struct {
		Id    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Color string `json:"color,omitempty"`
	} `json:"options"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigMultiSelect) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigMultiSelect) GetType() string {
	return PropertyTypeMultiSelect
}
