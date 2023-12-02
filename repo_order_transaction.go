package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderTransactionRepository ClientService

func (t OrderTransactionRepository) Search(ctx ApiContext, criteria Criteria) (*OrderTransactionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-transaction", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderTransactionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderTransactionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderTransactionCollection, *http.Response, error) {
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

func (t OrderTransactionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-transaction", criteria)

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

func (t OrderTransactionRepository) Upsert(ctx ApiContext, entity []OrderTransaction) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction": {
		Entity:  "order_transaction",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderTransactionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction": {
		Entity:  "order_transaction",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderTransaction struct {

	Id      string  `json:"id,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	Amount      interface{}  `json:"amount,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

	Captures      []OrderTransactionCapture  `json:"captures,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	StateId      string  `json:"stateId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type OrderTransactionCollection struct {
	EntityCollection

	Data []OrderTransaction `json:"data"`
}
