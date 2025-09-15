package notion

type RichText struct {
	Type string        `json:"type,omitempty"`
	Text *RichTextText `json:"text,omitempty"`

	// mention omitted
	// equation omitted
	// annotations omitted
	// plain_text omitted
	// href omitted
}

type RichTextText struct {
	Content string `json:"content"`

	// link omitted
}
