package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type OrderTransactionCaptureRefundRepository ClientService

func (t OrderTransactionCaptureRefundRepository) Search(ctx ApiContext, criteria Criteria) (*OrderTransactionCaptureRefundCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-transaction-capture-refund", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderTransactionCaptureRefundCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderTransactionCaptureRefundRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderTransactionCaptureRefundCollection, *http.Response, error) {
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

func (t OrderTransactionCaptureRefundRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-transaction-capture-refund", criteria)

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

func (t OrderTransactionCaptureRefundRepository) Upsert(ctx ApiContext, entity []OrderTransactionCaptureRefund) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction_capture_refund": {
		Entity:  "order_transaction_capture_refund",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderTransactionCaptureRefundRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction_capture_refund": {
		Entity:  "order_transaction_capture_refund",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderTransactionCaptureRefund struct {
	StateId string `json:"stateId,omitempty"`

	Reason string `json:"reason,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	CaptureId string `json:"captureId,omitempty"`

	TransactionCapture *OrderTransactionCapture `json:"transactionCapture,omitempty"`

	ExternalReference string `json:"externalReference,omitempty"`

	Amount interface{} `json:"amount,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CaptureVersionId string `json:"captureVersionId,omitempty"`

	Positions []OrderTransactionCaptureRefundPosition `json:"positions,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`
}

type OrderTransactionCaptureRefundCollection struct {
	EntityCollection

	Data []OrderTransactionCaptureRefund `json:"data"`
}
