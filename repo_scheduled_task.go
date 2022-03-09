package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ScheduledTaskRepository ClientService

func (t ScheduledTaskRepository) Search(ctx ApiContext, criteria Criteria) (*ScheduledTaskCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/scheduled-task", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ScheduledTaskCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ScheduledTaskRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/scheduled-task", criteria)

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

func (t ScheduledTaskRepository) Upsert(ctx ApiContext, entity []ScheduledTask) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"scheduled_task": {
		Entity:  "scheduled_task",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ScheduledTaskRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"scheduled_task": {
		Entity:  "scheduled_task",
		Action:  "delete",
		Payload: payload,
	}})
}

type ScheduledTask struct {
	Id string `json:"id,omitempty"`

	RunInterval float64 `json:"runInterval,omitempty"`

	Status string `json:"status,omitempty"`

	NextExecutionTime time.Time `json:"nextExecutionTime,omitempty"`

	DeadMessages []DeadMessage `json:"deadMessages,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Name string `json:"name,omitempty"`

	ScheduledTaskClass string `json:"scheduledTaskClass,omitempty"`

	LastExecutionTime time.Time `json:"lastExecutionTime,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ScheduledTaskCollection struct {
	EntityCollection

	Data []ScheduledTask `json:"data"`
}
