package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MediaRepository struct {
	*GenericRepository[Media]
}

func NewMediaRepository(client *Client) *MediaRepository {
	return &MediaRepository{
		GenericRepository: NewGenericRepository[Media](client),
	}
}

func (t *MediaRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Media], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media")
}

func (t *MediaRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Media], *http.Response, error) {
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

func (t *MediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media")
}

func (t *MediaRepository) Upsert(ctx ApiContext, entity []Media) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media")
}

func (t *MediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media")
}

type Media struct {

	Alt      string  `json:"alt,omitempty"`

	AppPaymentMethods      []AppPaymentMethod  `json:"appPaymentMethods,omitempty"`

	AppShippingMethods      []AppShippingMethod  `json:"appShippingMethods,omitempty"`

	AvatarUsers      []User  `json:"avatarUsers,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	CmsBlocks      []CmsBlock  `json:"cmsBlocks,omitempty"`

	CmsPages      []CmsPage  `json:"cmsPages,omitempty"`

	CmsSections      []CmsSection  `json:"cmsSections,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DocumentBaseConfigs      []DocumentBaseConfig  `json:"documentBaseConfigs,omitempty"`

	Documents      []Document  `json:"documents,omitempty"`

	FileExtension      string  `json:"fileExtension,omitempty"`

	FileHash      string  `json:"fileHash,omitempty"`

	FileName      string  `json:"fileName,omitempty"`

	FileSize      float64  `json:"fileSize,omitempty"`

	HasFile      bool  `json:"hasFile,omitempty"`

	Id      string  `json:"id,omitempty"`

	MailTemplateMedia      []MailTemplateMedia  `json:"mailTemplateMedia,omitempty"`

	MediaFolder      *MediaFolder  `json:"mediaFolder,omitempty"`

	MediaFolderId      string  `json:"mediaFolderId,omitempty"`

	MediaType      interface{}  `json:"mediaType,omitempty"`

	MediaTypeRaw      interface{}  `json:"mediaTypeRaw,omitempty"`

	MetaData      interface{}  `json:"metaData,omitempty"`

	MimeType      string  `json:"mimeType,omitempty"`

	OrderLineItemDownloads      []OrderLineItemDownload  `json:"orderLineItemDownloads,omitempty"`

	OrderLineItems      []OrderLineItem  `json:"orderLineItems,omitempty"`

	Path      string  `json:"path,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	Private      bool  `json:"private,omitempty"`

	ProductConfiguratorSettings      []ProductConfiguratorSetting  `json:"productConfiguratorSettings,omitempty"`

	ProductDownloads      []ProductDownload  `json:"productDownloads,omitempty"`

	ProductManufacturers      []ProductManufacturer  `json:"productManufacturers,omitempty"`

	ProductMedia      []ProductMedia  `json:"productMedia,omitempty"`

	PropertyGroupOptions      []PropertyGroupOption  `json:"propertyGroupOptions,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	ThemeMedia      []Theme  `json:"themeMedia,omitempty"`

	Themes      []Theme  `json:"themes,omitempty"`

	Thumbnails      []MediaThumbnail  `json:"thumbnails,omitempty"`

	ThumbnailsRo      interface{}  `json:"thumbnailsRo,omitempty"`

	Title      string  `json:"title,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []MediaTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	UploadedAt      time.Time  `json:"uploadedAt,omitempty"`

	Url      string  `json:"url,omitempty"`

	User      *User  `json:"user,omitempty"`

	UserId      string  `json:"userId,omitempty"`

}
