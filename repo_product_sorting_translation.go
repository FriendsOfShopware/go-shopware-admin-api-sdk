package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductSortingTranslationRepository ClientService

func (t ProductSortingTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ProductSortingTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-sorting-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductSortingTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductSortingTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductSortingTranslationCollection, *http.Response, error) {
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

func (t ProductSortingTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-sorting-translation", criteria)

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

func (t ProductSortingTranslationRepository) Upsert(ctx ApiContext, entity []ProductSortingTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_sorting_translation": {
		Entity:  "product_sorting_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductSortingTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_sorting_translation": {
		Entity:  "product_sorting_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductSortingTranslation struct {

	ProductSortingId      string  `json:"productSortingId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ProductSorting      *ProductSorting  `json:"productSorting,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Label      string  `json:"label,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type ProductSortingTranslationCollection struct {
	EntityCollection

	Data []ProductSortingTranslation `json:"data"`
}
