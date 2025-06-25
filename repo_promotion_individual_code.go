package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PromotionIndividualCodeRepository struct {
	*GenericRepository[PromotionIndividualCode]
}

func NewPromotionIndividualCodeRepository(client *Client) *PromotionIndividualCodeRepository {
	return &PromotionIndividualCodeRepository{
		GenericRepository: NewGenericRepository[PromotionIndividualCode](client),
	}
}

func (t *PromotionIndividualCodeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionIndividualCode], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-individual-code")
}

func (t *PromotionIndividualCodeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionIndividualCode], *http.Response, error) {
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

func (t *PromotionIndividualCodeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-individual-code")
}

func (t *PromotionIndividualCodeRepository) Upsert(ctx ApiContext, entity []PromotionIndividualCode) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_individual_code")
}

func (t *PromotionIndividualCodeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_individual_code")
}

type PromotionIndividualCode struct {

	Id      string  `json:"id,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	Code      string  `json:"code,omitempty"`

	Payload      interface{}  `json:"payload,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
