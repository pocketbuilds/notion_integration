package notion

func (c *Client) Page() *PageService {
	return &PageService{
		client: c,
	}
}

type PageService struct {
	client *Client
}

func (s *PageService) Create(page *Page) (*Page, error) {
	body := map[string]any{
		"parent":     page.Parent,
		"properties": page.Properties,
		// children omitted
		// icon omitted
		// cover omitted
	}
	resp, err := s.client.Post(body, "/pages")
	if err != nil {
		return nil, err
	}
	if err := resp.Bind(page); err != nil {
		return nil, err
	}
	return page, nil
}

func (s *PageService) Retrieve(id string) (*Page, error) {
	resp, err := s.client.Get("/pages/", id)
	if err != nil {
		return nil, err
	}
	page := &Page{}
	if err := resp.Bind(page); err != nil {
		return nil, err
	}
	return page, nil
}

func (s *PageService) Update(page *Page) error {
	properties := map[string]any{}
	for k, p := range page.Properties {
		if p.Editable() {
			properties[k] = p
		}
	}
	body := map[string]any{
		"properties": properties,
	}
	_, err := s.client.Patch(body, "/pages/", page.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PageService) Delete(page *Page) error {
	body := map[string]any{
		"archived": true,
	}
	_, err := s.client.Patch(body, "/pages/", page.Id)
	if err != nil {
		return err
	}
	return nil
}
