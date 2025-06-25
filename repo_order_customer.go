package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderCustomerRepository struct {
	*GenericRepository[OrderCustomer]
}

func NewOrderCustomerRepository(client *Client) *OrderCustomerRepository {
	return &OrderCustomerRepository{
		GenericRepository: NewGenericRepository[OrderCustomer](client),
	}
}

func (t *OrderCustomerRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderCustomer], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-customer")
}

func (t *OrderCustomerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderCustomer], *http.Response, error) {
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

func (t *OrderCustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-customer")
}

func (t *OrderCustomerRepository) Upsert(ctx ApiContext, entity []OrderCustomer) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_customer")
}

func (t *OrderCustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_customer")
}

type OrderCustomer struct {

	CustomFields      interface{}  `json:"customFields,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	Email      string  `json:"email,omitempty"`

	Company      string  `json:"company,omitempty"`

	CustomerNumber      string  `json:"customerNumber,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	Id      string  `json:"id,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	Title      string  `json:"title,omitempty"`

	VatIds      interface{}  `json:"vatIds,omitempty"`

	RemoteAddress      interface{}  `json:"remoteAddress,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Customer      *Customer  `json:"customer,omitempty"`

}
