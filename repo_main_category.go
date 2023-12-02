package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MainCategoryRepository ClientService

func (t MainCategoryRepository) Search(ctx ApiContext, criteria Criteria) (*MainCategoryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/main-category", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MainCategoryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MainCategoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MainCategoryCollection, *http.Response, error) {
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

func (t MainCategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/main-category", criteria)

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

func (t MainCategoryRepository) Upsert(ctx ApiContext, entity []MainCategory) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"main_category": {
		Entity:  "main_category",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MainCategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"main_category": {
		Entity:  "main_category",
		Action:  "delete",
		Payload: payload,
	}})
}

type MainCategory struct {

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	CategoryId      string  `json:"categoryId,omitempty"`

	CategoryVersionId      string  `json:"categoryVersionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	Category      *Category  `json:"category,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

}

type MainCategoryCollection struct {
	EntityCollection

	Data []MainCategory `json:"data"`
}
