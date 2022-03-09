package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductFeatureSetTranslationRepository ClientService

func (t ProductFeatureSetTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ProductFeatureSetTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-feature-set-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductFeatureSetTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductFeatureSetTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-feature-set-translation", criteria)

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

func (t ProductFeatureSetTranslationRepository) Upsert(ctx ApiContext, entity []ProductFeatureSetTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_feature_set_translation": {
		Entity:  "product_feature_set_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductFeatureSetTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_feature_set_translation": {
		Entity:  "product_feature_set_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductFeatureSetTranslation struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductFeatureSetId string `json:"productFeatureSetId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ProductFeatureSet *ProductFeatureSet `json:"productFeatureSet,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`
}

type ProductFeatureSetTranslationCollection struct {
	EntityCollection

	Data []ProductFeatureSetTranslation `json:"data"`
}
