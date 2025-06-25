package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ShippingMethodRepository struct {
	*GenericRepository[ShippingMethod]
}

func NewShippingMethodRepository(client *Client) *ShippingMethodRepository {
	return &ShippingMethodRepository{
		GenericRepository: NewGenericRepository[ShippingMethod](client),
	}
}

func (t *ShippingMethodRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethod], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "shipping-method")
}

func (t *ShippingMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethod], *http.Response, error) {
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

func (t *ShippingMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "shipping-method")
}

func (t *ShippingMethodRepository) Upsert(ctx ApiContext, entity []ShippingMethod) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "shipping_method")
}

func (t *ShippingMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "shipping_method")
}

type ShippingMethod struct {

	Active      bool  `json:"active,omitempty"`

	AppShippingMethod      *AppShippingMethod  `json:"appShippingMethod,omitempty"`

	AvailabilityRule      *Rule  `json:"availabilityRule,omitempty"`

	AvailabilityRuleId      string  `json:"availabilityRuleId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DeliveryTime      *DeliveryTime  `json:"deliveryTime,omitempty"`

	DeliveryTimeId      string  `json:"deliveryTimeId,omitempty"`

	Description      string  `json:"description,omitempty"`

	Id      string  `json:"id,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Name      string  `json:"name,omitempty"`

	OrderDeliveries      []OrderDelivery  `json:"orderDeliveries,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Prices      []ShippingMethodPrice  `json:"prices,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Tax      *Tax  `json:"tax,omitempty"`

	TaxId      string  `json:"taxId,omitempty"`

	TaxType      string  `json:"taxType,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	TrackingUrl      string  `json:"trackingUrl,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []ShippingMethodTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
