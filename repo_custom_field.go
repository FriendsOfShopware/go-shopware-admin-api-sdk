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

func (t CustomFieldRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomFieldCollection, *http.Response, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	c, resp, err := t.Search(ctx, criteria)

	if err != nil {
		return c, resp, err
	}

	for {
		criteria.Page++

		nextC, nextResp, nextErr := t.Search(ctx, criteria)

		if nextErr != nil {
			return c, nextResp, nextErr
		}

		if len(nextC.Data) == 0 {
			break
		}

		c.Data = append(c.Data, nextC.Data...)
	}

	c.Total = int64(len(c.Data))

	return c, resp, err
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
	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Type string `json:"type,omitempty"`

	Config interface{} `json:"config,omitempty"`

	Active bool `json:"active,omitempty"`

	CustomFieldSetId string `json:"customFieldSetId,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFieldSet *CustomFieldSet `json:"customFieldSet,omitempty"`

	ProductSearchConfigFields []ProductSearchConfigField `json:"productSearchConfigFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CustomFieldCollection struct {
	EntityCollection

	Data []CustomField `json:"data"`
}
