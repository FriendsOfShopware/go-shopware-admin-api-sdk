package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductManufacturerRepository ClientService

func (t ProductManufacturerRepository) Search(ctx ApiContext, criteria Criteria) (*ProductManufacturerCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-manufacturer", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductManufacturerCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductManufacturerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductManufacturerCollection, *http.Response, error) {
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

func (t ProductManufacturerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-manufacturer", criteria)

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

func (t ProductManufacturerRepository) Upsert(ctx ApiContext, entity []ProductManufacturer) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_manufacturer": {
		Entity:  "product_manufacturer",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductManufacturerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_manufacturer": {
		Entity:  "product_manufacturer",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductManufacturer struct {
	Description string `json:"description,omitempty"`

	Media *Media `json:"media,omitempty"`

	Translations []ProductManufacturerTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Link string `json:"link,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Products []Product `json:"products,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type ProductManufacturerCollection struct {
	EntityCollection

	Data []ProductManufacturer `json:"data"`
}
