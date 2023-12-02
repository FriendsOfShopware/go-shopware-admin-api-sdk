package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type FlowTemplateRepository ClientService

func (t FlowTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*FlowTemplateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/flow-template", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(FlowTemplateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t FlowTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*FlowTemplateCollection, *http.Response, error) {
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

func (t FlowTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/flow-template", criteria)

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

func (t FlowTemplateRepository) Upsert(ctx ApiContext, entity []FlowTemplate) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow_template": {
		Entity:  "flow_template",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t FlowTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow_template": {
		Entity:  "flow_template",
		Action:  "delete",
		Payload: payload,
	}})
}

type FlowTemplate struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type FlowTemplateCollection struct {
	EntityCollection

	Data []FlowTemplate `json:"data"`
}
