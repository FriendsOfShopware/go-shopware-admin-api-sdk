package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CategoryRepository ClientService

func (t CategoryRepository) Search(ctx ApiContext, criteria Criteria) (*CategoryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/category", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CategoryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CategoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CategoryCollection, *http.Response, error) {
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

func (t CategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/category", criteria)

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

func (t CategoryRepository) Upsert(ctx ApiContext, entity []Category) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"category": {
		Entity:  "category",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"category": {
		Entity:  "category",
		Action:  "delete",
		Payload: payload,
	}})
}

type Category struct {

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	NestedProducts      []Product  `json:"nestedProducts,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Active      bool  `json:"active,omitempty"`

	LinkNewTab      bool  `json:"linkNewTab,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Breadcrumb      interface{}  `json:"breadcrumb,omitempty"`

	Visible      bool  `json:"visible,omitempty"`

	VisibleChildCount      float64  `json:"visibleChildCount,omitempty"`

	Children      []Category  `json:"children,omitempty"`

	Id      string  `json:"id,omitempty"`

	Type      string  `json:"type,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	ChildCount      float64  `json:"childCount,omitempty"`

	CmsPageIdSwitched      bool  `json:"cmsPageIdSwitched,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	AfterCategoryVersionId      string  `json:"afterCategoryVersionId,omitempty"`

	AfterCategoryId      string  `json:"afterCategoryId,omitempty"`

	ExternalLink      string  `json:"externalLink,omitempty"`

	CustomEntityTypeId      string  `json:"customEntityTypeId,omitempty"`

	NavigationSalesChannels      []SalesChannel  `json:"navigationSalesChannels,omitempty"`

	Path      string  `json:"path,omitempty"`

	LinkType      string  `json:"linkType,omitempty"`

	Description      string  `json:"description,omitempty"`

	Parent      *Category  `json:"parent,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	FooterSalesChannels      []SalesChannel  `json:"footerSalesChannels,omitempty"`

	ProductAssignmentType      string  `json:"productAssignmentType,omitempty"`

	Translations      []CategoryTranslation  `json:"translations,omitempty"`

	ParentVersionId      string  `json:"parentVersionId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	DisplayNestedProducts      bool  `json:"displayNestedProducts,omitempty"`

	Name      string  `json:"name,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Level      float64  `json:"level,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	InternalLink      string  `json:"internalLink,omitempty"`

	ServiceSalesChannels      []SalesChannel  `json:"serviceSalesChannels,omitempty"`

}

type CategoryCollection struct {
	EntityCollection

	Data []Category `json:"data"`
}
