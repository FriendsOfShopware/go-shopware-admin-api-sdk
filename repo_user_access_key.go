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
	User *User `json:"user,omitempty"`

	Id string `json:"id,omitempty"`

	UserId string `json:"userId,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	SecretAccessKey interface{} `json:"secretAccessKey,omitempty"`

	WriteAccess bool `json:"writeAccess,omitempty"`

	LastUsageAt time.Time `json:"lastUsageAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type UserAccessKeyCollection struct {
	EntityCollection

	Data []UserAccessKey `json:"data"`
}
