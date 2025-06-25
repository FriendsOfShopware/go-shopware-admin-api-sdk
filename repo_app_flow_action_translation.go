package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppFlowActionTranslationRepository struct {
	*GenericRepository[AppFlowActionTranslation]
}

func NewAppFlowActionTranslationRepository(client *Client) *AppFlowActionTranslationRepository {
	return &AppFlowActionTranslationRepository{
		GenericRepository: NewGenericRepository[AppFlowActionTranslation](client),
	}
}

func (t *AppFlowActionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppFlowActionTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-flow-action-translation")
}

func (t *AppFlowActionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppFlowActionTranslation], *http.Response, error) {
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

func (t *AppFlowActionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-flow-action-translation")
}

func (t *AppFlowActionTranslationRepository) Upsert(ctx ApiContext, entity []AppFlowActionTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_flow_action_translation")
}

func (t *AppFlowActionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_flow_action_translation")
}

type AppFlowActionTranslation struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	AppFlowActionId      string  `json:"appFlowActionId,omitempty"`

	Label      string  `json:"label,omitempty"`

	Headline      string  `json:"headline,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	AppFlowAction      *AppFlowAction  `json:"appFlowAction,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Description      string  `json:"description,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}
