package go_shopware_admin_sdk

import (
	"net/http"
)

type MediaTagRepository ClientService

func (t MediaTagRepository) Search(ctx ApiContext, criteria Criteria) (*MediaTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-tag", criteria)

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

func (t MediaTagRepository) Upsert(ctx ApiContext, entity []MediaTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_tag": {
		Entity:  "media_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_tag": {
		Entity:  "media_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaTag struct {
	Media *Media `json:"media,omitempty"`

	Tag *Tag `json:"tag,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	TagId string `json:"tagId,omitempty"`
}

type MediaTagCollection struct {
	EntityCollection

	Data []MediaTag `json:"data"`
}
