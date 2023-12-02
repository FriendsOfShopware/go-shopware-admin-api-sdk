package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MediaRepository ClientService

func (t MediaRepository) Search(ctx ApiContext, criteria Criteria) (*MediaCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaCollection, *http.Response, error) {
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

func (t MediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media", criteria)

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

func (t MediaRepository) Upsert(ctx ApiContext, entity []Media) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media": {
		Entity:  "media",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media": {
		Entity:  "media",
		Action:  "delete",
		Payload: payload,
	}})
}

type Media struct {

	Path      string  `json:"path,omitempty"`

	Thumbnails      []MediaThumbnail  `json:"thumbnails,omitempty"`

	CmsPages      []CmsPage  `json:"cmsPages,omitempty"`

	Alt      string  `json:"alt,omitempty"`

	User      *User  `json:"user,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	MetaData      interface{}  `json:"metaData,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	MediaTypeRaw      interface{}  `json:"mediaTypeRaw,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	MediaFolder      *MediaFolder  `json:"mediaFolder,omitempty"`

	OrderLineItems      []OrderLineItem  `json:"orderLineItems,omitempty"`

	Documents      []Document  `json:"documents,omitempty"`

	ProductManufacturers      []ProductManufacturer  `json:"productManufacturers,omitempty"`

	ProductMedia      []ProductMedia  `json:"productMedia,omitempty"`

	MediaFolderId      string  `json:"mediaFolderId,omitempty"`

	Private      bool  `json:"private,omitempty"`

	Translations      []MediaTranslation  `json:"translations,omitempty"`

	ProductDownloads      []ProductDownload  `json:"productDownloads,omitempty"`

	DocumentBaseConfigs      []DocumentBaseConfig  `json:"documentBaseConfigs,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Themes      []Theme  `json:"themes,omitempty"`

	UploadedAt      time.Time  `json:"uploadedAt,omitempty"`

	Title      string  `json:"title,omitempty"`

	ThumbnailsRo      interface{}  `json:"thumbnailsRo,omitempty"`

	ProductConfiguratorSettings      []ProductConfiguratorSetting  `json:"productConfiguratorSettings,omitempty"`

	ThemeMedia      []Theme  `json:"themeMedia,omitempty"`

	MediaType      interface{}  `json:"mediaType,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	HasFile      bool  `json:"hasFile,omitempty"`

	OrderLineItemDownloads      []OrderLineItemDownload  `json:"orderLineItemDownloads,omitempty"`

	CmsSections      []CmsSection  `json:"cmsSections,omitempty"`

	Id      string  `json:"id,omitempty"`

	FileExtension      string  `json:"fileExtension,omitempty"`

	FileName      string  `json:"fileName,omitempty"`

	MailTemplateMedia      []MailTemplateMedia  `json:"mailTemplateMedia,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	Url      string  `json:"url,omitempty"`

	AvatarUsers      []User  `json:"avatarUsers,omitempty"`

	PropertyGroupOptions      []PropertyGroupOption  `json:"propertyGroupOptions,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	MimeType      string  `json:"mimeType,omitempty"`

	FileSize      float64  `json:"fileSize,omitempty"`

	CmsBlocks      []CmsBlock  `json:"cmsBlocks,omitempty"`

	AppPaymentMethods      []AppPaymentMethod  `json:"appPaymentMethods,omitempty"`

	AppShippingMethods      []AppShippingMethod  `json:"appShippingMethods,omitempty"`

}

type MediaCollection struct {
	EntityCollection

	Data []Media `json:"data"`
}
