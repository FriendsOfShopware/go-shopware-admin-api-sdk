package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MediaThumbnailRepository ClientService

func (t MediaThumbnailRepository) Search(ctx ApiContext, criteria Criteria) (*MediaThumbnailCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-thumbnail", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaThumbnailCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaThumbnailRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaThumbnailCollection, *http.Response, error) {
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

func (t MediaThumbnailRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-thumbnail", criteria)

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

func (t MediaThumbnailRepository) Upsert(ctx ApiContext, entity []MediaThumbnail) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_thumbnail": {
		Entity:  "media_thumbnail",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaThumbnailRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_thumbnail": {
		Entity:  "media_thumbnail",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaThumbnail struct {
	Url string `json:"url,omitempty"`

	Media *Media `json:"media,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Width float64 `json:"width,omitempty"`

	Height float64 `json:"height,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type MediaThumbnailCollection struct {
	EntityCollection

	Data []MediaThumbnail `json:"data"`
}
