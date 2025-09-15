package notion

import (
	"time"

	"github.com/spf13/cast"
)

const PropertyTypeCreatedTime = "created_time"

func init() {
	RegisterProperty(&PropertyCreatedTime{})
}

type PropertyCreatedTime struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	CreatedTime string `json:"created_time"`
}

// GetId implements Property.
func (p *PropertyCreatedTime) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyCreatedTime) GetType() string {
	return PropertyTypeCreatedTime
}

// GetValue implements Property.
func (p *PropertyCreatedTime) GetValue() any {
	return p.CreatedTime
}

// SetValue implements Property.
func (p *PropertyCreatedTime) SetValue(v any) {
	p.CreatedTime = cast.ToTime(v).UTC().Format(time.RFC3339)
}

// Editable implements Property.
func (p *PropertyCreatedTime) Editable() bool {
	return false
}

func init() {
	RegisterPropertyConfig(&PropertyConfigCreatedTime{})
}

type PropertyConfigCreatedTime struct {
	Id          string   `json:"id,omitempty"`
	Type        string   `json:"type"`
	CreatedTime struct{} `json:"created_time"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigCreatedTime) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigCreatedTime) GetType() string {
	return PropertyTypeCreatedTime
}
