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
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Height float64 `json:"height,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Width float64 `json:"width,omitempty"`

	Url string `json:"url,omitempty"`

	Media *Media `json:"media,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`
}

type MediaThumbnailCollection struct {
	EntityCollection

	Data []MediaThumbnail `json:"data"`
}
