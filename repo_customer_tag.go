package go_shopware_admin_sdk

import (
	"net/http"

)

type CustomerTagRepository struct {
	*GenericRepository[CustomerTag]
}

func NewCustomerTagRepository(client *Client) *CustomerTagRepository {
	return &CustomerTagRepository{
		GenericRepository: NewGenericRepository[CustomerTag](client),
	}
}

func (t *CustomerTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-tag")
}

func (t *CustomerTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerTag], *http.Response, error) {
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

func (t *CustomerTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-tag")
}

func (t *CustomerTagRepository) Upsert(ctx ApiContext, entity []CustomerTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_tag")
}

func (t *CustomerTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_tag")
}

type CustomerTag struct {

	Customer      *Customer  `json:"customer,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

}
