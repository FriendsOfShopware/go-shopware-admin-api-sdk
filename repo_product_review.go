package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type ProductReviewRepository ClientService

func (t ProductReviewRepository) Search(ctx ApiContext, criteria Criteria) (*ProductReviewCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-review", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductReviewCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductReviewRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductReviewCollection, *http.Response, error) {
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

func (t ProductReviewRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-review", criteria)

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

func (t ProductReviewRepository) Upsert(ctx ApiContext, entity []ProductReview) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_review": {
		Entity:  "product_review",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductReviewRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_review": {
		Entity:  "product_review",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductReview struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductId string `json:"productId,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	ExternalEmail string `json:"externalEmail,omitempty"`

	Title string `json:"title,omitempty"`

	Points float64 `json:"points,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	ExternalUser string `json:"externalUser,omitempty"`

	Content string `json:"content,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Status bool `json:"status,omitempty"`

	Comment string `json:"comment,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`
}

type ProductReviewCollection struct {
	EntityCollection

	Data []ProductReview `json:"data"`
}
