package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ThemeChildRepository ClientService

func (t ThemeChildRepository) Search(ctx ApiContext, criteria Criteria) (*ThemeChildCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/theme-child", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ThemeChildCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ThemeChildRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ThemeChildCollection, *http.Response, error) {
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

func (t ThemeChildRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/theme-child", criteria)

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

func (t ThemeChildRepository) Upsert(ctx ApiContext, entity []ThemeChild) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_child": {
		Entity:  "theme_child",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ThemeChildRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_child": {
		Entity:  "theme_child",
		Action:  "delete",
		Payload: payload,
	}})
}

type ThemeChild struct {

	ParentId      string  `json:"parentId,omitempty"`

	ChildId      string  `json:"childId,omitempty"`

	ParentTheme      *Theme  `json:"parentTheme,omitempty"`

	ChildTheme      *Theme  `json:"childTheme,omitempty"`

}

type ThemeChildCollection struct {
	EntityCollection

	Data []ThemeChild `json:"data"`
}
