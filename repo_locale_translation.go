package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type LocaleTranslationRepository struct {
	*GenericRepository[LocaleTranslation]
}

func NewLocaleTranslationRepository(client *Client) *LocaleTranslationRepository {
	return &LocaleTranslationRepository{
		GenericRepository: NewGenericRepository[LocaleTranslation](client),
	}
}

func (t *LocaleTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[LocaleTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "locale-translation")
}

func (t *LocaleTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[LocaleTranslation], *http.Response, error) {
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

func (t *LocaleTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "locale-translation")
}

func (t *LocaleTranslationRepository) Upsert(ctx ApiContext, entity []LocaleTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "locale_translation")
}

func (t *LocaleTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "locale_translation")
}

type LocaleTranslation struct {

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	Territory      string  `json:"territory,omitempty"`

	Locale      *Locale  `json:"locale,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LocaleId      string  `json:"localeId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

}
