package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type TaxProviderRepository struct {
	*GenericRepository[TaxProvider]
}

func NewTaxProviderRepository(client *Client) *TaxProviderRepository {
	return &TaxProviderRepository{
		GenericRepository: NewGenericRepository[TaxProvider](client),
	}
}

func (t *TaxProviderRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxProvider], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "tax-provider")
}

func (t *TaxProviderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxProvider], *http.Response, error) {
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

func (t *TaxProviderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "tax-provider")
}

func (t *TaxProviderRepository) Upsert(ctx ApiContext, entity []TaxProvider) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "tax_provider")
}

func (t *TaxProviderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "tax_provider")
}

type TaxProvider struct {

	ProcessUrl      string  `json:"processUrl,omitempty"`

	AvailabilityRuleId      string  `json:"availabilityRuleId,omitempty"`

	App      *App  `json:"app,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	AvailabilityRule      *Rule  `json:"availabilityRule,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Identifier      string  `json:"identifier,omitempty"`

	Translations      []TaxProviderTranslation  `json:"translations,omitempty"`

	Priority      float64  `json:"priority,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Name      string  `json:"name,omitempty"`

}
