package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type StateMachineTransitionRepository struct {
	*GenericRepository[StateMachineTransition]
}

func NewStateMachineTransitionRepository(client *Client) *StateMachineTransitionRepository {
	return &StateMachineTransitionRepository{
		GenericRepository: NewGenericRepository[StateMachineTransition](client),
	}
}

func (t *StateMachineTransitionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineTransition], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "state-machine-transition")
}

func (t *StateMachineTransitionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineTransition], *http.Response, error) {
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

func (t *StateMachineTransitionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "state-machine-transition")
}

func (t *StateMachineTransitionRepository) Upsert(ctx ApiContext, entity []StateMachineTransition) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "state_machine_transition")
}

func (t *StateMachineTransitionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "state_machine_transition")
}

type StateMachineTransition struct {

	ActionName      string  `json:"actionName,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	FromStateId      string  `json:"fromStateId,omitempty"`

	FromStateMachineState      *StateMachineState  `json:"fromStateMachineState,omitempty"`

	Id      string  `json:"id,omitempty"`

	StateMachine      *StateMachine  `json:"stateMachine,omitempty"`

	StateMachineId      string  `json:"stateMachineId,omitempty"`

	ToStateId      string  `json:"toStateId,omitempty"`

	ToStateMachineState      *StateMachineState  `json:"toStateMachineState,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
