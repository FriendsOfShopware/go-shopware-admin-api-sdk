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

	Products      []Product  `json:"products,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	ParentVersionId      string  `json:"parentVersionId,omitempty"`

	AfterCategoryId      string  `json:"afterCategoryId,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Breadcrumb      interface{}  `json:"breadcrumb,omitempty"`

	Level      float64  `json:"level,omitempty"`

	Path      string  `json:"path,omitempty"`

	Active      bool  `json:"active,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	ProductAssignmentType      string  `json:"productAssignmentType,omitempty"`

	LinkType      string  `json:"linkType,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	Visible      bool  `json:"visible,omitempty"`

	VisibleChildCount      float64  `json:"visibleChildCount,omitempty"`

	Parent      *Category  `json:"parent,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	CustomEntityTypeId      string  `json:"customEntityTypeId,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	FooterSalesChannels      []SalesChannel  `json:"footerSalesChannels,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ExternalLink      string  `json:"externalLink,omitempty"`

	Type      string  `json:"type,omitempty"`

	Translations      []CategoryTranslation  `json:"translations,omitempty"`

	NestedProducts      []Product  `json:"nestedProducts,omitempty"`

	NavigationSalesChannels      []SalesChannel  `json:"navigationSalesChannels,omitempty"`

	ServiceSalesChannels      []SalesChannel  `json:"serviceSalesChannels,omitempty"`

	Id      string  `json:"id,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	CmsPageIdSwitched      bool  `json:"cmsPageIdSwitched,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	Children      []Category  `json:"children,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	AfterCategoryVersionId      string  `json:"afterCategoryVersionId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	DisplayNestedProducts      bool  `json:"displayNestedProducts,omitempty"`

	ChildCount      float64  `json:"childCount,omitempty"`

	Name      string  `json:"name,omitempty"`

	InternalLink      string  `json:"internalLink,omitempty"`

	LinkNewTab      bool  `json:"linkNewTab,omitempty"`

}
