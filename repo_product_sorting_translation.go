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
	LanguageId string `json:"languageId,omitempty"`

	ProductSorting *ProductSorting `json:"productSorting,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductSortingId string `json:"productSortingId,omitempty"`
}

type ProductSortingTranslationCollection struct {
	EntityCollection

	Data []ProductSortingTranslation `json:"data"`
}
