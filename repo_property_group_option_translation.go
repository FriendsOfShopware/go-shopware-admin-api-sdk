package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PropertyGroupOptionTranslationRepository struct {
	*GenericRepository[PropertyGroupOptionTranslation]
}

func NewPropertyGroupOptionTranslationRepository(client *Client) *PropertyGroupOptionTranslationRepository {
	return &PropertyGroupOptionTranslationRepository{
		GenericRepository: NewGenericRepository[PropertyGroupOptionTranslation](client),
	}
}

func (t *PropertyGroupOptionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroupOptionTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "property-group-option-translation")
}

func (t *PropertyGroupOptionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroupOptionTranslation], *http.Response, error) {
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

func (t *PropertyGroupOptionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "property-group-option-translation")
}

func (t *PropertyGroupOptionTranslationRepository) Upsert(ctx ApiContext, entity []PropertyGroupOptionTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "property_group_option_translation")
}

func (t *PropertyGroupOptionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "property_group_option_translation")
}

type PropertyGroupOptionTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	PropertyGroupOption      *PropertyGroupOption  `json:"propertyGroupOption,omitempty"`

	PropertyGroupOptionId      string  `json:"propertyGroupOptionId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
