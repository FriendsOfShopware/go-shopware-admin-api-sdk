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
	Tags []Tag `json:"tags,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	CustomerComment string `json:"customerComment,omitempty"`

	OrderCustomer *OrderCustomer `json:"orderCustomer,omitempty"`

	BillingAddressVersionId string `json:"billingAddressVersionId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Documents []Document `json:"documents,omitempty"`

	RuleIds interface{} `json:"ruleIds,omitempty"`

	CreatedBy *User `json:"createdBy,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CurrencyFactor float64 `json:"currencyFactor,omitempty"`

	DeepLinkCode string `json:"deepLinkCode,omitempty"`

	CampaignCode string `json:"campaignCode,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	Price interface{} `json:"price,omitempty"`

	TaxStatus string `json:"taxStatus,omitempty"`

	ShippingTotal float64 `json:"shippingTotal,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	Id string `json:"id,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	OrderNumber string `json:"orderNumber,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	BillingAddress *OrderAddress `json:"billingAddress,omitempty"`

	UpdatedBy *User `json:"updatedBy,omitempty"`

	AmountNet float64 `json:"amountNet,omitempty"`

	PositionPrice float64 `json:"positionPrice,omitempty"`

	UpdatedById string `json:"updatedById,omitempty"`

	Language *Language `json:"language,omitempty"`

	Addresses []OrderAddress `json:"addresses,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	OrderDate time.Time `json:"orderDate,omitempty"`

	AmountTotal float64 `json:"amountTotal,omitempty"`

	LineItems []OrderLineItem `json:"lineItems,omitempty"`

	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`

	TotalRounding interface{} `json:"totalRounding,omitempty"`

	BillingAddressId string `json:"billingAddressId,omitempty"`

	OrderDateTime time.Time `json:"orderDateTime,omitempty"`

	StateId string `json:"stateId,omitempty"`

	Transactions []OrderTransaction `json:"transactions,omitempty"`

	ItemRounding interface{} `json:"itemRounding,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ShippingCosts interface{} `json:"shippingCosts,omitempty"`

	AffiliateCode string `json:"affiliateCode,omitempty"`

	CreatedById string `json:"createdById,omitempty"`

	Deliveries []OrderDelivery `json:"deliveries,omitempty"`
}

type OrderCollection struct {
	EntityCollection

	Data []Order `json:"data"`
}
