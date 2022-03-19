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

func (t MessageQueueStatsRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MessageQueueStatsCollection, *http.Response, error) {
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
	Size float64 `json:"size,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type MessageQueueStatsCollection struct {
	EntityCollection

	Data []MessageQueueStats `json:"data"`
}
