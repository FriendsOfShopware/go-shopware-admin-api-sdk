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
	HasFile bool `json:"hasFile,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	CmsBlocks []CmsBlock `json:"cmsBlocks,omitempty"`

	AppPaymentMethods []AppPaymentMethod `json:"appPaymentMethods,omitempty"`

	Documents []Document `json:"documents,omitempty"`

	Themes []Theme `json:"themes,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	ProductMedia []ProductMedia `json:"productMedia,omitempty"`

	MailTemplateMedia []MailTemplateMedia `json:"mailTemplateMedia,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	Id string `json:"id,omitempty"`

	MediaFolder *MediaFolder `json:"mediaFolder,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ThemeMedia []Theme `json:"themeMedia,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Title string `json:"title,omitempty"`

	Private bool `json:"private,omitempty"`

	ProductManufacturers []ProductManufacturer `json:"productManufacturers,omitempty"`

	CmsPages []CmsPage `json:"cmsPages,omitempty"`

	DocumentBaseConfigs []DocumentBaseConfig `json:"documentBaseConfigs,omitempty"`

	UserId string `json:"userId,omitempty"`

	UploadedAt time.Time `json:"uploadedAt,omitempty"`

	MediaTypeRaw interface{} `json:"mediaTypeRaw,omitempty"`

	Translations []MediaTranslation `json:"translations,omitempty"`

	PropertyGroupOptions []PropertyGroupOption `json:"propertyGroupOptions,omitempty"`

	OrderLineItems []OrderLineItem `json:"orderLineItems,omitempty"`

	CmsSections []CmsSection `json:"cmsSections,omitempty"`

	MediaFolderId string `json:"mediaFolderId,omitempty"`

	FileName string `json:"fileName,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	User *User `json:"user,omitempty"`

	Url string `json:"url,omitempty"`

	Thumbnails []MediaThumbnail `json:"thumbnails,omitempty"`

	ProductConfiguratorSettings []ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`

	MimeType string `json:"mimeType,omitempty"`

	MetaData interface{} `json:"metaData,omitempty"`

	MediaType interface{} `json:"mediaType,omitempty"`

	Alt string `json:"alt,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	FileExtension string `json:"fileExtension,omitempty"`

	FileSize float64 `json:"fileSize,omitempty"`

	ThumbnailsRo interface{} `json:"thumbnailsRo,omitempty"`

	AvatarUser *User `json:"avatarUser,omitempty"`
}

type MediaCollection struct {
	EntityCollection

	Data []Media `json:"data"`
}
