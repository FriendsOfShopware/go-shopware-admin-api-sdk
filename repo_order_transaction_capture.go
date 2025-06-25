package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderTransactionCaptureRepository struct {
	*GenericRepository[OrderTransactionCapture]
}

func NewOrderTransactionCaptureRepository(client *Client) *OrderTransactionCaptureRepository {
	return &OrderTransactionCaptureRepository{
		GenericRepository: NewGenericRepository[OrderTransactionCapture](client),
	}
}

func (t *OrderTransactionCaptureRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransactionCapture], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-transaction-capture")
}

func (t *OrderTransactionCaptureRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransactionCapture], *http.Response, error) {
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

func (t *OrderTransactionCaptureRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-transaction-capture")
}

func (t *OrderTransactionCaptureRepository) Upsert(ctx ApiContext, entity []OrderTransactionCapture) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_transaction_capture")
}

func (t *OrderTransactionCaptureRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_transaction_capture")
}

type OrderTransactionCapture struct {

	ExternalReference      string  `json:"externalReference,omitempty"`

	Amount      interface{}  `json:"amount,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	OrderTransactionId      string  `json:"orderTransactionId,omitempty"`

	StateId      string  `json:"stateId,omitempty"`

	Transaction      *OrderTransaction  `json:"transaction,omitempty"`

	Refunds      []OrderTransactionCaptureRefund  `json:"refunds,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	OrderTransactionVersionId      string  `json:"orderTransactionVersionId,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

}
