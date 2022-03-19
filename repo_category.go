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
	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	ServiceSalesChannels []SalesChannel `json:"serviceSalesChannels,omitempty"`

	Id string `json:"id,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	InternalLink string `json:"internalLink,omitempty"`

	Description string `json:"description,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	AfterCategoryVersionId string `json:"afterCategoryVersionId,omitempty"`

	DisplayNestedProducts bool `json:"displayNestedProducts,omitempty"`

	Name string `json:"name,omitempty"`

	ExternalLink string `json:"externalLink,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Level float64 `json:"level,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	LinkType string `json:"linkType,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	AfterCategoryId string `json:"afterCategoryId,omitempty"`

	Path string `json:"path,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	Type string `json:"type,omitempty"`

	ProductAssignmentType string `json:"productAssignmentType,omitempty"`

	NavigationSalesChannels []SalesChannel `json:"navigationSalesChannels,omitempty"`

	LinkNewTab bool `json:"linkNewTab,omitempty"`

	FooterSalesChannels []SalesChannel `json:"footerSalesChannels,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	Breadcrumb interface{} `json:"breadcrumb,omitempty"`

	Active bool `json:"active,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Translations []CategoryTranslation `json:"translations,omitempty"`

	Media *Media `json:"media,omitempty"`

	Products []Product `json:"products,omitempty"`

	NestedProducts []Product `json:"nestedProducts,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Visible bool `json:"visible,omitempty"`

	Parent *Category `json:"parent,omitempty"`

	Children []Category `json:"children,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CategoryCollection struct {
	EntityCollection

	Data []Category `json:"data"`
}
