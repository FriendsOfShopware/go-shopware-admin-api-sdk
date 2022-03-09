package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MessageQueueStatsRepository ClientService

func (t MessageQueueStatsRepository) Search(ctx ApiContext, criteria Criteria) (*MessageQueueStatsCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/message-queue-stats", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MessageQueueStatsCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MessageQueueStatsRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/message-queue-stats", criteria)

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

func (t MessageQueueStatsRepository) Upsert(ctx ApiContext, entity []MessageQueueStats) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"message_queue_stats": {
		Entity:  "message_queue_stats",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MessageQueueStatsRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"message_queue_stats": {
		Entity:  "message_queue_stats",
		Action:  "delete",
		Payload: payload,
	}})
}

type MessageQueueStats struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Size float64 `json:"size,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type MessageQueueStatsCollection struct {
	EntityCollection

	Data []MessageQueueStats `json:"data"`
}
