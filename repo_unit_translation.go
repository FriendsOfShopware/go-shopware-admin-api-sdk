package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type UnitTranslationRepository struct {
	*GenericRepository[UnitTranslation]
}

func NewUnitTranslationRepository(client *Client) *UnitTranslationRepository {
	return &UnitTranslationRepository{
		GenericRepository: NewGenericRepository[UnitTranslation](client),
	}
}

func (t *UnitTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[UnitTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "unit-translation")
}

func (t *UnitTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[UnitTranslation], *http.Response, error) {
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

func (t *UnitTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "unit-translation")
}

func (t *UnitTranslationRepository) Upsert(ctx ApiContext, entity []UnitTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "unit_translation")
}

func (t *UnitTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "unit_translation")
}

type UnitTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	ShortCode      string  `json:"shortCode,omitempty"`

	Unit      *Unit  `json:"unit,omitempty"`

	UnitId      string  `json:"unitId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
