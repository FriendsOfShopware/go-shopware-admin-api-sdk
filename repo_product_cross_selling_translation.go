package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type ProductCrossSellingTranslationRepository ClientService

func (t ProductCrossSellingTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCrossSellingTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-cross-selling-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCrossSellingTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductCrossSellingTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCrossSellingTranslationCollection, *http.Response, error) {
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

func (t ProductCrossSellingTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-cross-selling-translation", criteria)

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

func (t ProductCrossSellingTranslationRepository) Upsert(ctx ApiContext, entity []ProductCrossSellingTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_cross_selling_translation": {
		Entity:  "product_cross_selling_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductCrossSellingTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_cross_selling_translation": {
		Entity:  "product_cross_selling_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductCrossSellingTranslation struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductCrossSellingId string `json:"productCrossSellingId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ProductCrossSelling *ProductCrossSelling `json:"productCrossSelling,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`
}

type ProductCrossSellingTranslationCollection struct {
	EntityCollection

	Data []ProductCrossSellingTranslation `json:"data"`
}
