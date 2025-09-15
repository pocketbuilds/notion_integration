package notion

import (
	"github.com/spf13/cast"
)

const PropertyTypeTitle = "title"

func init() {
	RegisterProperty(&PropertyTitle{})
}

type PropertyTitle struct {
	Id    string      `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Type  string      `json:"type,omitempty"`
	Title []*RichText `json:"title"`
}

// GetId implements Property.
func (p *PropertyTitle) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyTitle) GetType() string {
	return PropertyTypeTitle
}

// GetValue implements Property.
func (p *PropertyTitle) GetValue() any {
	// TODO: improve
	if len(p.Title) != 0 {
		return p.Title[0].Text.Content
	}
	return nil
}

// SetValue implements Property.
func (p *PropertyTitle) SetValue(v any) {
	p.Title = []*RichText{
		{
			Text: &RichTextText{
				Content: cast.ToString(v),
			},
		},
	}
}

// Editable implements Property.
func (p *PropertyTitle) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigTitle{})
}

type PropertyConfigTitle struct {
	Id    string   `json:"id,omitempty"`
	Type  string   `json:"type"`
	Title struct{} `json:"title"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigTitle) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigTitle) GetType() string {
	return PropertyTypeTitle
}
