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
	Invalid bool `json:"invalid,omitempty"`

	ProductCrossSellings []ProductCrossSelling `json:"productCrossSellings,omitempty"`

	ProductExports []ProductExport `json:"productExports,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	ApiFilter interface{} `json:"apiFilter,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Translations []ProductStreamTranslation `json:"translations,omitempty"`

	Filters []ProductStreamFilter `json:"filters,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Categories []Category `json:"categories,omitempty"`
}

type ProductStreamCollection struct {
	EntityCollection

	Data []ProductStream `json:"data"`
}
