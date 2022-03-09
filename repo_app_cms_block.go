package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppCmsBlockRepository ClientService

func (t AppCmsBlockRepository) Search(ctx ApiContext, criteria Criteria) (*AppCmsBlockCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-cms-block", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppCmsBlockCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppCmsBlockRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-cms-block", criteria)

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

func (t AppCmsBlockRepository) Upsert(ctx ApiContext, entity []AppCmsBlock) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_cms_block": {
		Entity:  "app_cms_block",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppCmsBlockRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_cms_block": {
		Entity:  "app_cms_block",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppCmsBlock struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Block interface{} `json:"block,omitempty"`

	Template string `json:"template,omitempty"`

	Styles string `json:"styles,omitempty"`

	AppId string `json:"appId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Label string `json:"label,omitempty"`

	Translations []AppCmsBlockTranslation `json:"translations,omitempty"`

	App *App `json:"app,omitempty"`
}

type AppCmsBlockCollection struct {
	EntityCollection

	Data []AppCmsBlock `json:"data"`
}
