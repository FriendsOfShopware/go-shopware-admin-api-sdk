package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type ProductSearchConfigRepository ClientService

func (t ProductSearchConfigRepository) Search(ctx ApiContext, criteria Criteria) (*ProductSearchConfigCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-search-config", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductSearchConfigCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductSearchConfigRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductSearchConfigCollection, *http.Response, error) {
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

func (t ProductSearchConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-search-config", criteria)

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

func (t ProductSearchConfigRepository) Upsert(ctx ApiContext, entity []ProductSearchConfig) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_config": {
		Entity:  "product_search_config",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductSearchConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_config": {
		Entity:  "product_search_config",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductSearchConfig struct {
	Id string `json:"id,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ExcludedTerms interface{} `json:"excludedTerms,omitempty"`

	ConfigFields []ProductSearchConfigField `json:"configFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AndLogic bool `json:"andLogic,omitempty"`

	MinSearchLength float64 `json:"minSearchLength,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type ProductSearchConfigCollection struct {
	EntityCollection

	Data []ProductSearchConfig `json:"data"`
}
