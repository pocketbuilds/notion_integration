package notion

import (
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/spf13/cast"
)

const PropertyTypeUrl = "url"

func init() {
	RegisterProperty(&PropertyUrl{})
}

type PropertyUrl struct {
	Id   string  `json:"id,omitempty"`
	Type string  `json:"type,omitempty"`
	Url  *string `json:"url"`
}

// GetId implements Property.
func (p *PropertyUrl) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyUrl) GetType() string {
	return PropertyTypeUrl
}

// GetValue implements Property.
func (p *PropertyUrl) GetValue() any {
	if p.Url != nil {
		return *p.Url
	}
	return ""
}

// SetValue implements Property.
func (p *PropertyUrl) SetValue(v any) {
	if str := cast.ToString(v); str != "" {
		p.Url = types.Pointer(str)
	} else {
		p.Url = nil
	}
}

// Editable implements Property.
func (p *PropertyUrl) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigUrl{})
}

type PropertyConfigUrl struct {
	Id   string   `json:"id,omitempty"`
	Type string   `json:"type"`
	Url  struct{} `json:"url"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigUrl) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigUrl) GetType() string {
	return PropertyTypeUrl
}
