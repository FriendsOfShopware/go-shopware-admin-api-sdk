package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppActionButtonTranslationRepository struct {
	*GenericRepository[AppActionButtonTranslation]
}

func NewAppActionButtonTranslationRepository(client *Client) *AppActionButtonTranslationRepository {
	return &AppActionButtonTranslationRepository{
		GenericRepository: NewGenericRepository[AppActionButtonTranslation](client),
	}
}

func (t *AppActionButtonTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppActionButtonTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-action-button-translation")
}

func (t *AppActionButtonTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppActionButtonTranslation], *http.Response, error) {
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

func (t *AppActionButtonTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-action-button-translation")
}

func (t *AppActionButtonTranslationRepository) Upsert(ctx ApiContext, entity []AppActionButtonTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_action_button_translation")
}

func (t *AppActionButtonTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_action_button_translation")
}

type AppActionButtonTranslation struct {

	Label      string  `json:"label,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	AppActionButtonId      string  `json:"appActionButtonId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	AppActionButton      *AppActionButton  `json:"appActionButton,omitempty"`

	Language      *Language  `json:"language,omitempty"`

}
