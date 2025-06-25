package go_shopware_admin_sdk

import (
	"net/http"

)

type MediaTagRepository struct {
	*GenericRepository[MediaTag]
}

func NewMediaTagRepository(client *Client) *MediaTagRepository {
	return &MediaTagRepository{
		GenericRepository: NewGenericRepository[MediaTag](client),
	}
}

func (t *MediaTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-tag")
}

func (t *MediaTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaTag], *http.Response, error) {
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

func (t *MediaTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-tag")
}

func (t *MediaTagRepository) Upsert(ctx ApiContext, entity []MediaTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_tag")
}

func (t *MediaTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_tag")
}

type MediaTag struct {

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

}
