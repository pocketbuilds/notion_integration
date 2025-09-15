package notion

type Database struct {
	Object         string          `json:"object,omitempty"`
	Id             string          `json:"id,omitempty"`
	CreatedTime    string          `json:"created_time,omitempty"`
	LastEditedTime string          `json:"last_edited_time,omitempty"`
	Title          []*RichText     `json:"title,omitempty"`
	Description    []*RichText     `json:"description,omitempty"`
	Properties     PropertyConfigs `json:"properties,omitempty"`

	//last_edited_by omitted
	//icon omitted
	//cover omitted
	//parent omitted
	//url omitted
	//in_trash omitted
	//is_inline omitted
	//public_url omitted
}
