package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomerRecoveryRepository struct {
	*GenericRepository[CustomerRecovery]
}

func NewCustomerRecoveryRepository(client *Client) *CustomerRecoveryRepository {
	return &CustomerRecoveryRepository{
		GenericRepository: NewGenericRepository[CustomerRecovery](client),
	}
}

func (t *CustomerRecoveryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerRecovery], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-recovery")
}

func (t *CustomerRecoveryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerRecovery], *http.Response, error) {
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

func (t *CustomerRecoveryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-recovery")
}

func (t *CustomerRecoveryRepository) Upsert(ctx ApiContext, entity []CustomerRecovery) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_recovery")
}

func (t *CustomerRecoveryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_recovery")
}

type CustomerRecovery struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Customer      *Customer  `json:"customer,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	Id      string  `json:"id,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
