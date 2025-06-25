package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type DeliveryTimeTranslationRepository struct {
	*GenericRepository[DeliveryTimeTranslation]
}

func NewDeliveryTimeTranslationRepository(client *Client) *DeliveryTimeTranslationRepository {
	return &DeliveryTimeTranslationRepository{
		GenericRepository: NewGenericRepository[DeliveryTimeTranslation](client),
	}
}

func (t *DeliveryTimeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DeliveryTimeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "delivery-time-translation")
}

func (t *DeliveryTimeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DeliveryTimeTranslation], *http.Response, error) {
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

func (t *DeliveryTimeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "delivery-time-translation")
}

func (t *DeliveryTimeTranslationRepository) Upsert(ctx ApiContext, entity []DeliveryTimeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "delivery_time_translation")
}

func (t *DeliveryTimeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "delivery_time_translation")
}

type DeliveryTimeTranslation struct {

	DeliveryTime      *DeliveryTime  `json:"deliveryTime,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	DeliveryTimeId      string  `json:"deliveryTimeId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

}
