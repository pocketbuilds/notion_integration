package notion

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWebhookEventUnmarshallJSON(t *testing.T) {

	t.Run("page.created", func(t *testing.T) {
		data := `{
			"id": "367cba44-b6f3-4c92-81e7-6a2e9659efd4",
			"timestamp": "2024-12-05T23:55:34.285Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.created",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-809d-8dc4-ff2d96ae3090",
				"type": "page"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "367cba44-b6f3-4c92-81e7-6a2e9659efd4",
			Timestamp:      "2024-12-05T23:55:34.285Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.created",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-809d-8dc4-ff2d96ae3090",
				Type: "page",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("page.properties_updated", func(t *testing.T) {
		data := `{
			"id": "1782edd6-a853-4d4a-b02c-9c8c16f28e53",
			"timestamp": "2024-12-05T23:57:05.379Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.properties_updated",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-809d-8dc4-ff2d96ae3090",
				"type": "page"
			},
			"data": {
				"parent": {
					"id": "13950b26-c203-4f3b-b97d-93ec06319565",
					"type": "space"
				},
				"updated_properties": ["XGe%40", "bDf%5B", "DbAu"]
			}
		}`
		expected := WebhookEvent{
			Id:             "1782edd6-a853-4d4a-b02c-9c8c16f28e53",
			Timestamp:      "2024-12-05T23:57:05.379Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.properties_updated",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-809d-8dc4-ff2d96ae3090",
				Type: "page",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "13950b26-c203-4f3b-b97d-93ec06319565",
					Type: "space",
				},
				UpdatedPropertyValues: []string{"XGe%40", "bDf%5B", "DbAu"},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("page.content_updated", func(t *testing.T) {
		data := `{
			"id": "56c3e00c-4f0c-4566-9676-4b058a50a03d",
			"timestamp": "2024-12-05T19:49:36.997Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.content_updated",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
				"type": "page"
			},
			"data": {
				"updated_blocks": [
					{
						"id": "153104cd-477e-80ec-a87d-f7ff0236d35c",
						"type": "block"
					}
				],
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "56c3e00c-4f0c-4566-9676-4b058a50a03d",
			Timestamp:      "2024-12-05T19:49:36.997Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.content_updated",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
				Type: "page",
			},
			Data: &WebhookData{
				UpdatedBlocks: []*WebhookDataBlock{
					{
						Id:   "153104cd-477e-80ec-a87d-f7ff0236d35c",
						Type: "block",
					},
				},
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("page.moved", func(t *testing.T) {
		data := `{
			"id": "7de99a6f-2edd-4116-bf59-2d09407bddec",
			"timestamp": "2024-12-11T05:43:14.383Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.moved",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "154104cd-477e-8030-9989-d4daf352d900",
				"type": "page"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "7de99a6f-2edd-4116-bf59-2d09407bddec",
			Timestamp:      "2024-12-11T05:43:14.383Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.moved",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "154104cd-477e-8030-9989-d4daf352d900",
				Type: "page",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("page.deleted", func(t *testing.T) {
		data := `{
			"id": "ea6b8136-1db6-4f2e-b157-84a532437f62",
			"timestamp": "2024-12-05T23:59:31.215Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.deleted",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-8001-935c-c4b11828dfbd",
				"type": "page"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "ea6b8136-1db6-4f2e-b157-84a532437f62",
			Timestamp:      "2024-12-05T23:59:31.215Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.deleted",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-8001-935c-c4b11828dfbd",
				Type: "page",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("page.undeleted", func(t *testing.T) {
		data := `{
			"id": "ec37232c-a17b-4f02-bb7c-8d8e1f5f2250",
			"timestamp": "2024-12-06T00:00:03.356Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.undeleted",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-8001-935c-c4b11828dfbd",
				"type": "page"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "ec37232c-a17b-4f02-bb7c-8d8e1f5f2250",
			Timestamp:      "2024-12-06T00:00:03.356Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.undeleted",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-8001-935c-c4b11828dfbd",
				Type: "page",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("page.locked", func(t *testing.T) {
		data := `{
			"id": "e2a3092c-5af0-442f-9d11-b813145edb72",
			"timestamp": "2024-12-06T00:00:56.480Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.locked",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-8001-935c-c4b11828dfbd",
				"type": "page"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "e2a3092c-5af0-442f-9d11-b813145edb72",
			Timestamp:      "2024-12-06T00:00:56.480Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.locked",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-8001-935c-c4b11828dfbd",
				Type: "page",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("page.unlocked", func(t *testing.T) {
		data := `{
			"id": "e2a3092c-5af0-442f-9d11-b813145edb72",
			"timestamp": "2024-12-06T00:00:56.480Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "page.unlocked",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-8001-935c-c4b11828dfbd",
				"type": "page"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "e2a3092c-5af0-442f-9d11-b813145edb72",
			Timestamp:      "2024-12-06T00:00:56.480Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "page.unlocked",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-8001-935c-c4b11828dfbd",
				Type: "page",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("database.created", func(t *testing.T) {
		data := `{
			"id": "d0bd8927-0826-4db0-9e26-83d57253f1ff",
			"timestamp": "2024-12-05T23:50:35.868Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "database.created",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				"type": "database"
			},
			"data": {
				"parent": {
					"id": "153104cd-477e-803a-88dc-caececf26478",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "d0bd8927-0826-4db0-9e26-83d57253f1ff",
			Timestamp:      "2024-12-05T23:50:35.868Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "database.created",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				Type: "database",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "153104cd-477e-803a-88dc-caececf26478",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("database.content_updated", func(t *testing.T) {
		data := `{
			"id": "25e44fe0-6785-45bb-adc2-a321526c12c5",
			"timestamp": "2024-12-13T17:48:13.700Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "database.content_updated",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "15b104cd-477e-80c2-84a0-c32cefba5cff",
				"type": "database"
			},
			"data": {
				"updated_blocks": [
					{
						"id": "15b104cd-477e-80a4-bff3-cd05428a4d55",
						"type": "block"
					},
					{
						"id": "15b104cd-477e-80be-98e7-cdf0897fa5c9",
						"type": "block"
					}
				],
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "25e44fe0-6785-45bb-adc2-a321526c12c5",
			Timestamp:      "2024-12-13T17:48:13.700Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "database.content_updated",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "15b104cd-477e-80c2-84a0-c32cefba5cff",
				Type: "database",
			},
			Data: &WebhookData{
				UpdatedBlocks: []*WebhookDataBlock{
					{
						Id:   "15b104cd-477e-80a4-bff3-cd05428a4d55",
						Type: "block",
					},
					{
						Id:   "15b104cd-477e-80be-98e7-cdf0897fa5c9",
						Type: "block",
					},
				},
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("database.moved", func(t *testing.T) {
		data := `{
			"id": "f9c70013-d79d-4c4e-8d5b-939429949a2e",
			"timestamp": "2024-12-06T06:54:08.468Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "database.moved",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				"type": "database"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "f9c70013-d79d-4c4e-8d5b-939429949a2e",
			Timestamp:      "2024-12-06T06:54:08.468Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "database.moved",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				Type: "database",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("database.deleted", func(t *testing.T) {
		data := `{
			"id": "c00e2ea7-032a-4e20-ae05-d69028a09ae9",
			"timestamp": "2024-12-05T23:51:27.295Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "database.deleted",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				"type": "database"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "c00e2ea7-032a-4e20-ae05-d69028a09ae9",
			Timestamp:      "2024-12-05T23:51:27.295Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "database.deleted",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				Type: "database",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("database.undeleted", func(t *testing.T) {
		data := `{
			"id": "edd8ff6e-0f07-4621-934b-76ca55129cc2",
			"timestamp": "2024-12-05T23:52:16.149Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "database.undeleted",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				"type": "database"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "edd8ff6e-0f07-4621-934b-76ca55129cc2",
			Timestamp:      "2024-12-05T23:52:16.149Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "database.undeleted",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				Type: "database",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("database.schema_updated", func(t *testing.T) {
		data := `{
			"id": "5496f509-6988-4bab-b6a9-bdce0b720ca0",
			"timestamp": "2024-12-05T23:55:22.243Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "database.schema_updated",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				"type": "database"
			},
			"data": {
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				},
				"updated_properties": [
					{
						"id": "kqLW",
						"name": "Created at",
						"action": "created"
					},
					{
						"id": "wX%7Bd",
						"name": "Blurb",
						"action": "updated"
					},
					{
						"id": "LIM%5D",
						"name": "Description",
						"action": "deleted"
					}
				]
			}
		}`
		expected := WebhookEvent{
			Id:             "5496f509-6988-4bab-b6a9-bdce0b720ca0",
			Timestamp:      "2024-12-05T23:55:22.243Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "database.schema_updated",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-80eb-ae76-e1c2a32c7b35",
				Type: "database",
			},
			Data: &WebhookData{
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
				UpdatedPropertySchemas: []*WebhookDataProperty{
					{
						Id:     "kqLW",
						Name:   "Created at",
						Action: "created",
					},
					{
						Id:     "wX%7Bd",
						Name:   "Blurb",
						Action: "updated",
					},
					{
						Id:     "LIM%5D",
						Name:   "Description",
						Action: "deleted",
					},
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("comment.created", func(t *testing.T) {
		data := `{
			"id": "c6780f24-10b7-4f42-a6fd-230b6cf7ad69",
			"timestamp": "2024-12-05T20:46:45.854Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "comment.created",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-80ca-8f75-001d9e2b6839",
				"type": "comment"
			},
			"data": {
				"page_id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "c6780f24-10b7-4f42-a6fd-230b6cf7ad69",
			Timestamp:      "2024-12-05T20:46:45.854Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "comment.created",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-80ca-8f75-001d9e2b6839",
				Type: "comment",
			},
			Data: &WebhookData{
				PageId: "0ef104cd-477e-80e1-8571-cfd10e92339a",
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})

	t.Run("comment.deleted", func(t *testing.T) {
		data := `{
			"id": "68ad06e4-5b68-498d-8812-9a1d3e069e46",
			"timestamp": "2024-12-05T20:47:22.657Z",
			"workspace_id": "13950b26-c203-4f3b-b97d-93ec06319565",
			"workspace_name": "Quantify Labs",
			"subscription_id": "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			"integration_id": "0ef2e755-4912-8096-91c1-00376a88a5ca",
			"type": "comment.deleted",
			"authors": [
				{
					"id": "c7c11cca-1d73-471d-9b6e-bdef51470190",
					"type": "person"
				}
			],
			"accessible_by": [
				{
					"id": "556a1abf-4f08-40c6-878a-75890d2a88ba",
					"type": "person"
				},
				{
					"id": "1edc05f6-2702-81b5-8408-00279347f034",
					"type": "bot"
				}
			],
			"attempt_number": 1,
			"entity": {
				"id": "153104cd-477e-80ca-8f75-001d9e2b6839",
				"type": "comment"
			},
			"data": {
				"page_id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
				"parent": {
					"id": "0ef104cd-477e-80e1-8571-cfd10e92339a",
					"type": "page"
				}
			}
		}`
		expected := WebhookEvent{
			Id:             "68ad06e4-5b68-498d-8812-9a1d3e069e46",
			Timestamp:      "2024-12-05T20:47:22.657Z",
			WorkspaceId:    "13950b26-c203-4f3b-b97d-93ec06319565",
			WorkspaceName:  "Quantify Labs",
			SubscriptionId: "29d75c0d-5546-4414-8459-7b7a92f1fc4b",
			IntegrationId:  "0ef2e755-4912-8096-91c1-00376a88a5ca",
			Type:           "comment.deleted",
			Authors: []*WebhookPerson{
				{
					Id:   "c7c11cca-1d73-471d-9b6e-bdef51470190",
					Type: "person",
				},
			},
			AccessibleBy: []*WebhookPerson{
				{
					Id:   "556a1abf-4f08-40c6-878a-75890d2a88ba",
					Type: "person",
				},
				{
					Id:   "1edc05f6-2702-81b5-8408-00279347f034",
					Type: "bot",
				},
			},
			AttemptNumber: 1,
			Entity: &WebhookEntity{
				Id:   "153104cd-477e-80ca-8f75-001d9e2b6839",
				Type: "comment",
			},
			Data: &WebhookData{
				PageId: "0ef104cd-477e-80e1-8571-cfd10e92339a",
				Parent: &WebhookEntity{
					Id:   "0ef104cd-477e-80e1-8571-cfd10e92339a",
					Type: "page",
				},
			},
		}

		var result WebhookEvent
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got:\n%#v\nexpected:\n%#v", result, expected)
		}
	})
}
