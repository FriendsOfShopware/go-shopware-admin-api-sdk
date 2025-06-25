package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MediaThumbnailRepository struct {
	*GenericRepository[MediaThumbnail]
}

func NewMediaThumbnailRepository(client *Client) *MediaThumbnailRepository {
	return &MediaThumbnailRepository{
		GenericRepository: NewGenericRepository[MediaThumbnail](client),
	}
}

func (t *MediaThumbnailRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaThumbnail], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-thumbnail")
}

func (t *MediaThumbnailRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaThumbnail], *http.Response, error) {
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

func (t *MediaThumbnailRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-thumbnail")
}

func (t *MediaThumbnailRepository) Upsert(ctx ApiContext, entity []MediaThumbnail) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_thumbnail")
}

func (t *MediaThumbnailRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_thumbnail")
}

type MediaThumbnail struct {

	Id      string  `json:"id,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Width      float64  `json:"width,omitempty"`

	Height      float64  `json:"height,omitempty"`

	Path      string  `json:"path,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Url      string  `json:"url,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
