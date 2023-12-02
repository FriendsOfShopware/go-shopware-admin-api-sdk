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
	AmountTotal float64 `json:"amountTotal,omitempty"`

	AmountNet float64 `json:"amountNet,omitempty"`

	ItemRounding interface{} `json:"itemRounding,omitempty"`

	Price interface{} `json:"price,omitempty"`

	DeepLinkCode string `json:"deepLinkCode,omitempty"`

	CreatedBy *User `json:"createdBy,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	Language *Language `json:"language,omitempty"`

	Addresses []OrderAddress `json:"addresses,omitempty"`

	UpdatedBy *User `json:"updatedBy,omitempty"`

	BillingAddressId string `json:"billingAddressId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	AffiliateCode string `json:"affiliateCode,omitempty"`

	StateId string `json:"stateId,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	OrderNumber string `json:"orderNumber,omitempty"`

	ShippingTotal float64 `json:"shippingTotal,omitempty"`

	RuleIds interface{} `json:"ruleIds,omitempty"`

	BillingAddress *OrderAddress `json:"billingAddress,omitempty"`

	Documents []Document `json:"documents,omitempty"`

	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`

	CreatedById string `json:"createdById,omitempty"`

	Deliveries []OrderDelivery `json:"deliveries,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	OrderDateTime time.Time `json:"orderDateTime,omitempty"`

	PositionPrice float64 `json:"positionPrice,omitempty"`

	Source string `json:"source,omitempty"`

	LineItems []OrderLineItem `json:"lineItems,omitempty"`

	Transactions []OrderTransaction `json:"transactions,omitempty"`

	TotalRounding interface{} `json:"totalRounding,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ShippingCosts interface{} `json:"shippingCosts,omitempty"`

	CampaignCode string `json:"campaignCode,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	OrderCustomer *OrderCustomer `json:"orderCustomer,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CustomerComment string `json:"customerComment,omitempty"`

	UpdatedById string `json:"updatedById,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	BillingAddressVersionId string `json:"billingAddressVersionId,omitempty"`

	OrderDate time.Time `json:"orderDate,omitempty"`

	TaxStatus string `json:"taxStatus,omitempty"`

	CurrencyFactor float64 `json:"currencyFactor,omitempty"`
}

type OrderCollection struct {
	EntityCollection

	Data []Order `json:"data"`
}
