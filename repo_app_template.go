package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type AppTemplateRepository ClientService

func (t AppTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*AppTemplateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-template", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppTemplateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppTemplateCollection, *http.Response, error) {
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

func (t AppTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-template", criteria)

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

func (t AppTemplateRepository) Upsert(ctx ApiContext, entity []AppTemplate) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_template": {
		Entity:  "app_template",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_template": {
		Entity:  "app_template",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppTemplate struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Template string `json:"template,omitempty"`

	Path string `json:"path,omitempty"`

	Active bool `json:"active,omitempty"`

	AppId string `json:"appId,omitempty"`

	App *App `json:"app,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type AppTemplateCollection struct {
	EntityCollection

	Data []AppTemplate `json:"data"`
}
