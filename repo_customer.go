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
	Language *Language `json:"language,omitempty"`

	OrderCustomers []OrderCustomer `json:"orderCustomers,omitempty"`

	RequestedGroup *CustomerGroup `json:"requestedGroup,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	VatIds interface{} `json:"vatIds,omitempty"`

	LastOrderDate time.Time `json:"lastOrderDate,omitempty"`

	Addresses []CustomerAddress `json:"addresses,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	LastPaymentMethod *PaymentMethod `json:"lastPaymentMethod,omitempty"`

	Promotions []Promotion `json:"promotions,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	DefaultBillingAddressId string `json:"defaultBillingAddressId,omitempty"`

	Password interface{} `json:"password,omitempty"`

	OrderTotalAmount float64 `json:"orderTotalAmount,omitempty"`

	LegacyEncoder string `json:"legacyEncoder,omitempty"`

	BoundSalesChannelId string `json:"boundSalesChannelId,omitempty"`

	Id string `json:"id,omitempty"`

	FirstLogin time.Time `json:"firstLogin,omitempty"`

	LegacyPassword string `json:"legacyPassword,omitempty"`

	AffiliateCode string `json:"affiliateCode,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Company string `json:"company,omitempty"`

	Email string `json:"email,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	Active bool `json:"active,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	DoubleOptInConfirmDate time.Time `json:"doubleOptInConfirmDate,omitempty"`

	RecoveryCustomer *CustomerRecovery `json:"recoveryCustomer,omitempty"`

	RemoteAddress interface{} `json:"remoteAddress,omitempty"`

	DefaultPaymentMethod *PaymentMethod `json:"defaultPaymentMethod,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Title string `json:"title,omitempty"`

	CampaignCode string `json:"campaignCode,omitempty"`

	LastLogin time.Time `json:"lastLogin,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Guest bool `json:"guest,omitempty"`

	NewsletterSalesChannelIds interface{} `json:"newsletterSalesChannelIds,omitempty"`

	DefaultBillingAddress *CustomerAddress `json:"defaultBillingAddress,omitempty"`

	BoundSalesChannel *SalesChannel `json:"boundSalesChannel,omitempty"`

	DoubleOptInEmailSentDate time.Time `json:"doubleOptInEmailSentDate,omitempty"`

	Newsletter bool `json:"newsletter,omitempty"`

	RequestedGroupId string `json:"requestedGroupId,omitempty"`

	DefaultPaymentMethodId string `json:"defaultPaymentMethodId,omitempty"`

	CustomerNumber string `json:"customerNumber,omitempty"`

	Group *CustomerGroup `json:"group,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LastPaymentMethodId string `json:"lastPaymentMethodId,omitempty"`

	Hash string `json:"hash,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	DoubleOptInRegistration bool `json:"doubleOptInRegistration,omitempty"`

	DefaultShippingAddress *CustomerAddress `json:"defaultShippingAddress,omitempty"`

	Birthday time.Time `json:"birthday,omitempty"`

	OrderCount float64 `json:"orderCount,omitempty"`

	DefaultShippingAddressId string `json:"defaultShippingAddressId,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	LastName string `json:"lastName,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`

	Wishlists []CustomerWishlist `json:"wishlists,omitempty"`
}

type CustomerCollection struct {
	EntityCollection

	Data []Customer `json:"data"`
}
