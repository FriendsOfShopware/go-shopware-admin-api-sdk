package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductSortingRepository ClientService

func (t ProductSortingRepository) Search(ctx ApiContext, criteria Criteria) (*ProductSortingCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-sorting", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductSortingCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductSortingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductSortingCollection, *http.Response, error) {
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

func (t ProductSortingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-sorting", criteria)

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

func (t ProductSortingRepository) Upsert(ctx ApiContext, entity []ProductSorting) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_sorting": {
		Entity:  "product_sorting",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductSortingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_sorting": {
		Entity:  "product_sorting",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductSorting struct {
	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	Locked bool `json:"locked,omitempty"`

	Key string `json:"key,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	Active bool `json:"active,omitempty"`

	Fields interface{} `json:"fields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Label string `json:"label,omitempty"`

	Translations []ProductSortingTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ProductSortingCollection struct {
	EntityCollection

	Data []ProductSorting `json:"data"`
}
