package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type TaxRuleRepository ClientService

func (t TaxRuleRepository) Search(ctx ApiContext, criteria Criteria) (*TaxRuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/tax-rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TaxRuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t TaxRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/tax-rule", criteria)

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

func (t TaxRuleRepository) Upsert(ctx ApiContext, entity []TaxRule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_rule": {
		Entity:  "tax_rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t TaxRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_rule": {
		Entity:  "tax_rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type TaxRule struct {
	TaxId string `json:"taxId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	TaxRate float64 `json:"taxRate,omitempty"`

	TaxRuleTypeId string `json:"taxRuleTypeId,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	Data interface{} `json:"data,omitempty"`

	Type *TaxRuleType `json:"type,omitempty"`

	Country *Country `json:"country,omitempty"`

	Tax *Tax `json:"tax,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type TaxRuleCollection struct {
	EntityCollection

	Data []TaxRule `json:"data"`
}
