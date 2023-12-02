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
	CmsPages []CmsPage `json:"cmsPages,omitempty"`

	ProductMedia []ProductMedia `json:"productMedia,omitempty"`

	AvatarUsers []User `json:"avatarUsers,omitempty"`

	CmsSections []CmsSection `json:"cmsSections,omitempty"`

	Url string `json:"url,omitempty"`

	ThemeMedia []Theme `json:"themeMedia,omitempty"`

	MediaType interface{} `json:"mediaType,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UserId string `json:"userId,omitempty"`

	ProductConfiguratorSettings []ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`

	OrderLineItemDownloads []OrderLineItemDownload `json:"orderLineItemDownloads,omitempty"`

	AppShippingMethods []AppShippingMethod `json:"appShippingMethods,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	HasFile bool `json:"hasFile,omitempty"`

	Thumbnails []MediaThumbnail `json:"thumbnails,omitempty"`

	Themes []Theme `json:"themes,omitempty"`

	MediaFolderId string `json:"mediaFolderId,omitempty"`

	MediaFolder *MediaFolder `json:"mediaFolder,omitempty"`

	Id string `json:"id,omitempty"`

	MediaTypeRaw interface{} `json:"mediaTypeRaw,omitempty"`

	Translations []MediaTranslation `json:"translations,omitempty"`

	ProductDownloads []ProductDownload `json:"productDownloads,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	FileExtension string `json:"fileExtension,omitempty"`

	UploadedAt time.Time `json:"uploadedAt,omitempty"`

	DocumentBaseConfigs []DocumentBaseConfig `json:"documentBaseConfigs,omitempty"`

	ThumbnailsRo interface{} `json:"thumbnailsRo,omitempty"`

	MimeType string `json:"mimeType,omitempty"`

	Config interface{} `json:"config,omitempty"`

	Title string `json:"title,omitempty"`

	AppPaymentMethods []AppPaymentMethod `json:"appPaymentMethods,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	OrderLineItems []OrderLineItem `json:"orderLineItems,omitempty"`

	CmsBlocks []CmsBlock `json:"cmsBlocks,omitempty"`

	FileName string `json:"fileName,omitempty"`

	Alt string `json:"alt,omitempty"`

	Private bool `json:"private,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	MetaData interface{} `json:"metaData,omitempty"`

	ProductManufacturers []ProductManufacturer `json:"productManufacturers,omitempty"`

	MailTemplateMedia []MailTemplateMedia `json:"mailTemplateMedia,omitempty"`

	Documents []Document `json:"documents,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	FileSize float64 `json:"fileSize,omitempty"`

	Path string `json:"path,omitempty"`

	User *User `json:"user,omitempty"`

	PropertyGroupOptions []PropertyGroupOption `json:"propertyGroupOptions,omitempty"`
}

type MediaCollection struct {
	EntityCollection

	Data []Media `json:"data"`
}
