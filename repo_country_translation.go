package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CountryTranslationRepository struct {
	*GenericRepository[CountryTranslation]
}

func NewCountryTranslationRepository(client *Client) *CountryTranslationRepository {
	return &CountryTranslationRepository{
		GenericRepository: NewGenericRepository[CountryTranslation](client),
	}
}

func (t *CountryTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CountryTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "country-translation")
}

func (t *CountryTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CountryTranslation], *http.Response, error) {
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

func (t *CountryTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "country-translation")
}

func (t *CountryTranslationRepository) Upsert(ctx ApiContext, entity []CountryTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "country_translation")
}

func (t *CountryTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "country_translation")
}

type CountryTranslation struct {

	AddressFormat      interface{}  `json:"addressFormat,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Country      *Country  `json:"country,omitempty"`

}
