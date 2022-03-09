package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PromotionIndividualCodeRepository ClientService

func (t PromotionIndividualCodeRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionIndividualCodeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-individual-code", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionIndividualCodeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionIndividualCodeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-individual-code", criteria)

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

func (t PromotionIndividualCodeRepository) Upsert(ctx ApiContext, entity []PromotionIndividualCode) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_individual_code": {
		Entity:  "promotion_individual_code",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionIndividualCodeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_individual_code": {
		Entity:  "promotion_individual_code",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionIndividualCode struct {
	Code string `json:"code,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`
}

type PromotionIndividualCodeCollection struct {
	EntityCollection

	Data []PromotionIndividualCode `json:"data"`
}
