package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PromotionDiscountPricesRepository struct {
	*GenericRepository[PromotionDiscountPrices]
}

func NewPromotionDiscountPricesRepository(client *Client) *PromotionDiscountPricesRepository {
	return &PromotionDiscountPricesRepository{
		GenericRepository: NewGenericRepository[PromotionDiscountPrices](client),
	}
}

func (t *PromotionDiscountPricesRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionDiscountPrices], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-discount-prices")
}

func (t *PromotionDiscountPricesRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionDiscountPrices], *http.Response, error) {
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

func (t *PromotionDiscountPricesRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-discount-prices")
}

func (t *PromotionDiscountPricesRepository) Upsert(ctx ApiContext, entity []PromotionDiscountPrices) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_discount_prices")
}

func (t *PromotionDiscountPricesRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_discount_prices")
}

type PromotionDiscountPrices struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	DiscountId      string  `json:"discountId,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	Price      float64  `json:"price,omitempty"`

	PromotionDiscount      *PromotionDiscount  `json:"promotionDiscount,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}
