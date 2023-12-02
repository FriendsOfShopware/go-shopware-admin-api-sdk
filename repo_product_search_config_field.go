package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type ProductSearchConfigFieldRepository ClientService

func (t ProductSearchConfigFieldRepository) Search(ctx ApiContext, criteria Criteria) (*ProductSearchConfigFieldCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-search-config-field", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductSearchConfigFieldCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductSearchConfigFieldRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductSearchConfigFieldCollection, *http.Response, error) {
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

func (t ProductSearchConfigFieldRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-search-config-field", criteria)

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

func (t ProductSearchConfigFieldRepository) Upsert(ctx ApiContext, entity []ProductSearchConfigField) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_config_field": {
		Entity:  "product_search_config_field",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductSearchConfigFieldRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_config_field": {
		Entity:  "product_search_config_field",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductSearchConfigField struct {
	CustomField *CustomField `json:"customField,omitempty"`

	Id string `json:"id,omitempty"`

	CustomFieldId string `json:"customFieldId,omitempty"`

	Tokenize bool `json:"tokenize,omitempty"`

	Searchable bool `json:"searchable,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	SearchConfigId string `json:"searchConfigId,omitempty"`

	Field string `json:"field,omitempty"`

	Ranking float64 `json:"ranking,omitempty"`

	SearchConfig *ProductSearchConfig `json:"searchConfig,omitempty"`
}

type ProductSearchConfigFieldCollection struct {
	EntityCollection

	Data []ProductSearchConfigField `json:"data"`
}
