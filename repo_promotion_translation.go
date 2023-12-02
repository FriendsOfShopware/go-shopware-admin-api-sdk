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

func (t PromotionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionTranslationCollection, *http.Response, error) {
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
	LanguageId string `json:"languageId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`
}

type PromotionTranslationCollection struct {
	EntityCollection

	Data []PromotionTranslation `json:"data"`
}
