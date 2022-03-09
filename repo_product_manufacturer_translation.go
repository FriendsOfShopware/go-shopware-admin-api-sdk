package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductManufacturerTranslationRepository ClientService

func (t ProductManufacturerTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ProductManufacturerTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-manufacturer-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductManufacturerTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductManufacturerTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-manufacturer-translation", criteria)

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

func (t ProductManufacturerTranslationRepository) Upsert(ctx ApiContext, entity []ProductManufacturerTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_manufacturer_translation": {
		Entity:  "product_manufacturer_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductManufacturerTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_manufacturer_translation": {
		Entity:  "product_manufacturer_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductManufacturerTranslation struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductManufacturer *ProductManufacturer `json:"productManufacturer,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	ProductManufacturerVersionId string `json:"productManufacturerVersionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ProductManufacturerId string `json:"productManufacturerId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`
}

type ProductManufacturerTranslationCollection struct {
	EntityCollection

	Data []ProductManufacturerTranslation `json:"data"`
}
