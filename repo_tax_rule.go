package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type TaxRuleRepository struct {
	*GenericRepository[TaxRule]
}

func NewTaxRuleRepository(client *Client) *TaxRuleRepository {
	return &TaxRuleRepository{
		GenericRepository: NewGenericRepository[TaxRule](client),
	}
}

func (t *TaxRuleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxRule], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "tax-rule")
}

func (t *TaxRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxRule], *http.Response, error) {
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

func (t *TaxRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "tax-rule")
}

func (t *TaxRuleRepository) Upsert(ctx ApiContext, entity []TaxRule) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "tax_rule")
}

func (t *TaxRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "tax_rule")
}

type TaxRule struct {

	ActiveFrom      time.Time  `json:"activeFrom,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Data      interface{}  `json:"data,omitempty"`

	Id      string  `json:"id,omitempty"`

	Tax      *Tax  `json:"tax,omitempty"`

	TaxId      string  `json:"taxId,omitempty"`

	TaxRate      float64  `json:"taxRate,omitempty"`

	TaxRuleTypeId      string  `json:"taxRuleTypeId,omitempty"`

	Type      *TaxRuleType  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
