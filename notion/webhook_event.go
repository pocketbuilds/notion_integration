package notion

import (
	"encoding/json"
	"errors"
)

type WebhookEvent struct {
	Id                string           `json:"id"`
	Timestamp         string           `json:"timestamp"`
	WorkspaceId       string           `json:"workspace_id"`
	WorkspaceName     string           `json:"workspace_name"`
	SubscriptionId    string           `json:"subscription_id"`
	IntegrationId     string           `json:"integration_id"`
	Type              string           `json:"type"`
	Authors           []*WebhookPerson `json:"authors"`
	AccessibleBy      []*WebhookPerson `json:"accessible_by"`
	AttemptNumber     int              `json:"attempt_number"`
	Entity            *WebhookEntity   `json:"entity"`
	Data              *WebhookData     `json:"data"`
	VerificationToken string           `json:"verification_token"`
}

type WebhookEntity struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type WebhookData struct {
	Parent                 *WebhookEntity         `json:"parent"`
	UpdatedBlocks          []*WebhookDataBlock    `json:"updated_blocks"`
	UpdatedPropertyValues  []string               `json:"-"`
	UpdatedPropertySchemas []*WebhookDataProperty `json:"-"`
	PageId                 string                 `json:"page_id"`
}

func (d *WebhookData) UnmarshalJSON(data []byte) error {
	type alias WebhookData
	dummy1 := struct {
		*alias
		UpdatedProperties []string `json:"updated_properties"`
	}{
		alias: (*alias)(d),
	}
	err1 := json.Unmarshal(data, &dummy1)
	if err1 == nil {
		d.UpdatedPropertyValues = dummy1.UpdatedProperties
		return nil
	}

	dummy2 := struct {
		*alias
		UpdatedProperties []*WebhookDataProperty `json:"updated_properties"`
	}{
		alias: (*alias)(d),
	}
	err2 := json.Unmarshal(data, &dummy2)
	if err2 != nil {
		return errors.Join(err1, err2)
	}
	d.UpdatedPropertySchemas = dummy2.UpdatedProperties
	return nil
}

type WebhookDataProperty struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Action string `json:"action"`
}

type WebhookDataBlock struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type WebhookPerson struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
