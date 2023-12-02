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

func (t ProductTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductTranslationCollection, *http.Response, error) {
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
	CustomFields interface{} `json:"customFields,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Product *Product `json:"product,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Description string `json:"description,omitempty"`

	CustomSearchKeywords interface{} `json:"customSearchKeywords,omitempty"`

	ProductId string `json:"productId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Language *Language `json:"language,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	PackUnitPlural string `json:"packUnitPlural,omitempty"`

	PackUnit string `json:"packUnit,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Name string `json:"name,omitempty"`

	Keywords string `json:"keywords,omitempty"`
}

type ProductTranslationCollection struct {
	EntityCollection

	Data []ProductTranslation `json:"data"`
}
