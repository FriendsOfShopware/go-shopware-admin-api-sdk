package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderTransactionCaptureRefundRepository struct {
	*GenericRepository[OrderTransactionCaptureRefund]
}

func NewOrderTransactionCaptureRefundRepository(client *Client) *OrderTransactionCaptureRefundRepository {
	return &OrderTransactionCaptureRefundRepository{
		GenericRepository: NewGenericRepository[OrderTransactionCaptureRefund](client),
	}
}

func (t *OrderTransactionCaptureRefundRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransactionCaptureRefund], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-transaction-capture-refund")
}

func (t *OrderTransactionCaptureRefundRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransactionCaptureRefund], *http.Response, error) {
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

func (t *OrderTransactionCaptureRefundRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-transaction-capture-refund")
}

func (t *OrderTransactionCaptureRefundRepository) Upsert(ctx ApiContext, entity []OrderTransactionCaptureRefund) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_transaction_capture_refund")
}

func (t *OrderTransactionCaptureRefundRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_transaction_capture_refund")
}

type OrderTransactionCaptureRefund struct {

	TransactionCapture      *OrderTransactionCapture  `json:"transactionCapture,omitempty"`

	Reason      string  `json:"reason,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	CaptureId      string  `json:"captureId,omitempty"`

	CaptureVersionId      string  `json:"captureVersionId,omitempty"`

	Positions      []OrderTransactionCaptureRefundPosition  `json:"positions,omitempty"`

	ExternalReference      string  `json:"externalReference,omitempty"`

	Amount      interface{}  `json:"amount,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	StateId      string  `json:"stateId,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

}
