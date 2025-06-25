package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type TaxRuleTypeTranslationRepository struct {
	*GenericRepository[TaxRuleTypeTranslation]
}

func NewTaxRuleTypeTranslationRepository(client *Client) *TaxRuleTypeTranslationRepository {
	return &TaxRuleTypeTranslationRepository{
		GenericRepository: NewGenericRepository[TaxRuleTypeTranslation](client),
	}
}

func (t *TaxRuleTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxRuleTypeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "tax-rule-type-translation")
}

func (t *TaxRuleTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxRuleTypeTranslation], *http.Response, error) {
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

func (t *TaxRuleTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "tax-rule-type-translation")
}

func (t *TaxRuleTypeTranslationRepository) Upsert(ctx ApiContext, entity []TaxRuleTypeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "tax_rule_type_translation")
}

func (t *TaxRuleTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "tax_rule_type_translation")
}

type TaxRuleTypeTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	TaxRuleType      *TaxRuleType  `json:"taxRuleType,omitempty"`

	TaxRuleTypeId      string  `json:"taxRuleTypeId,omitempty"`

	TypeName      string  `json:"typeName,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
