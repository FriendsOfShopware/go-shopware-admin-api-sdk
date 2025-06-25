package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PromotionSetgroupRepository struct {
	*GenericRepository[PromotionSetgroup]
}

func NewPromotionSetgroupRepository(client *Client) *PromotionSetgroupRepository {
	return &PromotionSetgroupRepository{
		GenericRepository: NewGenericRepository[PromotionSetgroup](client),
	}
}

func (t *PromotionSetgroupRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionSetgroup], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-setgroup")
}

func (t *PromotionSetgroupRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionSetgroup], *http.Response, error) {
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

func (t *PromotionSetgroupRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-setgroup")
}

func (t *PromotionSetgroupRepository) Upsert(ctx ApiContext, entity []PromotionSetgroup) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_setgroup")
}

func (t *PromotionSetgroupRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_setgroup")
}

type PromotionSetgroup struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	PackagerKey      string  `json:"packagerKey,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	SetGroupRules      []Rule  `json:"setGroupRules,omitempty"`

	SorterKey      string  `json:"sorterKey,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Value      float64  `json:"value,omitempty"`

}
