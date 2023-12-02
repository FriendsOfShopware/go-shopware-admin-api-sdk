package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppFlowActionTranslationRepository ClientService

func (t AppFlowActionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*AppFlowActionTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-flow-action-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppFlowActionTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppFlowActionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppFlowActionTranslationCollection, *http.Response, error) {
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

func (t AppFlowActionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-flow-action-translation", criteria)

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

func (t AppFlowActionTranslationRepository) Upsert(ctx ApiContext, entity []AppFlowActionTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_flow_action_translation": {
		Entity:  "app_flow_action_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppFlowActionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_flow_action_translation": {
		Entity:  "app_flow_action_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppFlowActionTranslation struct {

	Language      *Language  `json:"language,omitempty"`

	Description      string  `json:"description,omitempty"`

	Headline      string  `json:"headline,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	AppFlowActionId      string  `json:"appFlowActionId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Label      string  `json:"label,omitempty"`

	AppFlowAction      *AppFlowAction  `json:"appFlowAction,omitempty"`

}

type AppFlowActionTranslationCollection struct {
	EntityCollection

	Data []AppFlowActionTranslation `json:"data"`
}
