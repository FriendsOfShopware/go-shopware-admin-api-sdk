package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type StateMachineHistoryRepository struct {
	*GenericRepository[StateMachineHistory]
}

func NewStateMachineHistoryRepository(client *Client) *StateMachineHistoryRepository {
	return &StateMachineHistoryRepository{
		GenericRepository: NewGenericRepository[StateMachineHistory](client),
	}
}

func (t *StateMachineHistoryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineHistory], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "state-machine-history")
}

func (t *StateMachineHistoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineHistory], *http.Response, error) {
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

func (t *StateMachineHistoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "state-machine-history")
}

func (t *StateMachineHistoryRepository) Upsert(ctx ApiContext, entity []StateMachineHistory) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "state_machine_history")
}

func (t *StateMachineHistoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "state_machine_history")
}

type StateMachineHistory struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	EntityName      string  `json:"entityName,omitempty"`

	FromStateId      string  `json:"fromStateId,omitempty"`

	FromStateMachineState      *StateMachineState  `json:"fromStateMachineState,omitempty"`

	Id      string  `json:"id,omitempty"`

	ReferencedId      string  `json:"referencedId,omitempty"`

	ReferencedVersionId      string  `json:"referencedVersionId,omitempty"`

	StateMachine      *StateMachine  `json:"stateMachine,omitempty"`

	StateMachineId      string  `json:"stateMachineId,omitempty"`

	ToStateId      string  `json:"toStateId,omitempty"`

	ToStateMachineState      *StateMachineState  `json:"toStateMachineState,omitempty"`

	TransitionActionName      string  `json:"transitionActionName,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	User      *User  `json:"user,omitempty"`

	UserId      string  `json:"userId,omitempty"`

}
