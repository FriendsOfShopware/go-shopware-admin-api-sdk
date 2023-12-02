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

func (t CustomFieldSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomFieldSetCollection, *http.Response, error) {
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

	Active      bool  `json:"active,omitempty"`

	Global      bool  `json:"global,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Id      string  `json:"id,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	App      *App  `json:"app,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CustomFields      []CustomField  `json:"customFields,omitempty"`

	Relations      []CustomFieldSetRelation  `json:"relations,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	Name      string  `json:"name,omitempty"`

}

type CustomFieldSetCollection struct {
	EntityCollection

	Data []CustomFieldSet `json:"data"`
}
