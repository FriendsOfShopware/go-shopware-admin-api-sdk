package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type OrderTransactionCaptureRepository ClientService

func (t OrderTransactionCaptureRepository) Search(ctx ApiContext, criteria Criteria) (*OrderTransactionCaptureCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-transaction-capture", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderTransactionCaptureCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderTransactionCaptureRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderTransactionCaptureCollection, *http.Response, error) {
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

func (t OrderTransactionCaptureRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-transaction-capture", criteria)

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

func (t OrderTransactionCaptureRepository) Upsert(ctx ApiContext, entity []OrderTransactionCapture) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction_capture": {
		Entity:  "order_transaction_capture",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderTransactionCaptureRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction_capture": {
		Entity:  "order_transaction_capture",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderTransactionCapture struct {
	StateId string `json:"stateId,omitempty"`

	Refunds []OrderTransactionCaptureRefund `json:"refunds,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`

	Transaction *OrderTransaction `json:"transaction,omitempty"`

	OrderTransactionId string `json:"orderTransactionId,omitempty"`

	OrderTransactionVersionId string `json:"orderTransactionVersionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ExternalReference string `json:"externalReference,omitempty"`

	Amount interface{} `json:"amount,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type OrderTransactionCaptureCollection struct {
	EntityCollection

	Data []OrderTransactionCapture `json:"data"`
}
