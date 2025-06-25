package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalutationTranslationRepository struct {
	*GenericRepository[SalutationTranslation]
}

func NewSalutationTranslationRepository(client *Client) *SalutationTranslationRepository {
	return &SalutationTranslationRepository{
		GenericRepository: NewGenericRepository[SalutationTranslation](client),
	}
}

func (t *SalutationTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalutationTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "salutation-translation")
}

func (t *SalutationTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalutationTranslation], *http.Response, error) {
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

func (t *SalutationTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "salutation-translation")
}

func (t *SalutationTranslationRepository) Upsert(ctx ApiContext, entity []SalutationTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "salutation_translation")
}

func (t *SalutationTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "salutation_translation")
}

type SalutationTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DisplayName      string  `json:"displayName,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	LetterName      string  `json:"letterName,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
