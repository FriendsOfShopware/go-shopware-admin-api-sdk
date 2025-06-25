package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderTransactionRepository struct {
	*GenericRepository[OrderTransaction]
}

func NewOrderTransactionRepository(client *Client) *OrderTransactionRepository {
	return &OrderTransactionRepository{
		GenericRepository: NewGenericRepository[OrderTransaction](client),
	}
}

func (t *OrderTransactionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransaction], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-transaction")
}

func (t *OrderTransactionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTransaction], *http.Response, error) {
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

func (t *OrderTransactionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-transaction")
}

func (t *OrderTransactionRepository) Upsert(ctx ApiContext, entity []OrderTransaction) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_transaction")
}

func (t *OrderTransactionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_transaction")
}

type OrderTransaction struct {

	Amount      interface{}  `json:"amount,omitempty"`

	Captures      []OrderTransactionCapture  `json:"captures,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	StateId      string  `json:"stateId,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ValidationData      interface{}  `json:"validationData,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}
