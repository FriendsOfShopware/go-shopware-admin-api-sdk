package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppActionButtonRepository ClientService

func (t AppActionButtonRepository) Search(ctx ApiContext, criteria Criteria) (*AppActionButtonCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-action-button", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppActionButtonCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppActionButtonRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppActionButtonCollection, *http.Response, error) {
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

func (t AppActionButtonRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-action-button", criteria)

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

func (t AppActionButtonRepository) Upsert(ctx ApiContext, entity []AppActionButton) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_action_button": {
		Entity:  "app_action_button",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppActionButtonRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_action_button": {
		Entity:  "app_action_button",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppActionButton struct {
	Translations []AppActionButtonTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	View string `json:"view,omitempty"`

	Label string `json:"label,omitempty"`

	Url string `json:"url,omitempty"`

	Action string `json:"action,omitempty"`

	OpenNewTab bool `json:"openNewTab,omitempty"`

	AppId string `json:"appId,omitempty"`

	App *App `json:"app,omitempty"`

	Id string `json:"id,omitempty"`

	Entity string `json:"entity,omitempty"`
}

type AppActionButtonCollection struct {
	EntityCollection

	Data []AppActionButton `json:"data"`
}
