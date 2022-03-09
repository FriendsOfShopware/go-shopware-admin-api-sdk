package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SystemConfigRepository ClientService

func (t SystemConfigRepository) Search(ctx ApiContext, criteria Criteria) (*SystemConfigCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/system-config", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SystemConfigCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SystemConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/system-config", criteria)

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

func (t SystemConfigRepository) Upsert(ctx ApiContext, entity []SystemConfig) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"system_config": {
		Entity:  "system_config",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SystemConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"system_config": {
		Entity:  "system_config",
		Action:  "delete",
		Payload: payload,
	}})
}

type SystemConfig struct {
	ConfigurationKey string `json:"configurationKey,omitempty"`

	ConfigurationValue interface{} `json:"configurationValue,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type SystemConfigCollection struct {
	EntityCollection

	Data []SystemConfig `json:"data"`
}
