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
	Id string `json:"id,omitempty"`

	EntityName string `json:"entityName,omitempty"`

	Template string `json:"template,omitempty"`

	IsValid bool `json:"isValid,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	RouteName string `json:"routeName,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type SeoUrlTemplateCollection struct {
	EntityCollection

	Data []SeoUrlTemplate `json:"data"`
}
