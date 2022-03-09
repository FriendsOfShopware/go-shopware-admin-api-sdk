package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SeoUrlRepository ClientService

func (t SeoUrlRepository) Search(ctx ApiContext, criteria Criteria) (*SeoUrlCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/seo-url", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SeoUrlCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SeoUrlRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/seo-url", criteria)

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

func (t SeoUrlRepository) Upsert(ctx ApiContext, entity []SeoUrl) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"seo_url": {
		Entity:  "seo_url",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SeoUrlRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"seo_url": {
		Entity:  "seo_url",
		Action:  "delete",
		Payload: payload,
	}})
}

type SeoUrl struct {
	IsCanonical bool `json:"isCanonical,omitempty"`

	Url string `json:"url,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	SeoPathInfo string `json:"seoPathInfo,omitempty"`

	PathInfo string `json:"pathInfo,omitempty"`

	IsDeleted bool `json:"isDeleted,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	IsModified bool `json:"isModified,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	ForeignKey string `json:"foreignKey,omitempty"`

	RouteName string `json:"routeName,omitempty"`

	Language *Language `json:"language,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type SeoUrlCollection struct {
	EntityCollection

	Data []SeoUrl `json:"data"`
}
