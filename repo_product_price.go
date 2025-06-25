package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductPriceRepository struct {
	*GenericRepository[ProductPrice]
}

func NewProductPriceRepository(client *Client) *ProductPriceRepository {
	return &ProductPriceRepository{
		GenericRepository: NewGenericRepository[ProductPrice](client),
	}
}

func (t *ProductPriceRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductPrice], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-price")
}

func (t *ProductPriceRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductPrice], *http.Response, error) {
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

func (t *ProductPriceRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-price")
}

func (t *ProductPriceRepository) Upsert(ctx ApiContext, entity []ProductPrice) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_price")
}

func (t *ProductPriceRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_price")
}

type ProductPrice struct {

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	QuantityStart      float64  `json:"quantityStart,omitempty"`

	QuantityEnd      float64  `json:"quantityEnd,omitempty"`

	Product      *Product  `json:"product,omitempty"`

}
