package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type PromotionDiscountPricesRepository ClientService

func (t PromotionDiscountPricesRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionDiscountPricesCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-discount-prices", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionDiscountPricesCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionDiscountPricesRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionDiscountPricesCollection, *http.Response, error) {
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

func (t PromotionDiscountPricesRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-discount-prices", criteria)

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

func (t PromotionDiscountPricesRepository) Upsert(ctx ApiContext, entity []PromotionDiscountPrices) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_discount_prices": {
		Entity:  "promotion_discount_prices",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionDiscountPricesRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_discount_prices": {
		Entity:  "promotion_discount_prices",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionDiscountPrices struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	DiscountId string `json:"discountId,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	Price float64 `json:"price,omitempty"`

	PromotionDiscount *PromotionDiscount `json:"promotionDiscount,omitempty"`

	Currency *Currency `json:"currency,omitempty"`
}

type PromotionDiscountPricesCollection struct {
	EntityCollection

	Data []PromotionDiscountPrices `json:"data"`
}
