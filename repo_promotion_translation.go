package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PromotionTranslationRepository struct {
	*GenericRepository[PromotionTranslation]
}

func NewPromotionTranslationRepository(client *Client) *PromotionTranslationRepository {
	return &PromotionTranslationRepository{
		GenericRepository: NewGenericRepository[PromotionTranslation](client),
	}
}

func (t *PromotionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-translation")
}

func (t *PromotionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionTranslation], *http.Response, error) {
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

func (t *PromotionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-translation")
}

func (t *PromotionTranslationRepository) Upsert(ctx ApiContext, entity []PromotionTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_translation")
}

func (t *PromotionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_translation")
}

type PromotionTranslation struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}
