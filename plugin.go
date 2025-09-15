package notion_integration

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"github.com/pocketbuilds/notion_integration/notion"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbuilds/xpb"
)

type Plugin struct {
	ApiBaseUrl  string               `json:"api_base_url"`
	ApiVersion  string               `json:"api_version:"`
	Secret      string               `json:"secret" env:"SECRET"`
	Collections []*CollectionsConfig `json:"collections"`

	client *notion.Client
}

type CollectionsConfig struct {
	NotionIdPocketbaseFieldName string         `json:"notion_id_pocketbase_field_name"`
	CollectionName              string         `json:"collection_name"`
	DatabaseId                  string         `json:"database_id"`
	Fields                      []*FieldConfig `json:"fields"`
}

type FieldConfig struct {
	PocketbaseName string `json:"pocketbase_name"`
	NotionName     string `json:"notion_name"`
}

func init() {
	xpb.Register(&Plugin{
		ApiBaseUrl: "https://api.notion.com/v1",
		ApiVersion: "2022-06-28",
	})
}

const notionSkipContextKey string = "notion_skip"

func (p *Plugin) Name() string {
	return "notion_integration"
}

func (p *Plugin) Version() string {
	return "v0.0.1"
}

func (p *Plugin) Description() string {
	return "Sync Pocketbase records with your Notion pages"
}

func (p *Plugin) Init(app core.App) error {

	app.OnServe().BindFunc(p.setupNotionClient)
	app.OnServe().BindFunc(p.setupWebhook)

	for _, c := range p.Collections {
		app.OnRecordCreate(c.CollectionName).BindFunc(p.saveRecordHandler(c))
		app.OnRecordUpdate(c.CollectionName).BindFunc(p.saveRecordHandler(c))
		app.OnRecordDelete(c.CollectionName).BindFunc(p.deleteRecordHandler(c))
	}

	return nil
}

func (p *Plugin) setupNotionClient(e *core.ServeEvent) (err error) {
	p.client, err = notion.NewClient(p.Secret,
		notion.ClientWithApiBaseUrl(p.ApiBaseUrl),
		notion.ClientWithApiVersion(p.ApiVersion),
		notion.ClientWithLogger(e.App.Logger()),
	)
	if err != nil {
		return err
	}
	return e.Next()
}

func (p *Plugin) setupWebhook(e *core.ServeEvent) error {
	e.Router.POST("/notion/webhook", p.handleWebhook)
	return e.Next()
}

func (p *Plugin) handleWebhook(e *core.RequestEvent) error {
	we := &notion.WebhookEvent{}
	if err := e.BindBody(we); err != nil {
		return e.BadRequestError("failed to unmarshal json request body", err)
	}
	if we.VerificationToken != "" {
		return p.handleVerification(e, we.VerificationToken)
	}
	switch we.Type {
	case "page.created":
		fallthrough
	case "page.properties_updated":
		return p.handleWebhookPageCreatedOrPropertiesUpdated(e, we)
	case "page.deleted":
		return p.handleWebhookPageDeleted(e, we)
	}
	return e.NoContent(http.StatusOK)
}

func (p *Plugin) handleVerification(e *core.RequestEvent, token string) error {
	if e.App.IsDev() {
		fmt.Println("notion verification token:", token)
	}
	superusers, err := e.App.FindAllRecords(core.CollectionNameSuperusers)
	if err != nil {
		return e.InternalServerError("failed to query superusers to notifiy of verification token", err)
	}
	to := []mail.Address{}
	for _, su := range superusers {
		to = append(to, mail.Address{
			Address: su.Email(),
		})
	}
	msg := &mailer.Message{
		From: mail.Address{
			Name:    e.App.Settings().Meta.SenderName,
			Address: e.App.Settings().Meta.SenderAddress,
		},
		To:      to,
		Subject: "Notion Verification",
		Text: fmt.Sprintf(
			"A Notion integration verification token received:\n\n%s",
			token,
		),
	}
	if err := e.App.NewMailClient().Send(msg); err != nil {
		return e.InternalServerError("failed to notify superusers", err)
	}
	return e.NoContent(http.StatusOK)
}

func (p *Plugin) handleWebhookPageCreatedOrPropertiesUpdated(e *core.RequestEvent, we *notion.WebhookEvent) error {
	page, err := p.client.Page().Retrieve(we.Entity.Id)
	if err != nil {
		return err
	}
	databaseId := strings.ReplaceAll(page.Parent.DatabaseId, "-", "")
	var record *core.Record
	for _, c := range p.Collections {
		if c.DatabaseId == databaseId {
			collection, err := e.App.FindCachedCollectionByNameOrId(c.CollectionName)
			if err != nil {
				return err
			}
			record, err = e.App.FindFirstRecordByData(collection, c.NotionIdPocketbaseFieldName, page.Id)
			if err != nil || record == nil {
				record = core.NewRecord(collection)
			}
			record.Set(c.NotionIdPocketbaseFieldName, page.Id)
			for _, f := range c.Fields {
				record.Set(f.PocketbaseName, page.Properties.Get(f.NotionName))
			}
			ctx := context.WithValue(e.Request.Context(), notionSkipContextKey, true)
			return e.App.SaveWithContext(ctx, record)
		}
	}
	return e.NoContent(http.StatusOK)
}

func (p *Plugin) handleWebhookPageDeleted(e *core.RequestEvent, we *notion.WebhookEvent) error {
	pageId := we.Entity.Id
	databaseId := strings.ReplaceAll(we.Data.Parent.Id, "-", "")
	for _, c := range p.Collections {
		if c.DatabaseId == databaseId {
			record, err := e.App.FindFirstRecordByData(c.CollectionName, c.NotionIdPocketbaseFieldName, pageId)
			if err != nil {
				return nil // TODO: log not found
			}
			ctx := context.WithValue(e.Request.Context(), notionSkipContextKey, true)
			return e.App.DeleteWithContext(ctx, record)
		}
	}
	return e.NoContent(http.StatusOK)
}

func (p *Plugin) saveRecordHandler(config *CollectionsConfig) func(e *core.RecordEvent) error {
	return func(e *core.RecordEvent) error {
		if skip, ok := e.Context.Value(notionSkipContextKey).(bool); ok && skip {
			return e.Next()
		}

		if pageId := e.Record.GetString(config.NotionIdPocketbaseFieldName); pageId != "" {
			page, err := p.client.Page().Retrieve(pageId)
			if err != nil {
				return err
			}
			for _, field := range config.Fields {
				page.Properties.Set(field.NotionName, e.Record.Get(field.PocketbaseName))
			}
			if err := p.client.Page().Update(page); err != nil {
				return err
			}
			return e.Next()
		}

		db, err := p.client.Database().Get(config.DatabaseId)
		if err != nil {
			return err
		}

		page := notion.NewPage(&notion.Parent{
			Type:       notion.ParentTypeDatabase,
			DatabaseId: config.DatabaseId,
		})

		for _, field := range config.Fields {
			prop := db.Properties[field.NotionName]
			if prop == nil {
				continue
			}
			pageProp := notion.NewProperty(prop.GetType())
			if pageProp == nil {
				continue
			}
			pageProp.SetValue(e.Record.Get(field.PocketbaseName))
			page.Properties[field.NotionName] = pageProp
		}
		page, err = p.client.Page().Create(page)
		if err != nil {
			return err
		}
		e.Record.Set(config.NotionIdPocketbaseFieldName, page.Id)

		return e.Next()
	}
}

func (p *Plugin) deleteRecordHandler(config *CollectionsConfig) func(e *core.RecordEvent) error {
	return func(e *core.RecordEvent) error {
		if skip, ok := e.Context.Value(notionSkipContextKey).(bool); ok && skip {
			return e.Next()
		}
		if pageId := e.Record.GetString(config.NotionIdPocketbaseFieldName); pageId != "" {
			page, err := p.client.Page().Retrieve(pageId)
			if err != nil {
				return err
			}
			if page.Archived {
				return e.Next()
			}
			if err := p.client.Page().Delete(page); err != nil {
				return err
			}
		}
		return e.Next()
	}
}
