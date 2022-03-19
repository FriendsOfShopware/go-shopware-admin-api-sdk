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

func (t MediaTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaTagCollection, *http.Response, error) {
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
	MediaId string `json:"mediaId,omitempty"`

	TagId string `json:"tagId,omitempty"`

	Media *Media `json:"media,omitempty"`

	Tag *Tag `json:"tag,omitempty"`
}

type MediaTagCollection struct {
	EntityCollection

	Data []MediaTag `json:"data"`
}
