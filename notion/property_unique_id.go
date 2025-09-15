package notion

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

const PropertyTypeUniqueId = "unique_id"

func init() {
	RegisterProperty(&PropertyUniqueId{})
}

type PropertyUniqueId struct {
	Id       string                 `json:"id,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Type     string                 `json:"type,omitempty"`
	UniqueId *PropertyUniqueIdValue `json:"unique_id"`
}

type PropertyUniqueIdValue struct {
	Number int64  `json:"number"`
	Prefix string `json:"prefix"`
}

// GetId implements Property.
func (p *PropertyUniqueId) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyUniqueId) GetType() string {
	return PropertyTypeUniqueId
}

// GetValue implements Property.
func (p *PropertyUniqueId) GetValue() any {
	var result string
	if p.UniqueId.Prefix != "" {
		result = fmt.Sprintf("%s.", p.UniqueId.Prefix)
	}
	return result + fmt.Sprintf("%d", p.UniqueId.Number)
}

// SetValue implements Property.
func (p *PropertyUniqueId) SetValue(v any) {
	num := cast.ToString(v)
	if i := strings.LastIndex(num, "."); i != -1 {
		p.UniqueId.Prefix, num = num[:i], num[i+1:]
	}
	p.UniqueId.Number = cast.ToInt64(num)
}

// Editable implements Property.
func (p *PropertyUniqueId) Editable() bool {
	return false
}

func init() {
	RegisterPropertyConfig(&PropertyConfigUniqueId{})
}

type PropertyConfigUniqueId struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type"`
	UniqueId struct {
		Prefix string `json:"prefix"`
	} `json:"unique_id"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigUniqueId) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigUniqueId) GetType() string {
	return PropertyTypeUniqueId
}
