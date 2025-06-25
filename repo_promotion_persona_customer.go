package go_shopware_admin_sdk

import (
	"net/http"

)

type PromotionPersonaCustomerRepository struct {
	*GenericRepository[PromotionPersonaCustomer]
}

func NewPromotionPersonaCustomerRepository(client *Client) *PromotionPersonaCustomerRepository {
	return &PromotionPersonaCustomerRepository{
		GenericRepository: NewGenericRepository[PromotionPersonaCustomer](client),
	}
}

func (t *PromotionPersonaCustomerRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionPersonaCustomer], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-persona-customer")
}

func (t *PromotionPersonaCustomerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionPersonaCustomer], *http.Response, error) {
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

func (t *PromotionPersonaCustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-persona-customer")
}

func (t *PromotionPersonaCustomerRepository) Upsert(ctx ApiContext, entity []PromotionPersonaCustomer) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_persona_customer")
}

func (t *PromotionPersonaCustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_persona_customer")
}

type PromotionPersonaCustomer struct {

	Customer      *Customer  `json:"customer,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

}
