package go_shopware_admin_sdk

import (
	"net/http"

)

type SalesChannelShippingMethodRepository struct {
	*GenericRepository[SalesChannelShippingMethod]
}

func NewSalesChannelShippingMethodRepository(client *Client) *SalesChannelShippingMethodRepository {
	return &SalesChannelShippingMethodRepository{
		GenericRepository: NewGenericRepository[SalesChannelShippingMethod](client),
	}
}

func (t *SalesChannelShippingMethodRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelShippingMethod], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-shipping-method")
}

func (t *SalesChannelShippingMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelShippingMethod], *http.Response, error) {
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

func (t *SalesChannelShippingMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-shipping-method")
}

func (t *SalesChannelShippingMethodRepository) Upsert(ctx ApiContext, entity []SalesChannelShippingMethod) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_shipping_method")
}

func (t *SalesChannelShippingMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_shipping_method")
}

type SalesChannelShippingMethod struct {

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

}
