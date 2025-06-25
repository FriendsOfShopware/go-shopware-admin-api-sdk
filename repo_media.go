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

	AppShippingMethods      []AppShippingMethod  `json:"appShippingMethods,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ProductManufacturers      []ProductManufacturer  `json:"productManufacturers,omitempty"`

	MailTemplateMedia      []MailTemplateMedia  `json:"mailTemplateMedia,omitempty"`

	Themes      []Theme  `json:"themes,omitempty"`

	MediaFolderId      string  `json:"mediaFolderId,omitempty"`

	FileSize      float64  `json:"fileSize,omitempty"`

	MetaData      interface{}  `json:"metaData,omitempty"`

	Alt      string  `json:"alt,omitempty"`

	Path      string  `json:"path,omitempty"`

	HasFile      bool  `json:"hasFile,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	ThumbnailsRo      interface{}  `json:"thumbnailsRo,omitempty"`

	Id      string  `json:"id,omitempty"`

	Thumbnails      []MediaThumbnail  `json:"thumbnails,omitempty"`

	PropertyGroupOptions      []PropertyGroupOption  `json:"propertyGroupOptions,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	ProductConfiguratorSettings      []ProductConfiguratorSetting  `json:"productConfiguratorSettings,omitempty"`

	OrderLineItems      []OrderLineItem  `json:"orderLineItems,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	FileName      string  `json:"fileName,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	ProductDownloads      []ProductDownload  `json:"productDownloads,omitempty"`

	OrderLineItemDownloads      []OrderLineItemDownload  `json:"orderLineItemDownloads,omitempty"`

	DocumentBaseConfigs      []DocumentBaseConfig  `json:"documentBaseConfigs,omitempty"`

	CmsPages      []CmsPage  `json:"cmsPages,omitempty"`

	AppPaymentMethods      []AppPaymentMethod  `json:"appPaymentMethods,omitempty"`

	ThemeMedia      []Theme  `json:"themeMedia,omitempty"`

	MediaTypeRaw      interface{}  `json:"mediaTypeRaw,omitempty"`

	User      *User  `json:"user,omitempty"`

	CmsBlocks      []CmsBlock  `json:"cmsBlocks,omitempty"`

	Documents      []Document  `json:"documents,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	MediaType      interface{}  `json:"mediaType,omitempty"`

	Url      string  `json:"url,omitempty"`

	Private      bool  `json:"private,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	ProductMedia      []ProductMedia  `json:"productMedia,omitempty"`

	AvatarUsers      []User  `json:"avatarUsers,omitempty"`

	MediaFolder      *MediaFolder  `json:"mediaFolder,omitempty"`

	UploadedAt      time.Time  `json:"uploadedAt,omitempty"`

	Translations      []MediaTranslation  `json:"translations,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	CmsSections      []CmsSection  `json:"cmsSections,omitempty"`

	FileHash      string  `json:"fileHash,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	MimeType      string  `json:"mimeType,omitempty"`

	FileExtension      string  `json:"fileExtension,omitempty"`

	Title      string  `json:"title,omitempty"`

}
