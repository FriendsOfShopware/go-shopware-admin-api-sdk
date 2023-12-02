package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PromotionPersonaCustomerRepository ClientService

func (t PromotionPersonaCustomerRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionPersonaCustomerCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-persona-customer", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionPersonaCustomerCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionPersonaCustomerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionPersonaCustomerCollection, *http.Response, error) {
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

func (t PromotionPersonaCustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-persona-customer", criteria)

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

func (t PromotionPersonaCustomerRepository) Upsert(ctx ApiContext, entity []PromotionPersonaCustomer) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_persona_customer": {
		Entity:  "promotion_persona_customer",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionPersonaCustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_persona_customer": {
		Entity:  "promotion_persona_customer",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionPersonaCustomer struct {

	PromotionId      string  `json:"promotionId,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	Customer      *Customer  `json:"customer,omitempty"`

}

type PromotionPersonaCustomerCollection struct {
	EntityCollection

	Data []PromotionPersonaCustomer `json:"data"`
}
