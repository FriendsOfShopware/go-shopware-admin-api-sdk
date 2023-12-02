package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type AppScriptConditionRepository ClientService

func (t AppScriptConditionRepository) Search(ctx ApiContext, criteria Criteria) (*AppScriptConditionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-script-condition", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppScriptConditionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppScriptConditionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppScriptConditionCollection, *http.Response, error) {
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

func (t AppScriptConditionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-script-condition", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppScriptConditionRepository) Upsert(ctx ApiContext, entity []AppScriptCondition) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_script_condition": {
		Entity:  "app_script_condition",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppScriptConditionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_script_condition": {
		Entity:  "app_script_condition",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppScriptCondition struct {
	RuleConditions []RuleCondition `json:"ruleConditions,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	App *App `json:"app,omitempty"`

	Config interface{} `json:"config,omitempty"`

	AppId string `json:"appId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Group string `json:"group,omitempty"`

	Script string `json:"script,omitempty"`

	Constraints interface{} `json:"constraints,omitempty"`

	Active bool `json:"active,omitempty"`

	Translations []AppScriptConditionTranslation `json:"translations,omitempty"`

	Identifier string `json:"identifier,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type AppScriptConditionCollection struct {
	EntityCollection

	Data []AppScriptCondition `json:"data"`
}
