package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PropertyGroupTranslationRepository struct {
	*GenericRepository[PropertyGroupTranslation]
}

func NewPropertyGroupTranslationRepository(client *Client) *PropertyGroupTranslationRepository {
	return &PropertyGroupTranslationRepository{
		GenericRepository: NewGenericRepository[PropertyGroupTranslation](client),
	}
}

func (t *PropertyGroupTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroupTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "property-group-translation")
}

func (t *PropertyGroupTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroupTranslation], *http.Response, error) {
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

func (t *PropertyGroupTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "property-group-translation")
}

func (t *PropertyGroupTranslationRepository) Upsert(ctx ApiContext, entity []PropertyGroupTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "property_group_translation")
}

func (t *PropertyGroupTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "property_group_translation")
}

type PropertyGroupTranslation struct {

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	Position      float64  `json:"position,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	PropertyGroupId      string  `json:"propertyGroupId,omitempty"`

	PropertyGroup      *PropertyGroup  `json:"propertyGroup,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

}
