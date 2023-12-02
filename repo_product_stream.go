package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductStreamRepository ClientService

func (t ProductStreamRepository) Search(ctx ApiContext, criteria Criteria) (*ProductStreamCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-stream", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductStreamCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductStreamRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductStreamCollection, *http.Response, error) {
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

func (t ProductStreamRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-stream", criteria)

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

func (t ProductStreamRepository) Upsert(ctx ApiContext, entity []ProductStream) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream": {
		Entity:  "product_stream",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductStreamRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream": {
		Entity:  "product_stream",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductStream struct {

	Id      string  `json:"id,omitempty"`

	Invalid      bool  `json:"invalid,omitempty"`

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	ProductCrossSellings      []ProductCrossSelling  `json:"productCrossSellings,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Translations      []ProductStreamTranslation  `json:"translations,omitempty"`

	Filters      []ProductStreamFilter  `json:"filters,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	ApiFilter      interface{}  `json:"apiFilter,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type ProductStreamCollection struct {
	EntityCollection

	Data []ProductStream `json:"data"`
}
