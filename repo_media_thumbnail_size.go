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
	Width float64 `json:"width,omitempty"`

	Height float64 `json:"height,omitempty"`

	MediaFolderConfigurations []MediaFolderConfiguration `json:"mediaFolderConfigurations,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type MediaThumbnailSizeCollection struct {
	EntityCollection

	Data []MediaThumbnailSize `json:"data"`
}
