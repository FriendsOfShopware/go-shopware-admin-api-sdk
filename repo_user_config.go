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

func (t UserConfigRepository) SearchAll(ctx ApiContext, criteria Criteria) (*UserConfigCollection, *http.Response, error) {
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
	Id string `json:"id,omitempty"`

	UserId string `json:"userId,omitempty"`

	Key string `json:"key,omitempty"`

	Value interface{} `json:"value,omitempty"`

	User *User `json:"user,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type UserConfigCollection struct {
	EntityCollection

	Data []UserConfig `json:"data"`
}
