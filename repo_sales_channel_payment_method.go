package go_shopware_admin_sdk

import (
	"net/http"

)

type SalesChannelPaymentMethodRepository struct {
	*GenericRepository[SalesChannelPaymentMethod]
}

func NewSalesChannelPaymentMethodRepository(client *Client) *SalesChannelPaymentMethodRepository {
	return &SalesChannelPaymentMethodRepository{
		GenericRepository: NewGenericRepository[SalesChannelPaymentMethod](client),
	}
}

func (t *SalesChannelPaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelPaymentMethod], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-payment-method")
}

func (t *SalesChannelPaymentMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelPaymentMethod], *http.Response, error) {
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

func (t *SalesChannelPaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-payment-method")
}

func (t *SalesChannelPaymentMethodRepository) Upsert(ctx ApiContext, entity []SalesChannelPaymentMethod) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_payment_method")
}

func (t *SalesChannelPaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_payment_method")
}

type SalesChannelPaymentMethod struct {

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

}
