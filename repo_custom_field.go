package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomFieldRepository ClientService

func (t CustomFieldRepository) Search(ctx ApiContext, criteria Criteria) (*CustomFieldCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/custom-field", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomFieldCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomFieldRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/custom-field", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomFieldRepository) Upsert(ctx ApiContext, entity []CustomField) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_field": {
		Entity:  "custom_field",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomFieldRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_field": {
		Entity:  "custom_field",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomField struct {
	Config interface{} `json:"config,omitempty"`

	CustomFieldSet *CustomFieldSet `json:"customFieldSet,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	Active bool `json:"active,omitempty"`

	CustomFieldSetId string `json:"customFieldSetId,omitempty"`

	ProductSearchConfigFields []ProductSearchConfigField `json:"productSearchConfigFields,omitempty"`
}

type CustomFieldCollection struct {
	EntityCollection

	Data []CustomField `json:"data"`
}
