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
	Comment string `json:"comment,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	Title string `json:"title,omitempty"`

	Content string `json:"content,omitempty"`

	Status bool `json:"status,omitempty"`

	Points float64 `json:"points,omitempty"`

	Product *Product `json:"product,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	ProductId string `json:"productId,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	ExternalUser string `json:"externalUser,omitempty"`

	ExternalEmail string `json:"externalEmail,omitempty"`

	Language *Language `json:"language,omitempty"`

	Id string `json:"id,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ProductReviewCollection struct {
	EntityCollection

	Data []ProductReview `json:"data"`
}
