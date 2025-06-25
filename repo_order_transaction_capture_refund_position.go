package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderTransactionCaptureRefundPositionRepository struct {
	*GenericRepository[OrderTransactionCaptureRefundPosition]
}

func NewOrderTransactionCaptureRefundPositionRepository(client *Client) *OrderTransactionCaptureRefundPositionRepository {
	return &OrderTransactionCaptureRefundPositionRepository{
		GenericRepository: NewGenericRepository[OrderTransactionCaptureRefundPosition](client),
	}
}

func (t *OrderTransactionCaptureRefundPositionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransactionCaptureRefundPosition], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-transaction-capture-refund-position")
}

func (t *OrderTransactionCaptureRefundPositionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransactionCaptureRefundPosition], *http.Response, error) {
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

func (t *OrderTransactionCaptureRefundPositionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-transaction-capture-refund-position")
}

func (t *OrderTransactionCaptureRefundPositionRepository) Upsert(ctx ApiContext, entity []OrderTransactionCaptureRefundPosition) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_transaction_capture_refund_position")
}

func (t *OrderTransactionCaptureRefundPositionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_transaction_capture_refund_position")
}

type OrderTransactionCaptureRefundPosition struct {

	Amount      interface{}  `json:"amount,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	ExternalReference      string  `json:"externalReference,omitempty"`

	Id      string  `json:"id,omitempty"`

	OrderLineItem      *OrderLineItem  `json:"orderLineItem,omitempty"`

	OrderLineItemId      string  `json:"orderLineItemId,omitempty"`

	OrderLineItemVersionId      string  `json:"orderLineItemVersionId,omitempty"`

	OrderTransactionCaptureRefund      *OrderTransactionCaptureRefund  `json:"orderTransactionCaptureRefund,omitempty"`

	Quantity      float64  `json:"quantity,omitempty"`

	Reason      string  `json:"reason,omitempty"`

	RefundId      string  `json:"refundId,omitempty"`

	RefundVersionId      string  `json:"refundVersionId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}
