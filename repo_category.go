package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CategoryRepository struct {
	*GenericRepository[Category]
}

func NewCategoryRepository(client *Client) *CategoryRepository {
	return &CategoryRepository{
		GenericRepository: NewGenericRepository[Category](client),
	}
}

func (t *CategoryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Category], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "category")
}

func (t *CategoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Category], *http.Response, error) {
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

func (t *CategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "category")
}

func (t *CategoryRepository) Upsert(ctx ApiContext, entity []Category) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "category")
}

func (t *CategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "category")
}

type Category struct {

	Active      bool  `json:"active,omitempty"`

	AfterCategoryId      string  `json:"afterCategoryId,omitempty"`

	AfterCategoryVersionId      string  `json:"afterCategoryVersionId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	Breadcrumb      interface{}  `json:"breadcrumb,omitempty"`

	ChildCount      float64  `json:"childCount,omitempty"`

	Children      []Category  `json:"children,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	CmsPageIdSwitched      bool  `json:"cmsPageIdSwitched,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomEntityTypeId      string  `json:"customEntityTypeId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	DisplayNestedProducts      bool  `json:"displayNestedProducts,omitempty"`

	ExternalLink      string  `json:"externalLink,omitempty"`

	FooterSalesChannels      []SalesChannel  `json:"footerSalesChannels,omitempty"`

	Id      string  `json:"id,omitempty"`

	InternalLink      string  `json:"internalLink,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	Level      float64  `json:"level,omitempty"`

	LinkNewTab      bool  `json:"linkNewTab,omitempty"`

	LinkType      string  `json:"linkType,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	Name      string  `json:"name,omitempty"`

	NavigationSalesChannels      []SalesChannel  `json:"navigationSalesChannels,omitempty"`

	NestedProducts      []Product  `json:"nestedProducts,omitempty"`

	Parent      *Category  `json:"parent,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	ParentVersionId      string  `json:"parentVersionId,omitempty"`

	Path      string  `json:"path,omitempty"`

	ProductAssignmentType      string  `json:"productAssignmentType,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	ServiceSalesChannels      []SalesChannel  `json:"serviceSalesChannels,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []CategoryTranslation  `json:"translations,omitempty"`

	Type      string  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Visible      bool  `json:"visible,omitempty"`

	VisibleChildCount      float64  `json:"visibleChildCount,omitempty"`

}
