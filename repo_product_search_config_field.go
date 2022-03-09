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
	Ranking float64 `json:"ranking,omitempty"`

	CustomField *CustomField `json:"customField,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Tokenize bool `json:"tokenize,omitempty"`

	Searchable bool `json:"searchable,omitempty"`

	SearchConfig *ProductSearchConfig `json:"searchConfig,omitempty"`

	SearchConfigId string `json:"searchConfigId,omitempty"`

	CustomFieldId string `json:"customFieldId,omitempty"`

	Field string `json:"field,omitempty"`
}

type ProductSearchConfigFieldCollection struct {
	EntityCollection

	Data []ProductSearchConfigField `json:"data"`
}
