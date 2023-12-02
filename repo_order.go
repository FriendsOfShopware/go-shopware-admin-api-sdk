package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderRepository ClientService

func (t OrderRepository) Search(ctx ApiContext, criteria Criteria) (*OrderCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderCollection, *http.Response, error) {
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

func (t OrderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order", criteria)

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

func (t OrderRepository) Upsert(ctx ApiContext, entity []Order) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order": {
		Entity:  "order",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order": {
		Entity:  "order",
		Action:  "delete",
		Payload: payload,
	}})
}

type Order struct {

	OrderDateTime      time.Time  `json:"orderDateTime,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Id      string  `json:"id,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	OrderNumber      string  `json:"orderNumber,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	BillingAddressVersionId      string  `json:"billingAddressVersionId,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	AmountNet      float64  `json:"amountNet,omitempty"`

	PositionPrice      float64  `json:"positionPrice,omitempty"`

	ShippingTotal      float64  `json:"shippingTotal,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	Transactions      []OrderTransaction  `json:"transactions,omitempty"`

	ShippingCosts      interface{}  `json:"shippingCosts,omitempty"`

	StateId      string  `json:"stateId,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	OrderDate      time.Time  `json:"orderDate,omitempty"`

	RuleIds      interface{}  `json:"ruleIds,omitempty"`

	CreatedById      string  `json:"createdById,omitempty"`

	LineItems      []OrderLineItem  `json:"lineItems,omitempty"`

	BillingAddress      *OrderAddress  `json:"billingAddress,omitempty"`

	Deliveries      []OrderDelivery  `json:"deliveries,omitempty"`

	Documents      []Document  `json:"documents,omitempty"`

	UpdatedBy      *User  `json:"updatedBy,omitempty"`

	AmountTotal      float64  `json:"amountTotal,omitempty"`

	TaxStatus      string  `json:"taxStatus,omitempty"`

	CampaignCode      string  `json:"campaignCode,omitempty"`

	OrderCustomer      *OrderCustomer  `json:"orderCustomer,omitempty"`

	UpdatedById      string  `json:"updatedById,omitempty"`

	ItemRounding      interface{}  `json:"itemRounding,omitempty"`

	TotalRounding      interface{}  `json:"totalRounding,omitempty"`

	AffiliateCode      string  `json:"affiliateCode,omitempty"`

	CustomerComment      string  `json:"customerComment,omitempty"`

	Source      string  `json:"source,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	BillingAddressId      string  `json:"billingAddressId,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	CurrencyFactor      float64  `json:"currencyFactor,omitempty"`

	DeepLinkCode      string  `json:"deepLinkCode,omitempty"`

	Addresses      []OrderAddress  `json:"addresses,omitempty"`

	CreatedBy      *User  `json:"createdBy,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type OrderCollection struct {
	EntityCollection

	Data []Order `json:"data"`
}
