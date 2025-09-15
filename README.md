# Notion Integration

## Setup

### Create a Notion Database
1. Login or Create a [Notion](https://notion.so/) Account.
2. Create a Notion Page.
3. Turn the page into a Notion database.
4. Add any properties you require for this database.

### Create a Pocketbase Collection
1. Create a Pocketbase collection.
2. Add a Plain Text field (recommended name is "notion_id") that will hold the id to the corresponding Notion page.
3. Add a index with a unique constraint on the Notion Id field.
4. Add any fields that you want to correspond to the properties in your Notion database (see [Supported Notion Property Types](#supported-notion-property-types)

### Create A Notion Integration
1. Go to your [Notion Integrations](https://www.notion.so/profile/integrations).
2. Give your integration access to Read content, Update content, and Insert content. Also, give it access to any Notion database you want to integrate with pocketbase.
3. Create an integration for Pocketbase and add your secret key to your xpb config.

### Edit the Plugin Config

```toml
# pocketbuilds.toml

[notion_integration]
# The base url of the Notion api server.
#   - default: https://api.notion.com/v1
#   - optional
api_base_url = "https://api.notion.com/v1"
# The Notion api verion.
#   - default: 2022-06-28
#   - optional
api_version = "2022-06-28"
# Your Notion integration secret.
#   - required
#   - env: XPB__NOTION_INTEGRATION__SECRET
secret = "xxxxxxxxxxxxxxxxx"

# Array of collections to sync with Notion
[[notion_integration.collections]]
# The name of the pocketbase field that will hold
#   the id of the corresponding Notion page.
#   - recommended: notion_id
#   - required
notion_id_pocketbase_field_name = "notion_id"
# The collection name that will be synced with the
#   Notion database you specify.
#   - required
collection_name = "my_collection"
# The Notion database name that will be synced with the
#   Pocketbase collection you specify.
#   - required
database_id = "my_database"
# Array of pocketbase fields to sync.
[[notion_integration.collections.fields]]
# The name of the Pocketbase field to be synced.
#   - required
pocketbase_name = "my_field"
# The name of the Notion property to be synced.
#   - required
notion_name = "my_property"
```

### Setup Notion Webhooks

Note this requires a domain your pocketbase app to be deployed with https. Without setting up webhooks, your app will work in one direction (changes to records in pocketbase will be reflected in notion but not vise versa).

1. Go to your [Notion Integrations](https://www.notion.so/profile/integrations).
2. Click on your pocketbase integration, and go to "Webhooks".
3. Set your webhook URL (defaults to /notion/webhook).
4. Notion will then ask for a verification code that was sent to the webhook endpoint. This verification code will be emailed to all pocketbase superusers (it will also be printed to stdout if the --dev flag was passed).

## Supported Notion Property Types

Below is the currently supported Notion property types, and their Pocketbase field type counterparts.

| Notion Property Type | Pocketbase Field Type |
| -------------------- | --------------------- |
| Title                | Plain Text            |
| Text                 | Plain Text            |
| Number               | Number                |
| Select               | Select (Single)       |
| Multi-select         | Select (Multiple)     |
| Date                 | DateTime              |
| Checkbox             | Bool                  |
| URL                  | URL                   |
| Email                | Email                 |
| Phone                | Plain Text            |
