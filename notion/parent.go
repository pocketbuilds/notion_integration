package notion

// https://developers.notion.com/reference/parent-object

const (
	ParentTypeDatabase  = "database_id"
	ParentTypePage      = "page_id"
	ParentTypeWorkspace = "workspace"
	ParentTypeBlock     = "block_id"
)

type Parent struct {
	Type       string `json:"type,omitempty"`
	DatabaseId string `json:"database_id,omitempty"`
	PageId     string `json:"page_id,omitempty"`
	Workspace  bool   `json:"workspace,omitempty"`
	BlockId    string `json:"block_id,omitempty"`
}
