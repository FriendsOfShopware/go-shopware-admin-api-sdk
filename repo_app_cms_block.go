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

func (t AppCmsBlockRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppCmsBlockCollection, *http.Response, error) {
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

	Name      string  `json:"name,omitempty"`

	Translations      []AppCmsBlockTranslation  `json:"translations,omitempty"`

	App      *App  `json:"app,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	Block      interface{}  `json:"block,omitempty"`

	Template      string  `json:"template,omitempty"`

	Styles      string  `json:"styles,omitempty"`

	Label      string  `json:"label,omitempty"`

}

type AppCmsBlockCollection struct {
	EntityCollection

	Data []AppCmsBlock `json:"data"`
}
