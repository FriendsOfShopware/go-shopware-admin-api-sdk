package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppTranslationRepository struct {
	*GenericRepository[AppTranslation]
}

func NewAppTranslationRepository(client *Client) *AppTranslationRepository {
	return &AppTranslationRepository{
		GenericRepository: NewGenericRepository[AppTranslation](client),
	}
}

func (t *AppTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-translation")
}

func (t *AppTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppTranslation], *http.Response, error) {
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

func (t *AppTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-translation")
}

func (t *AppTranslationRepository) Upsert(ctx ApiContext, entity []AppTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_translation")
}

func (t *AppTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_translation")
}

type AppTranslation struct {

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Label      string  `json:"label,omitempty"`

	PrivacyPolicyExtensions      string  `json:"privacyPolicyExtensions,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	App      *App  `json:"app,omitempty"`

	Description      string  `json:"description,omitempty"`

}
