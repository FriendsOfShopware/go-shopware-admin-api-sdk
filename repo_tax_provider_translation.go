package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type TaxProviderTranslationRepository struct {
	*GenericRepository[TaxProviderTranslation]
}

func NewTaxProviderTranslationRepository(client *Client) *TaxProviderTranslationRepository {
	return &TaxProviderTranslationRepository{
		GenericRepository: NewGenericRepository[TaxProviderTranslation](client),
	}
}

func (t *TaxProviderTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxProviderTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "tax-provider-translation")
}

func (t *TaxProviderTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[TaxProviderTranslation], *http.Response, error) {
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

func (t *TaxProviderTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "tax-provider-translation")
}

func (t *TaxProviderTranslationRepository) Upsert(ctx ApiContext, entity []TaxProviderTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "tax_provider_translation")
}

func (t *TaxProviderTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "tax_provider_translation")
}

type TaxProviderTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	TaxProvider      *TaxProvider  `json:"taxProvider,omitempty"`

	TaxProviderId      string  `json:"taxProviderId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
