package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type TaxRuleTypeRepository struct {
	*GenericRepository[TaxRuleType]
}

func NewTaxRuleTypeRepository(client *Client) *TaxRuleTypeRepository {
	return &TaxRuleTypeRepository{
		GenericRepository: NewGenericRepository[TaxRuleType](client),
	}
}

func (t *TaxRuleTypeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxRuleType], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "tax-rule-type")
}

func (t *TaxRuleTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxRuleType], *http.Response, error) {
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

func (t *TaxRuleTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "tax-rule-type")
}

func (t *TaxRuleTypeRepository) Upsert(ctx ApiContext, entity []TaxRuleType) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "tax_rule_type")
}

func (t *TaxRuleTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "tax_rule_type")
}

type TaxRuleType struct {

	Id      string  `json:"id,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Translations      []TaxRuleTypeTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	TypeName      string  `json:"typeName,omitempty"`

	Rules      []TaxRule  `json:"rules,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

}
