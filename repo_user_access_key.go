package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type UserAccessKeyRepository ClientService

func (t UserAccessKeyRepository) Search(ctx ApiContext, criteria Criteria) (*UserAccessKeyCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/user-access-key", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(UserAccessKeyCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UserAccessKeyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*UserAccessKeyCollection, *http.Response, error) {
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

func (t UserAccessKeyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/user-access-key", criteria)

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

func (t UserAccessKeyRepository) Upsert(ctx ApiContext, entity []UserAccessKey) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user_access_key": {
		Entity:  "user_access_key",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t UserAccessKeyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user_access_key": {
		Entity:  "user_access_key",
		Action:  "delete",
		Payload: payload,
	}})
}

type UserAccessKey struct {
	Id string `json:"id,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	SecretAccessKey interface{} `json:"secretAccessKey,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UserId string `json:"userId,omitempty"`

	LastUsageAt time.Time `json:"lastUsageAt,omitempty"`

	User *User `json:"user,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type UserAccessKeyCollection struct {
	EntityCollection

	Data []UserAccessKey `json:"data"`
}
