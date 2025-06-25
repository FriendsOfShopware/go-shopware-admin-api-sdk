package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SeoUrlTemplateRepository struct {
	*GenericRepository[SeoUrlTemplate]
}

func NewSeoUrlTemplateRepository(client *Client) *SeoUrlTemplateRepository {
	return &SeoUrlTemplateRepository{
		GenericRepository: NewGenericRepository[SeoUrlTemplate](client),
	}
}

func (t *SeoUrlTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SeoUrlTemplate], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "seo-url-template")
}

func (t *SeoUrlTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SeoUrlTemplate], *http.Response, error) {
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

func (t *SeoUrlTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "seo-url-template")
}

func (t *SeoUrlTemplateRepository) Upsert(ctx ApiContext, entity []SeoUrlTemplate) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "seo_url_template")
}

func (t *SeoUrlTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "seo_url_template")
}

type SeoUrlTemplate struct {

	Id      string  `json:"id,omitempty"`

	EntityName      string  `json:"entityName,omitempty"`

	RouteName      string  `json:"routeName,omitempty"`

	Template      string  `json:"template,omitempty"`

	IsValid      bool  `json:"isValid,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}
