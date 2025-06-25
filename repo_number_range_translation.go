package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type NumberRangeTranslationRepository struct {
	*GenericRepository[NumberRangeTranslation]
}

func NewNumberRangeTranslationRepository(client *Client) *NumberRangeTranslationRepository {
	return &NumberRangeTranslationRepository{
		GenericRepository: NewGenericRepository[NumberRangeTranslation](client),
	}
}

func (t *NumberRangeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "number-range-translation")
}

func (t *NumberRangeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeTranslation], *http.Response, error) {
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

func (t *NumberRangeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "number-range-translation")
}

func (t *NumberRangeTranslationRepository) Upsert(ctx ApiContext, entity []NumberRangeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "number_range_translation")
}

func (t *NumberRangeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "number_range_translation")
}

type NumberRangeTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	NumberRange      *NumberRange  `json:"numberRange,omitempty"`

	NumberRangeId      string  `json:"numberRangeId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
