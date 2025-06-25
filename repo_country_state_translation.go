package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CountryStateTranslationRepository struct {
	*GenericRepository[CountryStateTranslation]
}

func NewCountryStateTranslationRepository(client *Client) *CountryStateTranslationRepository {
	return &CountryStateTranslationRepository{
		GenericRepository: NewGenericRepository[CountryStateTranslation](client),
	}
}

func (t *CountryStateTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CountryStateTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "country-state-translation")
}

func (t *CountryStateTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CountryStateTranslation], *http.Response, error) {
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

func (t *CountryStateTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "country-state-translation")
}

func (t *CountryStateTranslationRepository) Upsert(ctx ApiContext, entity []CountryStateTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "country_state_translation")
}

func (t *CountryStateTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "country_state_translation")
}

type CountryStateTranslation struct {

	CountryState      *CountryState  `json:"countryState,omitempty"`

	CountryStateId      string  `json:"countryStateId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
