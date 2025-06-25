package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SeoUrlRepository struct {
	*GenericRepository[SeoUrl]
}

func NewSeoUrlRepository(client *Client) *SeoUrlRepository {
	return &SeoUrlRepository{
		GenericRepository: NewGenericRepository[SeoUrl](client),
	}
}

func (t *SeoUrlRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SeoUrl], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "seo-url")
}

func (t *SeoUrlRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SeoUrl], *http.Response, error) {
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

func (t *SeoUrlRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "seo-url")
}

func (t *SeoUrlRepository) Upsert(ctx ApiContext, entity []SeoUrl) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "seo_url")
}

func (t *SeoUrlRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "seo_url")
}

type SeoUrl struct {

	Id      string  `json:"id,omitempty"`

	PathInfo      string  `json:"pathInfo,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ForeignKey      string  `json:"foreignKey,omitempty"`

	IsModified      bool  `json:"isModified,omitempty"`

	IsDeleted      bool  `json:"isDeleted,omitempty"`

	Url      string  `json:"url,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SeoPathInfo      string  `json:"seoPathInfo,omitempty"`

	IsCanonical      bool  `json:"isCanonical,omitempty"`

	Error      string  `json:"error,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	RouteName      string  `json:"routeName,omitempty"`

}
