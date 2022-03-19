package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SeoUrlTemplateRepository ClientService

func (t SeoUrlTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*SeoUrlTemplateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/seo-url-template", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SeoUrlTemplateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SeoUrlTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SeoUrlTemplateCollection, *http.Response, error) {
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

func (t SeoUrlTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/seo-url-template", criteria)

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

func (t SeoUrlTemplateRepository) Upsert(ctx ApiContext, entity []SeoUrlTemplate) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"seo_url_template": {
		Entity:  "seo_url_template",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SeoUrlTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"seo_url_template": {
		Entity:  "seo_url_template",
		Action:  "delete",
		Payload: payload,
	}})
}

type SeoUrlTemplate struct {
	SalesChannelId string `json:"salesChannelId,omitempty"`

	EntityName string `json:"entityName,omitempty"`

	RouteName string `json:"routeName,omitempty"`

	Template string `json:"template,omitempty"`

	IsValid bool `json:"isValid,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Id string `json:"id,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type SeoUrlTemplateCollection struct {
	EntityCollection

	Data []SeoUrlTemplate `json:"data"`
}
