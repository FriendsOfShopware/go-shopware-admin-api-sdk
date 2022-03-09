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
	Title string `json:"title,omitempty"`

	DoubleOptInEmailSentDate time.Time `json:"doubleOptInEmailSentDate,omitempty"`

	DefaultPaymentMethod *PaymentMethod `json:"defaultPaymentMethod,omitempty"`

	DefaultShippingAddress *CustomerAddress `json:"defaultShippingAddress,omitempty"`

	Id string `json:"id,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Active bool `json:"active,omitempty"`

	LastLogin time.Time `json:"lastLogin,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Email string `json:"email,omitempty"`

	AffiliateCode string `json:"affiliateCode,omitempty"`

	Newsletter bool `json:"newsletter,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	RequestedGroupId string `json:"requestedGroupId,omitempty"`

	LastOrderDate time.Time `json:"lastOrderDate,omitempty"`

	LastPaymentMethod *PaymentMethod `json:"lastPaymentMethod,omitempty"`

	Password interface{} `json:"password,omitempty"`

	FirstLogin time.Time `json:"firstLogin,omitempty"`

	OrderCount float64 `json:"orderCount,omitempty"`

	DefaultShippingAddressId string `json:"defaultShippingAddressId,omitempty"`

	LastName string `json:"lastName,omitempty"`

	CampaignCode string `json:"campaignCode,omitempty"`

	OrderCustomers []OrderCustomer `json:"orderCustomers,omitempty"`

	RecoveryCustomer *CustomerRecovery `json:"recoveryCustomer,omitempty"`

	Wishlists []CustomerWishlist `json:"wishlists,omitempty"`

	LastPaymentMethodId string `json:"lastPaymentMethodId,omitempty"`

	CustomerNumber string `json:"customerNumber,omitempty"`

	NewsletterSalesChannelIds interface{} `json:"newsletterSalesChannelIds,omitempty"`

	OrderTotalAmount float64 `json:"orderTotalAmount,omitempty"`

	LegacyPassword string `json:"legacyPassword,omitempty"`

	Group *CustomerGroup `json:"group,omitempty"`

	RemoteAddress interface{} `json:"remoteAddress,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	DoubleOptInRegistration bool `json:"doubleOptInRegistration,omitempty"`

	DoubleOptInConfirmDate time.Time `json:"doubleOptInConfirmDate,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	DefaultPaymentMethodId string `json:"defaultPaymentMethodId,omitempty"`

	DefaultBillingAddress *CustomerAddress `json:"defaultBillingAddress,omitempty"`

	Promotions []Promotion `json:"promotions,omitempty"`

	Language *Language `json:"language,omitempty"`

	RequestedGroup *CustomerGroup `json:"requestedGroup,omitempty"`

	Addresses []CustomerAddress `json:"addresses,omitempty"`

	BoundSalesChannelId string `json:"boundSalesChannelId,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	DefaultBillingAddressId string `json:"defaultBillingAddressId,omitempty"`

	Hash string `json:"hash,omitempty"`

	Company string `json:"company,omitempty"`

	VatIds interface{} `json:"vatIds,omitempty"`

	Guest bool `json:"guest,omitempty"`

	Birthday time.Time `json:"birthday,omitempty"`

	LegacyEncoder string `json:"legacyEncoder,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	BoundSalesChannel *SalesChannel `json:"boundSalesChannel,omitempty"`
}

type CustomerCollection struct {
	EntityCollection

	Data []Customer `json:"data"`
}
