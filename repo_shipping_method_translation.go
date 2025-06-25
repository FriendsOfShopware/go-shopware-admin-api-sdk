package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ShippingMethodTranslationRepository struct {
	*GenericRepository[ShippingMethodTranslation]
}

func NewShippingMethodTranslationRepository(client *Client) *ShippingMethodTranslationRepository {
	return &ShippingMethodTranslationRepository{
		GenericRepository: NewGenericRepository[ShippingMethodTranslation](client),
	}
}

func (t *ShippingMethodTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethodTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "shipping-method-translation")
}

func (t *ShippingMethodTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethodTranslation], *http.Response, error) {
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

func (t *ShippingMethodTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "shipping-method-translation")
}

func (t *ShippingMethodTranslationRepository) Upsert(ctx ApiContext, entity []ShippingMethodTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "shipping_method_translation")
}

func (t *ShippingMethodTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "shipping_method_translation")
}

type ShippingMethodTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	TrackingUrl      string  `json:"trackingUrl,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
