package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type LogEntryRepository ClientService

func (t LogEntryRepository) Search(ctx ApiContext, criteria Criteria) (*LogEntryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/log-entry", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LogEntryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LogEntryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/log-entry", criteria)

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

func (t LogEntryRepository) Upsert(ctx ApiContext, entity []LogEntry) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"log_entry": {
		Entity:  "log_entry",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LogEntryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"log_entry": {
		Entity:  "log_entry",
		Action:  "delete",
		Payload: payload,
	}})
}

type LogEntry struct {
	Extra interface{} `json:"extra,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	Level float64 `json:"level,omitempty"`

	Channel string `json:"channel,omitempty"`

	Context interface{} `json:"context,omitempty"`
}

type LogEntryCollection struct {
	EntityCollection

	Data []LogEntry `json:"data"`
}
