package notion

func (c *Client) Database() *DatabaseService {
	return &DatabaseService{
		client: c,
	}
}

type DatabaseService struct {
	client *Client
}

func (s *DatabaseService) Get(id string) (db *Database, err error) {
	resp, err := s.client.Get("/databases/", id)
	if err != nil {
		return nil, err
	}
	db = &Database{}
	if err := resp.Bind(db); err != nil {
		return nil, err
	}
	return db, nil
}

func (s *DatabaseService) Update(id string, body *DatabaseUpdateRequestBody) error {
	_, err := s.client.Patch(body, "/databases/", id)
	if err != nil {
		return err
	}
	return nil
}

type DatabaseUpdateRequestBody struct {
	Properties Properties `json:"properties"`
}

func (s *DatabaseService) Query(id string, body *DatabaseQueryRequestBody) (result *DatabaseQueryResponseBody, err error) {
	resp, err := s.client.Post(body, "/databases/", id, "/query")
	if err != nil {
		return nil, err
	}
	result = new(DatabaseQueryResponseBody)
	if err := resp.Bind(result); err != nil {
		return nil, err
	}
	return result, nil
}

type DatabaseQueryRequestBody struct {
	Filter *DatabaseFilter `json:"filter"`
}

type DatabaseFilter struct {
	Property       string                `json:"property,omitempty"`
	String         *DatabaseStringFilter `json:"string,omitempty"`
	Timestamp      string                `json:"timestamp,omitempty"`
	CreatedTime    *DatabaseDateFilter   `json:"created_time,omitempty"`
	LastEditedTime *DatabaseDateFilter   `json:"last_edited_time,omitempty"`
}

type DatabaseStringFilter struct {
	Equals string `json:"equals,omitempty"`
}

type DatabaseDateFilter struct {
	After      string `json:"after,omitempty"`
	Before     string `json:"before,omitempty"`
	Equals     string `json:"equals,omitempty"`
	IsEmpty    bool   `json:"is_empty,omitempty"`
	IsNotEmpty bool   `json:"is_not_empty,omitempty"`
	OnOrAfter  string `json:"on_or_after,omitempty"`
	OnOrBefore string `json:"on_or_before,omitempty"`
	// Theres more but this is fine for now
}

type DatabaseQueryResponseBody struct {
	Object  string  `json:"object"`
	Results []*Page `json:"results"`
	HasMore bool    `json:"has_more"`
	Type    string  `json:"type"`

	//next_cursor omitted
	//page_or_database omitted
}
