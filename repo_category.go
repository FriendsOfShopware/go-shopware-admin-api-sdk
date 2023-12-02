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
	Children []Category `json:"children,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	CustomEntityTypeId string `json:"customEntityTypeId,omitempty"`

	Breadcrumb interface{} `json:"breadcrumb,omitempty"`

	Products []Product `json:"products,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	DisplayNestedProducts bool `json:"displayNestedProducts,omitempty"`

	Level float64 `json:"level,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	ProductAssignmentType string `json:"productAssignmentType,omitempty"`

	Description string `json:"description,omitempty"`

	Parent *Category `json:"parent,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Type string `json:"type,omitempty"`

	LinkType string `json:"linkType,omitempty"`

	VisibleChildCount float64 `json:"visibleChildCount,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	AfterCategoryId string `json:"afterCategoryId,omitempty"`

	LinkNewTab bool `json:"linkNewTab,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	CmsPageIdSwitched bool `json:"cmsPageIdSwitched,omitempty"`

	NestedProducts []Product `json:"nestedProducts,omitempty"`

	NavigationSalesChannels []SalesChannel `json:"navigationSalesChannels,omitempty"`

	Visible bool `json:"visible,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	FooterSalesChannels []SalesChannel `json:"footerSalesChannels,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	ExternalLink string `json:"externalLink,omitempty"`

	Path string `json:"path,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	Active bool `json:"active,omitempty"`

	InternalLink string `json:"internalLink,omitempty"`

	Media *Media `json:"media,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`

	Id string `json:"id,omitempty"`

	AfterCategoryVersionId string `json:"afterCategoryVersionId,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translations []CategoryTranslation `json:"translations,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	ServiceSalesChannels []SalesChannel `json:"serviceSalesChannels,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CategoryCollection struct {
	EntityCollection

	Data []Category `json:"data"`
}
