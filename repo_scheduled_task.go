package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ScheduledTaskRepository struct {
	*GenericRepository[ScheduledTask]
}

func NewScheduledTaskRepository(client *Client) *ScheduledTaskRepository {
	return &ScheduledTaskRepository{
		GenericRepository: NewGenericRepository[ScheduledTask](client),
	}
}

func (t *ScheduledTaskRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ScheduledTask], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "scheduled-task")
}

func (t *ScheduledTaskRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ScheduledTask], *http.Response, error) {
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

func (t *ScheduledTaskRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "scheduled-task")
}

func (t *ScheduledTaskRepository) Upsert(ctx ApiContext, entity []ScheduledTask) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "scheduled_task")
}

func (t *ScheduledTaskRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "scheduled_task")
}

type ScheduledTask struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	RunInterval      float64  `json:"runInterval,omitempty"`

	DefaultRunInterval      float64  `json:"defaultRunInterval,omitempty"`

	LastExecutionTime      time.Time  `json:"lastExecutionTime,omitempty"`

	NextExecutionTime      time.Time  `json:"nextExecutionTime,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ScheduledTaskClass      string  `json:"scheduledTaskClass,omitempty"`

	Status      string  `json:"status,omitempty"`

}
