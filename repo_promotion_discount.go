package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PromotionDiscountRepository struct {
	*GenericRepository[PromotionDiscount]
}

func NewPromotionDiscountRepository(client *Client) *PromotionDiscountRepository {
	return &PromotionDiscountRepository{
		GenericRepository: NewGenericRepository[PromotionDiscount](client),
	}
}

func (t *PromotionDiscountRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionDiscount], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-discount")
}

func (t *PromotionDiscountRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionDiscount], *http.Response, error) {
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

func (t *PromotionDiscountRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-discount")
}

func (t *PromotionDiscountRepository) Upsert(ctx ApiContext, entity []PromotionDiscount) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_discount")
}

func (t *PromotionDiscountRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_discount")
}

type PromotionDiscount struct {

	Type      string  `json:"type,omitempty"`

	Value      float64  `json:"value,omitempty"`

	ConsiderAdvancedRules      bool  `json:"considerAdvancedRules,omitempty"`

	UsageKey      string  `json:"usageKey,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	SorterKey      string  `json:"sorterKey,omitempty"`

	PickerKey      string  `json:"pickerKey,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	MaxValue      float64  `json:"maxValue,omitempty"`

	ApplierKey      string  `json:"applierKey,omitempty"`

	PromotionDiscountPrices      []PromotionDiscountPrices  `json:"promotionDiscountPrices,omitempty"`

	Scope      string  `json:"scope,omitempty"`

	DiscountRules      []Rule  `json:"discountRules,omitempty"`

}
