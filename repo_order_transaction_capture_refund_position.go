package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderTransactionCaptureRefundPositionRepository ClientService

func (t OrderTransactionCaptureRefundPositionRepository) Search(ctx ApiContext, criteria Criteria) (*OrderTransactionCaptureRefundPositionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-transaction-capture-refund-position", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderTransactionCaptureRefundPositionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderTransactionCaptureRefundPositionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderTransactionCaptureRefundPositionCollection, *http.Response, error) {
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

func (t OrderTransactionCaptureRefundPositionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-transaction-capture-refund-position", criteria)

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

func (t OrderTransactionCaptureRefundPositionRepository) Upsert(ctx ApiContext, entity []OrderTransactionCaptureRefundPosition) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction_capture_refund_position": {
		Entity:  "order_transaction_capture_refund_position",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderTransactionCaptureRefundPositionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_transaction_capture_refund_position": {
		Entity:  "order_transaction_capture_refund_position",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderTransactionCaptureRefundPosition struct {

	Amount      interface{}  `json:"amount,omitempty"`

	RefundId      string  `json:"refundId,omitempty"`

	OrderLineItem      *OrderLineItem  `json:"orderLineItem,omitempty"`

	OrderTransactionCaptureRefund      *OrderTransactionCaptureRefund  `json:"orderTransactionCaptureRefund,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Quantity      float64  `json:"quantity,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	OrderLineItemVersionId      string  `json:"orderLineItemVersionId,omitempty"`

	ExternalReference      string  `json:"externalReference,omitempty"`

	Reason      string  `json:"reason,omitempty"`

	Id      string  `json:"id,omitempty"`

	RefundVersionId      string  `json:"refundVersionId,omitempty"`

	OrderLineItemId      string  `json:"orderLineItemId,omitempty"`

}

type OrderTransactionCaptureRefundPositionCollection struct {
	EntityCollection

	Data []OrderTransactionCaptureRefundPosition `json:"data"`
}
