package go_shopware_admin_sdk

import (
	"net/http"
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
	PromotionId string `json:"promotionId,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Customer *Customer `json:"customer,omitempty"`
}

type PromotionPersonaCustomerCollection struct {
	EntityCollection

	Data []PromotionPersonaCustomer `json:"data"`
}
