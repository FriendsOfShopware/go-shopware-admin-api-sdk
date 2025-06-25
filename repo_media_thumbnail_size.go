package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MediaThumbnailSizeRepository struct {
	*GenericRepository[MediaThumbnailSize]
}

func NewMediaThumbnailSizeRepository(client *Client) *MediaThumbnailSizeRepository {
	return &MediaThumbnailSizeRepository{
		GenericRepository: NewGenericRepository[MediaThumbnailSize](client),
	}
}

func (t *MediaThumbnailSizeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaThumbnailSize], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-thumbnail-size")
}

func (t *MediaThumbnailSizeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaThumbnailSize], *http.Response, error) {
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

func (t *MediaThumbnailSizeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-thumbnail-size")
}

func (t *MediaThumbnailSizeRepository) Upsert(ctx ApiContext, entity []MediaThumbnailSize) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_thumbnail_size")
}

func (t *MediaThumbnailSizeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_thumbnail_size")
}

type MediaThumbnailSize struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Height      float64  `json:"height,omitempty"`

	Id      string  `json:"id,omitempty"`

	MediaFolderConfigurations      []MediaFolderConfiguration  `json:"mediaFolderConfigurations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Width      float64  `json:"width,omitempty"`

}
