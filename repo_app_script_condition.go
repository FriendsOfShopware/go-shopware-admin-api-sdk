package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppScriptConditionRepository struct {
	*GenericRepository[AppScriptCondition]
}

func NewAppScriptConditionRepository(client *Client) *AppScriptConditionRepository {
	return &AppScriptConditionRepository{
		GenericRepository: NewGenericRepository[AppScriptCondition](client),
	}
}

func (t *AppScriptConditionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppScriptCondition], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-script-condition")
}

func (t *AppScriptConditionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppScriptCondition], *http.Response, error) {
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

func (t *AppScriptConditionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-script-condition")
}

func (t *AppScriptConditionRepository) Upsert(ctx ApiContext, entity []AppScriptCondition) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_script_condition")
}

func (t *AppScriptConditionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_script_condition")
}

type AppScriptCondition struct {

	Active      bool  `json:"active,omitempty"`

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	Constraints      interface{}  `json:"constraints,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Group      string  `json:"group,omitempty"`

	Id      string  `json:"id,omitempty"`

	Identifier      string  `json:"identifier,omitempty"`

	Name      string  `json:"name,omitempty"`

	RuleConditions      []RuleCondition  `json:"ruleConditions,omitempty"`

	Script      string  `json:"script,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []AppScriptConditionTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
