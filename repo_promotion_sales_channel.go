package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PromotionSalesChannelRepository struct {
	*GenericRepository[PromotionSalesChannel]
}

func NewPromotionSalesChannelRepository(client *Client) *PromotionSalesChannelRepository {
	return &PromotionSalesChannelRepository{
		GenericRepository: NewGenericRepository[PromotionSalesChannel](client),
	}
}

func (t *PromotionSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionSalesChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-sales-channel")
}

func (t *PromotionSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionSalesChannel], *http.Response, error) {
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

func (t *PromotionSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-sales-channel")
}

func (t *PromotionSalesChannelRepository) Upsert(ctx ApiContext, entity []PromotionSalesChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_sales_channel")
}

func (t *PromotionSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_sales_channel")
}

type PromotionSalesChannel struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Priority      float64  `json:"priority,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
