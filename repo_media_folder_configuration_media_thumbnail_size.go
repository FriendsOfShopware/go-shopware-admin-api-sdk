package go_shopware_admin_sdk

import (
	"net/http"

)

type MediaFolderConfigurationMediaThumbnailSizeRepository struct {
	*GenericRepository[MediaFolderConfigurationMediaThumbnailSize]
}

func NewMediaFolderConfigurationMediaThumbnailSizeRepository(client *Client) *MediaFolderConfigurationMediaThumbnailSizeRepository {
	return &MediaFolderConfigurationMediaThumbnailSizeRepository{
		GenericRepository: NewGenericRepository[MediaFolderConfigurationMediaThumbnailSize](client),
	}
}

func (t *MediaFolderConfigurationMediaThumbnailSizeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaFolderConfigurationMediaThumbnailSize], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-folder-configuration-media-thumbnail-size")
}

func (t *MediaFolderConfigurationMediaThumbnailSizeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaFolderConfigurationMediaThumbnailSize], *http.Response, error) {
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

func (t *MediaFolderConfigurationMediaThumbnailSizeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-folder-configuration-media-thumbnail-size")
}

func (t *MediaFolderConfigurationMediaThumbnailSizeRepository) Upsert(ctx ApiContext, entity []MediaFolderConfigurationMediaThumbnailSize) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_folder_configuration_media_thumbnail_size")
}

func (t *MediaFolderConfigurationMediaThumbnailSizeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_folder_configuration_media_thumbnail_size")
}

type MediaFolderConfigurationMediaThumbnailSize struct {

	MediaFolderConfigurationId      string  `json:"mediaFolderConfigurationId,omitempty"`

	MediaThumbnailSizeId      string  `json:"mediaThumbnailSizeId,omitempty"`

	MediaFolderConfiguration      *MediaFolderConfiguration  `json:"mediaFolderConfiguration,omitempty"`

	MediaThumbnailSize      *MediaThumbnailSize  `json:"mediaThumbnailSize,omitempty"`

}
