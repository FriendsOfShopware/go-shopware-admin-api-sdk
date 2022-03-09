package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PromotionTranslationRepository ClientService

func (t PromotionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-translation", criteria)

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

func (t PromotionTranslationRepository) Upsert(ctx ApiContext, entity []PromotionTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_translation": {
		Entity:  "promotion_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_translation": {
		Entity:  "promotion_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionTranslation struct {
	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type PromotionTranslationCollection struct {
	EntityCollection

	Data []PromotionTranslation `json:"data"`
}
