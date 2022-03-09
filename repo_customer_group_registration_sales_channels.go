package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerGroupRegistrationSalesChannelsRepository ClientService

func (t CustomerGroupRegistrationSalesChannelsRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerGroupRegistrationSalesChannelsCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-group-registration-sales-channels", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerGroupRegistrationSalesChannelsCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerGroupRegistrationSalesChannelsRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-group-registration-sales-channels", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerGroupRegistrationSalesChannelsRepository) Upsert(ctx ApiContext, entity []CustomerGroupRegistrationSalesChannels) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_group_registration_sales_channels": {
		Entity:  "customer_group_registration_sales_channels",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerGroupRegistrationSalesChannelsRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_group_registration_sales_channels": {
		Entity:  "customer_group_registration_sales_channels",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerGroupRegistrationSalesChannels struct {
	CustomerGroupId string `json:"customerGroupId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	CustomerGroup *CustomerGroup `json:"customerGroup,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type CustomerGroupRegistrationSalesChannelsCollection struct {
	EntityCollection

	Data []CustomerGroupRegistrationSalesChannels `json:"data"`
}
