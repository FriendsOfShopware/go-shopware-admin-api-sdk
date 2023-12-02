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
	Company string `json:"company,omitempty"`

	Title string `json:"title,omitempty"`

	VatIds interface{} `json:"vatIds,omitempty"`

	ActiveShippingAddress *CustomerAddress `json:"activeShippingAddress,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	LastPaymentMethod *PaymentMethod `json:"lastPaymentMethod,omitempty"`

	LastOrderDate time.Time `json:"lastOrderDate,omitempty"`

	DefaultPaymentMethod *PaymentMethod `json:"defaultPaymentMethod,omitempty"`

	Password interface{} `json:"password,omitempty"`

	CampaignCode string `json:"campaignCode,omitempty"`

	Active bool `json:"active,omitempty"`

	Group *CustomerGroup `json:"group,omitempty"`

	BoundSalesChannel *SalesChannel `json:"boundSalesChannel,omitempty"`

	CreatedById string `json:"createdById,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	DefaultShippingAddressId string `json:"defaultShippingAddressId,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	FirstLogin time.Time `json:"firstLogin,omitempty"`

	UpdatedById string `json:"updatedById,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	Guest bool `json:"guest,omitempty"`

	LegacyEncoder string `json:"legacyEncoder,omitempty"`

	Language *Language `json:"language,omitempty"`

	ActiveBillingAddress *CustomerAddress `json:"activeBillingAddress,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	Addresses []CustomerAddress `json:"addresses,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	DefaultPaymentMethodId string `json:"defaultPaymentMethodId,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	Hash string `json:"hash,omitempty"`

	AccountType string `json:"accountType,omitempty"`

	CreatedBy *User `json:"createdBy,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LastPaymentMethodId string `json:"lastPaymentMethodId,omitempty"`

	DefaultBillingAddressId string `json:"defaultBillingAddressId,omitempty"`

	LastName string `json:"lastName,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	BoundSalesChannelId string `json:"boundSalesChannelId,omitempty"`

	NewsletterSalesChannelIds interface{} `json:"newsletterSalesChannelIds,omitempty"`

	Promotions []Promotion `json:"promotions,omitempty"`

	Birthday time.Time `json:"birthday,omitempty"`

	RequestedGroup *CustomerGroup `json:"requestedGroup,omitempty"`

	Wishlists []CustomerWishlist `json:"wishlists,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	UpdatedBy *User `json:"updatedBy,omitempty"`

	DoubleOptInConfirmDate time.Time `json:"doubleOptInConfirmDate,omitempty"`

	OrderCount float64 `json:"orderCount,omitempty"`

	LegacyPassword string `json:"legacyPassword,omitempty"`

	DefaultBillingAddress *CustomerAddress `json:"defaultBillingAddress,omitempty"`

	RecoveryCustomer *CustomerRecovery `json:"recoveryCustomer,omitempty"`

	RemoteAddress interface{} `json:"remoteAddress,omitempty"`

	RequestedGroupId string `json:"requestedGroupId,omitempty"`

	Email string `json:"email,omitempty"`

	LastLogin time.Time `json:"lastLogin,omitempty"`

	CustomerNumber string `json:"customerNumber,omitempty"`

	AffiliateCode string `json:"affiliateCode,omitempty"`

	DoubleOptInRegistration bool `json:"doubleOptInRegistration,omitempty"`

	DoubleOptInEmailSentDate time.Time `json:"doubleOptInEmailSentDate,omitempty"`

	ReviewCount float64 `json:"reviewCount,omitempty"`

	OrderCustomers []OrderCustomer `json:"orderCustomers,omitempty"`

	OrderTotalAmount float64 `json:"orderTotalAmount,omitempty"`

	DefaultShippingAddress *CustomerAddress `json:"defaultShippingAddress,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`
}

type CustomerCollection struct {
	EntityCollection

	Data []Customer `json:"data"`
}
