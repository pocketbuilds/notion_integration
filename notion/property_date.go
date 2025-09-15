package notion

import (
	"time"

	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/spf13/cast"
)

const PropertyTypeDate = "date"

func init() {
	RegisterProperty(&PropertyDate{})
}

type PropertyDate struct {
	Id   string     `json:"id,omitempty"`
	Type string     `json:"type,omitempty"`
	Date *DateValue `json:"date"`
}

type DateValue struct {
	End   string `json:"end,omitempty"`
	Start string `json:"start,omitempty"`
}

// GetId implements Property.
func (p *PropertyDate) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyDate) GetType() string {
	return PropertyTypeDate
}

// GetValue implements Property.
func (p *PropertyDate) GetValue() any {
	if p.Date != nil {
		return cast.ToTime(p.Date.Start)
	}
	return nil
}

// SetValue implements Property.
func (p *PropertyDate) SetValue(v any) {
	switch t := v.(type) {
	case types.DateTime:
		p.SetValue(t.Time())
	case time.Time:
		if t.IsZero() {
			p.Date = nil
		} else {
			p.Date = &DateValue{
				Start: t.Format(time.RFC3339),
			}
		}
	default:
		p.Date = nil
	}
}

// Editable implements Property.
func (p *PropertyDate) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigDate{})
}

type PropertyConfigDate struct {
	Id   string   `json:"id,omitempty"`
	Type string   `json:"type"`
	Date struct{} `json:"date"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigDate) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigDate) GetType() string {
	return PropertyTypeDate
}
