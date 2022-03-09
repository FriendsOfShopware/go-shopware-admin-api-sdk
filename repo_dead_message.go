package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DeadMessageRepository ClientService

func (t DeadMessageRepository) Search(ctx ApiContext, criteria Criteria) (*DeadMessageCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/dead-message", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DeadMessageCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DeadMessageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/dead-message", criteria)

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

func (t DeadMessageRepository) Upsert(ctx ApiContext, entity []DeadMessage) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"dead_message": {
		Entity:  "dead_message",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DeadMessageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"dead_message": {
		Entity:  "dead_message",
		Action:  "delete",
		Payload: payload,
	}})
}

type DeadMessage struct {
	NextExecutionTime time.Time `json:"nextExecutionTime,omitempty"`

	ScheduledTaskId string `json:"scheduledTaskId,omitempty"`

	ScheduledTask *ScheduledTask `json:"scheduledTask,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	OriginalMessageClass string `json:"originalMessageClass,omitempty"`

	ErrorCount float64 `json:"errorCount,omitempty"`

	Exception string `json:"exception,omitempty"`

	ExceptionFile string `json:"exceptionFile,omitempty"`

	Id string `json:"id,omitempty"`

	ExceptionLine float64 `json:"exceptionLine,omitempty"`

	HandlerClass string `json:"handlerClass,omitempty"`

	Encrypted bool `json:"encrypted,omitempty"`

	ExceptionMessage string `json:"exceptionMessage,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	SerializedOriginalMessage interface{} `json:"serializedOriginalMessage,omitempty"`
}

type DeadMessageCollection struct {
	EntityCollection

	Data []DeadMessage `json:"data"`
}
