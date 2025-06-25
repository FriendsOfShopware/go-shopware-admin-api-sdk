package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CountryStateRepository struct {
	*GenericRepository[CountryState]
}

func NewCountryStateRepository(client *Client) *CountryStateRepository {
	return &CountryStateRepository{
		GenericRepository: NewGenericRepository[CountryState](client),
	}
}

func (t *CountryStateRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CountryState], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "country-state")
}

func (t *CountryStateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CountryState], *http.Response, error) {
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

func (t *CountryStateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "country-state")
}

func (t *CountryStateRepository) Upsert(ctx ApiContext, entity []CountryState) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "country_state")
}

func (t *CountryStateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "country_state")
}

type CountryState struct {

	Active      bool  `json:"active,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomerAddresses      []CustomerAddress  `json:"customerAddresses,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	OrderAddresses      []OrderAddress  `json:"orderAddresses,omitempty"`

	Position      float64  `json:"position,omitempty"`

	ShortCode      string  `json:"shortCode,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []CountryStateTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
