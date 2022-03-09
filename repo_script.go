package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ScriptRepository ClientService

func (t ScriptRepository) Search(ctx ApiContext, criteria Criteria) (*ScriptCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/script", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ScriptCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ScriptRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/script", criteria)

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

func (t ScriptRepository) Upsert(ctx ApiContext, entity []Script) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"script": {
		Entity:  "script",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ScriptRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"script": {
		Entity:  "script",
		Action:  "delete",
		Payload: payload,
	}})
}

type Script struct {
	Id string `json:"id,omitempty"`

	AppId string `json:"appId,omitempty"`

	App *App `json:"app,omitempty"`

	Script string `json:"script,omitempty"`

	Hook string `json:"hook,omitempty"`

	Name string `json:"name,omitempty"`

	Active bool `json:"active,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ScriptCollection struct {
	EntityCollection

	Data []Script `json:"data"`
}
