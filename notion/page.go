package notion

type Page struct {
	Object         string     `json:"object,omitempty"`
	Id             string     `json:"id,omitempty"`
	CreatedTime    string     `json:"created_time,omitempty"`
	LastEditedTime string     `json:"last_edited_time,omitempty"`
	Archived       bool       `json:"archived,omitempty"`
	InTrash        bool       `json:"in_trash,omitempty"`
	Properties     Properties `json:"properties,omitempty"`
	Parent         *Parent    `json:"parent,omitempty"`

	//created_by omitted
	//last_edited_by omitted
	//cover omitted
	//icon omitted
	//url omitted
}

func NewPage(parent *Parent) *Page {
	return &Page{
		Parent:     parent,
		Properties: Properties{},
	}
}
