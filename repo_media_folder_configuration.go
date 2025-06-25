package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MediaFolderConfigurationRepository struct {
	*GenericRepository[MediaFolderConfiguration]
}

func NewMediaFolderConfigurationRepository(client *Client) *MediaFolderConfigurationRepository {
	return &MediaFolderConfigurationRepository{
		GenericRepository: NewGenericRepository[MediaFolderConfiguration](client),
	}
}

func (t *MediaFolderConfigurationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaFolderConfiguration], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-folder-configuration")
}

func (t *MediaFolderConfigurationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaFolderConfiguration], *http.Response, error) {
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

func (t *MediaFolderConfigurationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-folder-configuration")
}

func (t *MediaFolderConfigurationRepository) Upsert(ctx ApiContext, entity []MediaFolderConfiguration) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_folder_configuration")
}

func (t *MediaFolderConfigurationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_folder_configuration")
}

type MediaFolderConfiguration struct {

	NoAssociation      bool  `json:"noAssociation,omitempty"`

	MediaThumbnailSizes      []MediaThumbnailSize  `json:"mediaThumbnailSizes,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	MediaFolders      []MediaFolder  `json:"mediaFolders,omitempty"`

	MediaThumbnailSizesRo      interface{}  `json:"mediaThumbnailSizesRo,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CreateThumbnails      bool  `json:"createThumbnails,omitempty"`

	KeepAspectRatio      bool  `json:"keepAspectRatio,omitempty"`

	ThumbnailQuality      float64  `json:"thumbnailQuality,omitempty"`

	Private      bool  `json:"private,omitempty"`

}
