package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type NumberRangeTypeTranslationRepository struct {
	*GenericRepository[NumberRangeTypeTranslation]
}

func NewNumberRangeTypeTranslationRepository(client *Client) *NumberRangeTypeTranslationRepository {
	return &NumberRangeTypeTranslationRepository{
		GenericRepository: NewGenericRepository[NumberRangeTypeTranslation](client),
	}
}

func (t *NumberRangeTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeTypeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "number-range-type-translation")
}

func (t *NumberRangeTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeTypeTranslation], *http.Response, error) {
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

func (t *NumberRangeTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "number-range-type-translation")
}

func (t *NumberRangeTypeTranslationRepository) Upsert(ctx ApiContext, entity []NumberRangeTypeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "number_range_type_translation")
}

func (t *NumberRangeTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "number_range_type_translation")
}

type NumberRangeTypeTranslation struct {

	LanguageId      string  `json:"languageId,omitempty"`

	NumberRangeType      *NumberRangeType  `json:"numberRangeType,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	TypeName      string  `json:"typeName,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	NumberRangeTypeId      string  `json:"numberRangeTypeId,omitempty"`

}
