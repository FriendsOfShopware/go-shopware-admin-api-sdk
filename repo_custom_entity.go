package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type CustomEntityRepository ClientService

func (t CustomEntityRepository) Search(ctx ApiContext, criteria Criteria) (*CustomEntityCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/custom-entity", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomEntityCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomEntityRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomEntityCollection, *http.Response, error) {
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

func (t CustomEntityRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/custom-entity", criteria)

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

func (t CustomEntityRepository) Upsert(ctx ApiContext, entity []CustomEntity) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_entity": {
		Entity:  "custom_entity",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomEntityRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_entity": {
		Entity:  "custom_entity",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomEntity struct {
	StoreApiAware bool `json:"storeApiAware,omitempty"`

	DeletedAt time.Time `json:"deletedAt,omitempty"`

	Name string `json:"name,omitempty"`

	Flags interface{} `json:"flags,omitempty"`

	AppId string `json:"appId,omitempty"`

	CustomFieldsAware bool `json:"customFieldsAware,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Fields interface{} `json:"fields,omitempty"`

	CmsAware bool `json:"cmsAware,omitempty"`

	Id string `json:"id,omitempty"`

	PluginId string `json:"pluginId,omitempty"`

	LabelProperty string `json:"labelProperty,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CustomEntityCollection struct {
	EntityCollection

	Data []CustomEntity `json:"data"`
}
