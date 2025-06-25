package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderRepository struct {
	*GenericRepository[Order]
}

func NewOrderRepository(client *Client) *OrderRepository {
	return &OrderRepository{
		GenericRepository: NewGenericRepository[Order](client),
	}
}

func (t *OrderRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Order], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order")
}

func (t *OrderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Order], *http.Response, error) {
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

func (t *OrderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order")
}

func (t *OrderRepository) Upsert(ctx ApiContext, entity []Order) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order")
}

func (t *OrderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order")
}

type Order struct {

	CustomerComment      string  `json:"customerComment,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	CreatedBy      *User  `json:"createdBy,omitempty"`

	ItemRounding      interface{}  `json:"itemRounding,omitempty"`

	Id      string  `json:"id,omitempty"`

	TaxStatus      string  `json:"taxStatus,omitempty"`

	DeepLinkCode      string  `json:"deepLinkCode,omitempty"`

	Source      string  `json:"source,omitempty"`

	UpdatedById      string  `json:"updatedById,omitempty"`

	LineItems      []OrderLineItem  `json:"lineItems,omitempty"`

	TotalRounding      interface{}  `json:"totalRounding,omitempty"`

	OrderDateTime      time.Time  `json:"orderDateTime,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	Transactions      []OrderTransaction  `json:"transactions,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	BillingAddressVersionId      string  `json:"billingAddressVersionId,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	ShippingTotal      float64  `json:"shippingTotal,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CampaignCode      string  `json:"campaignCode,omitempty"`

	StateId      string  `json:"stateId,omitempty"`

	RuleIds      interface{}  `json:"ruleIds,omitempty"`

	OrderCustomer      *OrderCustomer  `json:"orderCustomer,omitempty"`

	Addresses      []OrderAddress  `json:"addresses,omitempty"`

	BillingAddress      *OrderAddress  `json:"billingAddress,omitempty"`

	Deliveries      []OrderDelivery  `json:"deliveries,omitempty"`

	OrderDate      time.Time  `json:"orderDate,omitempty"`

	AmountTotal      float64  `json:"amountTotal,omitempty"`

	PositionPrice      float64  `json:"positionPrice,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	UpdatedBy      *User  `json:"updatedBy,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	OrderNumber      string  `json:"orderNumber,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	CurrencyFactor      float64  `json:"currencyFactor,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Documents      []Document  `json:"documents,omitempty"`

	AmountNet      float64  `json:"amountNet,omitempty"`

	CreatedById      string  `json:"createdById,omitempty"`

	BillingAddressId      string  `json:"billingAddressId,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	ShippingCosts      interface{}  `json:"shippingCosts,omitempty"`

	AffiliateCode      string  `json:"affiliateCode,omitempty"`

}
