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
	App *App `json:"app,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Template string `json:"template,omitempty"`

	Path string `json:"path,omitempty"`

	Active bool `json:"active,omitempty"`

	AppId string `json:"appId,omitempty"`
}

type AppTemplateCollection struct {
	EntityCollection

	Data []AppTemplate `json:"data"`
}
