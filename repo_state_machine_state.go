package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type StateMachineStateRepository ClientService

func (t StateMachineStateRepository) Search(ctx ApiContext, criteria Criteria) (*StateMachineStateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/state-machine-state", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(StateMachineStateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t StateMachineStateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*StateMachineStateCollection, *http.Response, error) {
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

func (t StateMachineStateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/state-machine-state", criteria)

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

func (t StateMachineStateRepository) Upsert(ctx ApiContext, entity []StateMachineState) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_state": {
		Entity:  "state_machine_state",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t StateMachineStateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_state": {
		Entity:  "state_machine_state",
		Action:  "delete",
		Payload: payload,
	}})
}

type StateMachineState struct {

	StateMachineId      string  `json:"stateMachineId,omitempty"`

	StateMachine      *StateMachine  `json:"stateMachine,omitempty"`

	FromStateMachineTransitions      []StateMachineTransition  `json:"fromStateMachineTransitions,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	Id      string  `json:"id,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	OrderTransactions      []OrderTransaction  `json:"orderTransactions,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	OrderTransactionCaptures      []OrderTransactionCapture  `json:"orderTransactionCaptures,omitempty"`

	ToStateMachineHistoryEntries      []StateMachineHistory  `json:"toStateMachineHistoryEntries,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	FromStateMachineHistoryEntries      []StateMachineHistory  `json:"fromStateMachineHistoryEntries,omitempty"`

	ToStateMachineTransitions      []StateMachineTransition  `json:"toStateMachineTransitions,omitempty"`

	Translations      []StateMachineStateTranslation  `json:"translations,omitempty"`

	OrderDeliveries      []OrderDelivery  `json:"orderDeliveries,omitempty"`

	OrderTransactionCaptureRefunds      []OrderTransactionCaptureRefund  `json:"orderTransactionCaptureRefunds,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}

type StateMachineStateCollection struct {
	EntityCollection

	Data []StateMachineState `json:"data"`
}
