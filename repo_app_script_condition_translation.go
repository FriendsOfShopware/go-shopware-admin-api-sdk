package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppScriptConditionTranslationRepository struct {
	*GenericRepository[AppScriptConditionTranslation]
}

func NewAppScriptConditionTranslationRepository(client *Client) *AppScriptConditionTranslationRepository {
	return &AppScriptConditionTranslationRepository{
		GenericRepository: NewGenericRepository[AppScriptConditionTranslation](client),
	}
}

func (t *AppScriptConditionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppScriptConditionTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-script-condition-translation")
}

func (t *AppScriptConditionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppScriptConditionTranslation], *http.Response, error) {
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

func (t *AppScriptConditionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-script-condition-translation")
}

func (t *AppScriptConditionTranslationRepository) Upsert(ctx ApiContext, entity []AppScriptConditionTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_script_condition_translation")
}

func (t *AppScriptConditionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_script_condition_translation")
}

type AppScriptConditionTranslation struct {

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	AppScriptConditionId      string  `json:"appScriptConditionId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	AppScriptCondition      *AppScriptCondition  `json:"appScriptCondition,omitempty"`

}
