package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductTranslationRepository ClientService

func (t ProductTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ProductTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-translation", criteria)

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

func (t ProductTranslationRepository) Upsert(ctx ApiContext, entity []ProductTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_translation": {
		Entity:  "product_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_translation": {
		Entity:  "product_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductTranslation struct {
	PackUnit string `json:"packUnit,omitempty"`

	PackUnitPlural string `json:"packUnitPlural,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	Description string `json:"description,omitempty"`

	Language *Language `json:"language,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	Product *Product `json:"product,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	CustomSearchKeywords interface{} `json:"customSearchKeywords,omitempty"`

	ProductId string `json:"productId,omitempty"`
}

type ProductTranslationCollection struct {
	EntityCollection

	Data []ProductTranslation `json:"data"`
}
