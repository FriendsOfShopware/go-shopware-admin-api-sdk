package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CurrencyTranslationRepository struct {
	*GenericRepository[CurrencyTranslation]
}

func NewCurrencyTranslationRepository(client *Client) *CurrencyTranslationRepository {
	return &CurrencyTranslationRepository{
		GenericRepository: NewGenericRepository[CurrencyTranslation](client),
	}
}

func (t *CurrencyTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CurrencyTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "currency-translation")
}

func (t *CurrencyTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CurrencyTranslation], *http.Response, error) {
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

func (t *CurrencyTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "currency-translation")
}

func (t *CurrencyTranslationRepository) Upsert(ctx ApiContext, entity []CurrencyTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "currency_translation")
}

func (t *CurrencyTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "currency_translation")
}

type CurrencyTranslation struct {

	ShortName      string  `json:"shortName,omitempty"`

	Name      string  `json:"name,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

}
