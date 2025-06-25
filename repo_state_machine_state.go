package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type StateMachineStateRepository struct {
	*GenericRepository[StateMachineState]
}

func NewStateMachineStateRepository(client *Client) *StateMachineStateRepository {
	return &StateMachineStateRepository{
		GenericRepository: NewGenericRepository[StateMachineState](client),
	}
}

func (t *StateMachineStateRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineState], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "state-machine-state")
}

func (t *StateMachineStateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineState], *http.Response, error) {
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

func (t *StateMachineStateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "state-machine-state")
}

func (t *StateMachineStateRepository) Upsert(ctx ApiContext, entity []StateMachineState) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "state_machine_state")
}

func (t *StateMachineStateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "state_machine_state")
}

type StateMachineState struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	FromStateMachineHistoryEntries      []StateMachineHistory  `json:"fromStateMachineHistoryEntries,omitempty"`

	FromStateMachineTransitions      []StateMachineTransition  `json:"fromStateMachineTransitions,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	OrderDeliveries      []OrderDelivery  `json:"orderDeliveries,omitempty"`

	OrderTransactionCaptureRefunds      []OrderTransactionCaptureRefund  `json:"orderTransactionCaptureRefunds,omitempty"`

	OrderTransactionCaptures      []OrderTransactionCapture  `json:"orderTransactionCaptures,omitempty"`

	OrderTransactions      []OrderTransaction  `json:"orderTransactions,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	StateMachine      *StateMachine  `json:"stateMachine,omitempty"`

	StateMachineId      string  `json:"stateMachineId,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	ToStateMachineHistoryEntries      []StateMachineHistory  `json:"toStateMachineHistoryEntries,omitempty"`

	ToStateMachineTransitions      []StateMachineTransition  `json:"toStateMachineTransitions,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []StateMachineStateTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
