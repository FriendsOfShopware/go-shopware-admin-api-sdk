package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductReviewRepository struct {
	*GenericRepository[ProductReview]
}

func NewProductReviewRepository(client *Client) *ProductReviewRepository {
	return &ProductReviewRepository{
		GenericRepository: NewGenericRepository[ProductReview](client),
	}
}

func (t *ProductReviewRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductReview], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-review")
}

func (t *ProductReviewRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductReview], *http.Response, error) {
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

func (t *ProductReviewRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-review")
}

func (t *ProductReviewRepository) Upsert(ctx ApiContext, entity []ProductReview) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_review")
}

func (t *ProductReviewRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_review")
}

type ProductReview struct {

	Comment      string  `json:"comment,omitempty"`

	Content      string  `json:"content,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Customer      *Customer  `json:"customer,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	ExternalEmail      string  `json:"externalEmail,omitempty"`

	ExternalUser      string  `json:"externalUser,omitempty"`

	Id      string  `json:"id,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Points      float64  `json:"points,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	Status      bool  `json:"status,omitempty"`

	Title      string  `json:"title,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
