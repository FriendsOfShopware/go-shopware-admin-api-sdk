package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomFieldSetRepository ClientService

func (t CustomFieldSetRepository) Search(ctx ApiContext, criteria Criteria) (*CustomFieldSetCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/custom-field-set", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomFieldSetCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomFieldSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/custom-field-set", criteria)

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

func (t CustomFieldSetRepository) Upsert(ctx ApiContext, entity []CustomFieldSet) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_field_set": {
		Entity:  "custom_field_set",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomFieldSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_field_set": {
		Entity:  "custom_field_set",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomFieldSet struct {
	Products []Product `json:"products,omitempty"`

	App *App `json:"app,omitempty"`

	Global bool `json:"global,omitempty"`

	Position float64 `json:"position,omitempty"`

	AppId string `json:"appId,omitempty"`

	CustomFields []CustomField `json:"customFields,omitempty"`

	Relations []CustomFieldSetRelation `json:"relations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Config interface{} `json:"config,omitempty"`

	Active bool `json:"active,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type CustomFieldSetCollection struct {
	EntityCollection

	Data []CustomFieldSet `json:"data"`
}
