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
	ChildCount float64 `json:"childCount,omitempty"`

	Type string `json:"type,omitempty"`

	Active bool `json:"active,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	Translations []CategoryTranslation `json:"translations,omitempty"`

	AfterCategoryId string `json:"afterCategoryId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ServiceSalesChannels []SalesChannel `json:"serviceSalesChannels,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	Parent *Category `json:"parent,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	FooterSalesChannels []SalesChannel `json:"footerSalesChannels,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	Visible bool `json:"visible,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	LinkType string `json:"linkType,omitempty"`

	ExternalLink string `json:"externalLink,omitempty"`

	LinkNewTab bool `json:"linkNewTab,omitempty"`

	Children []Category `json:"children,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	DisplayNestedProducts bool `json:"displayNestedProducts,omitempty"`

	Breadcrumb interface{} `json:"breadcrumb,omitempty"`

	ProductAssignmentType string `json:"productAssignmentType,omitempty"`

	InternalLink string `json:"internalLink,omitempty"`

	Media *Media `json:"media,omitempty"`

	Products []Product `json:"products,omitempty"`

	NavigationSalesChannels []SalesChannel `json:"navigationSalesChannels,omitempty"`

	Id string `json:"id,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`

	AfterCategoryVersionId string `json:"afterCategoryVersionId,omitempty"`

	Level float64 `json:"level,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	NestedProducts []Product `json:"nestedProducts,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Path string `json:"path,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`

	VersionId string `json:"versionId,omitempty"`
}

type CategoryCollection struct {
	EntityCollection

	Data []Category `json:"data"`
}
