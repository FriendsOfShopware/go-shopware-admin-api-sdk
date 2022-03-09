package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type TaxRepository ClientService

func (t TaxRepository) Search(ctx ApiContext, criteria Criteria) (*TaxCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/tax", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TaxCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t TaxRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/tax", criteria)

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

func (t TaxRepository) Upsert(ctx ApiContext, entity []Tax) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax": {
		Entity:  "tax",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t TaxRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax": {
		Entity:  "tax",
		Action:  "delete",
		Payload: payload,
	}})
}

type Tax struct {
	TaxRate float64 `json:"taxRate,omitempty"`

	Name string `json:"name,omitempty"`

	Rules []TaxRule `json:"rules,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	Id string `json:"id,omitempty"`

	Position float64 `json:"position,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Products []Product `json:"products,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type TaxCollection struct {
	EntityCollection

	Data []Tax `json:"data"`
}
