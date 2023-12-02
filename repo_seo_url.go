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

func (t SeoUrlRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SeoUrlCollection, *http.Response, error) {
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
	Language *Language `json:"language,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	ForeignKey string `json:"foreignKey,omitempty"`

	RouteName string `json:"routeName,omitempty"`

	PathInfo string `json:"pathInfo,omitempty"`

	SeoPathInfo string `json:"seoPathInfo,omitempty"`

	IsModified bool `json:"isModified,omitempty"`

	IsValid bool `json:"isValid,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	IsDeleted bool `json:"isDeleted,omitempty"`

	Url string `json:"url,omitempty"`

	Id string `json:"id,omitempty"`

	IsCanonical bool `json:"isCanonical,omitempty"`
}

type SeoUrlCollection struct {
	EntityCollection

	Data []SeoUrl `json:"data"`
}
