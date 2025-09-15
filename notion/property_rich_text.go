package notion

import (
	"github.com/spf13/cast"
)

const PropertyTypeRichText = "rich_text"

func init() {
	RegisterProperty(&PropertyRichText{})
}

type PropertyRichText struct {
	Id       string      `json:"id,omitempty"`
	Type     string      `json:"type,omitempty"`
	RichText []*RichText `json:"rich_text"`
}

// GetId implements Property.
func (p *PropertyRichText) GetId() string {
	return p.Id
}

// GetType implements Property.
func (p *PropertyRichText) GetType() string {
	return PropertyTypeRichText
}

// GetValue implements Property.
func (p *PropertyRichText) GetValue() any {
	// TODO: improve
	if len(p.RichText) != 0 {
		return p.RichText[0].Text.Content
	}
	return nil
}

// SetValue implements Property.
func (p *PropertyRichText) SetValue(v any) {
	p.RichText = []*RichText{
		{
			Text: &RichTextText{
				Content: cast.ToString(v),
			},
		},
	}
}

// Editable implements Property.
func (p *PropertyRichText) Editable() bool {
	return true
}

func init() {
	RegisterPropertyConfig(&PropertyConfigRichText{})
}

type PropertyConfigRichText struct {
	Id       string   `json:"id,omitempty"`
	Type     string   `json:"type"`
	RichText struct{} `json:"rich_text"`
}

// GetId implements PropertyConfig.
func (p *PropertyConfigRichText) GetId() string {
	return p.Id
}

// GetType implements PropertyConfig.
func (p *PropertyConfigRichText) GetType() string {
	return PropertyTypeRichText
}
