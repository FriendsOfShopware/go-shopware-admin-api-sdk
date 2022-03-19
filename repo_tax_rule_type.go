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

func (t TaxRuleTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*TaxRuleTypeCollection, *http.Response, error) {
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
	Id string `json:"id,omitempty"`

	Rules []TaxRule `json:"rules,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	Position float64 `json:"position,omitempty"`

	TypeName string `json:"typeName,omitempty"`

	Translations []TaxRuleTypeTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type TaxRuleTypeCollection struct {
	EntityCollection

	Data []TaxRuleType `json:"data"`
}
