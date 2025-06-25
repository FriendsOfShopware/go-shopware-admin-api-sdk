package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomerGroupRegistrationSalesChannelsRepository struct {
	*GenericRepository[CustomerGroupRegistrationSalesChannels]
}

func NewCustomerGroupRegistrationSalesChannelsRepository(client *Client) *CustomerGroupRegistrationSalesChannelsRepository {
	return &CustomerGroupRegistrationSalesChannelsRepository{
		GenericRepository: NewGenericRepository[CustomerGroupRegistrationSalesChannels](client),
	}
}

func (t *CustomerGroupRegistrationSalesChannelsRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerGroupRegistrationSalesChannels], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-group-registration-sales-channels")
}

func (t *CustomerGroupRegistrationSalesChannelsRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerGroupRegistrationSalesChannels], *http.Response, error) {
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

func (t *CustomerGroupRegistrationSalesChannelsRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-group-registration-sales-channels")
}

func (t *CustomerGroupRegistrationSalesChannelsRepository) Upsert(ctx ApiContext, entity []CustomerGroupRegistrationSalesChannels) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_group_registration_sales_channels")
}

func (t *CustomerGroupRegistrationSalesChannelsRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_group_registration_sales_channels")
}

type CustomerGroupRegistrationSalesChannels struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomerGroup      *CustomerGroup  `json:"customerGroup,omitempty"`

	CustomerGroupId      string  `json:"customerGroupId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

}
