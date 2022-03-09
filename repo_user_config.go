package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type UserConfigRepository ClientService

func (t UserConfigRepository) Search(ctx ApiContext, criteria Criteria) (*UserConfigCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/user-config", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(UserConfigCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UserConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/user-config", criteria)

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

func (t UserConfigRepository) Upsert(ctx ApiContext, entity []UserConfig) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user_config": {
		Entity:  "user_config",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t UserConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user_config": {
		Entity:  "user_config",
		Action:  "delete",
		Payload: payload,
	}})
}

type UserConfig struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	UserId string `json:"userId,omitempty"`

	Key string `json:"key,omitempty"`

	Value interface{} `json:"value,omitempty"`

	User *User `json:"user,omitempty"`
}

type UserConfigCollection struct {
	EntityCollection

	Data []UserConfig `json:"data"`
}
