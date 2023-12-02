package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CategoryTagRepository ClientService

func (t CategoryTagRepository) Search(ctx ApiContext, criteria Criteria) (*CategoryTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/category-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CategoryTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CategoryTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CategoryTagCollection, *http.Response, error) {
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

func (t CategoryTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/category-tag", criteria)

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

func (t CategoryTagRepository) Upsert(ctx ApiContext, entity []CategoryTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"category_tag": {
		Entity:  "category_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CategoryTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"category_tag": {
		Entity:  "category_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type CategoryTag struct {

	CategoryId      string  `json:"categoryId,omitempty"`

	CategoryVersionId      string  `json:"categoryVersionId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	Category      *Category  `json:"category,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

}

type CategoryTagCollection struct {
	EntityCollection

	Data []CategoryTag `json:"data"`
}
