package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerRepository ClientService

func (t CustomerRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomerCollection, *http.Response, error) {
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

func (t CustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer", criteria)

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

func (t CustomerRepository) Upsert(ctx ApiContext, entity []Customer) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer": {
		Entity:  "customer",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer": {
		Entity:  "customer",
		Action:  "delete",
		Payload: payload,
	}})
}

type Customer struct {

	DoubleOptInRegistration      bool  `json:"doubleOptInRegistration,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	NewsletterSalesChannelIds      interface{}  `json:"newsletterSalesChannelIds,omitempty"`

	Group      *CustomerGroup  `json:"group,omitempty"`

	DefaultShippingAddressId      string  `json:"defaultShippingAddressId,omitempty"`

	Password      interface{}  `json:"password,omitempty"`

	FirstLogin      time.Time  `json:"firstLogin,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	UpdatedBy      *User  `json:"updatedBy,omitempty"`

	Company      string  `json:"company,omitempty"`

	CampaignCode      string  `json:"campaignCode,omitempty"`

	ReviewCount      float64  `json:"reviewCount,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Addresses      []CustomerAddress  `json:"addresses,omitempty"`

	RemoteAddress      interface{}  `json:"remoteAddress,omitempty"`

	AccountType      string  `json:"accountType,omitempty"`

	GroupId      string  `json:"groupId,omitempty"`

	DefaultBillingAddressId      string  `json:"defaultBillingAddressId,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	LastOrderDate      time.Time  `json:"lastOrderDate,omitempty"`

	OrderTotalAmount      float64  `json:"orderTotalAmount,omitempty"`

	DefaultBillingAddress      *CustomerAddress  `json:"defaultBillingAddress,omitempty"`

	ActiveBillingAddress      *CustomerAddress  `json:"activeBillingAddress,omitempty"`

	Id      string  `json:"id,omitempty"`

	Email      string  `json:"email,omitempty"`

	Title      string  `json:"title,omitempty"`

	Birthday      time.Time  `json:"birthday,omitempty"`

	CreatedBy      *User  `json:"createdBy,omitempty"`

	DefaultPaymentMethodId      string  `json:"defaultPaymentMethodId,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	DefaultShippingAddress      *CustomerAddress  `json:"defaultShippingAddress,omitempty"`

	BoundSalesChannel      *SalesChannel  `json:"boundSalesChannel,omitempty"`

	LastPaymentMethod      *PaymentMethod  `json:"lastPaymentMethod,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	LastPaymentMethodId      string  `json:"lastPaymentMethodId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	LegacyPassword      string  `json:"legacyPassword,omitempty"`

	Guest      bool  `json:"guest,omitempty"`

	OrderCustomers      []OrderCustomer  `json:"orderCustomers,omitempty"`

	RecoveryCustomer      *CustomerRecovery  `json:"recoveryCustomer,omitempty"`

	RequestedGroupId      string  `json:"requestedGroupId,omitempty"`

	DoubleOptInEmailSentDate      time.Time  `json:"doubleOptInEmailSentDate,omitempty"`

	LegacyEncoder      string  `json:"legacyEncoder,omitempty"`

	DefaultPaymentMethod      *PaymentMethod  `json:"defaultPaymentMethod,omitempty"`

	BoundSalesChannelId      string  `json:"boundSalesChannelId,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	RequestedGroup      *CustomerGroup  `json:"requestedGroup,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	DoubleOptInConfirmDate      time.Time  `json:"doubleOptInConfirmDate,omitempty"`

	OrderCount      float64  `json:"orderCount,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	CreatedById      string  `json:"createdById,omitempty"`

	CustomerNumber      string  `json:"customerNumber,omitempty"`

	VatIds      interface{}  `json:"vatIds,omitempty"`

	AffiliateCode      string  `json:"affiliateCode,omitempty"`

	ActiveShippingAddress      *CustomerAddress  `json:"activeShippingAddress,omitempty"`

	UpdatedById      string  `json:"updatedById,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	Promotions      []Promotion  `json:"promotions,omitempty"`

	Wishlists      []CustomerWishlist  `json:"wishlists,omitempty"`

	LastLogin      time.Time  `json:"lastLogin,omitempty"`

	TagIds      interface{}  `json:"tagIds,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type CustomerCollection struct {
	EntityCollection

	Data []Customer `json:"data"`
}
