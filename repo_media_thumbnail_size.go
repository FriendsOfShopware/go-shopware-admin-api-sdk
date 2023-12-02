package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type MediaThumbnailSizeRepository ClientService

func (t MediaThumbnailSizeRepository) Search(ctx ApiContext, criteria Criteria) (*MediaThumbnailSizeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-thumbnail-size", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaThumbnailSizeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaThumbnailSizeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaThumbnailSizeCollection, *http.Response, error) {
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

func (t MediaThumbnailSizeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-thumbnail-size", criteria)

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

func (t MediaThumbnailSizeRepository) Upsert(ctx ApiContext, entity []MediaThumbnailSize) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_thumbnail_size": {
		Entity:  "media_thumbnail_size",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaThumbnailSizeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_thumbnail_size": {
		Entity:  "media_thumbnail_size",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaThumbnailSize struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Width float64 `json:"width,omitempty"`

	Height float64 `json:"height,omitempty"`

	MediaFolderConfigurations []MediaFolderConfiguration `json:"mediaFolderConfigurations,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`
}

type MediaThumbnailSizeCollection struct {
	EntityCollection

	Data []MediaThumbnailSize `json:"data"`
}
