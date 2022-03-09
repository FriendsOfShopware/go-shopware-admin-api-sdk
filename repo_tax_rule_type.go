package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type TaxRuleTypeRepository ClientService

func (t TaxRuleTypeRepository) Search(ctx ApiContext, criteria Criteria) (*TaxRuleTypeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/tax-rule-type", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TaxRuleTypeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t TaxRuleTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/tax-rule-type", criteria)

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

func (t TaxRuleTypeRepository) Upsert(ctx ApiContext, entity []TaxRuleType) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_rule_type": {
		Entity:  "tax_rule_type",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t TaxRuleTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_rule_type": {
		Entity:  "tax_rule_type",
		Action:  "delete",
		Payload: payload,
	}})
}

type TaxRuleType struct {
	Position float64 `json:"position,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	TypeName string `json:"typeName,omitempty"`

	Rules []TaxRule `json:"rules,omitempty"`

	Translations []TaxRuleTypeTranslation `json:"translations,omitempty"`
}

type TaxRuleTypeCollection struct {
	EntityCollection

	Data []TaxRuleType `json:"data"`
}
