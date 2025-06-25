package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomerRepository struct {
	*GenericRepository[Customer]
}

func NewCustomerRepository(client *Client) *CustomerRepository {
	return &CustomerRepository{
		GenericRepository: NewGenericRepository[Customer](client),
	}
}

func (t *CustomerRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Customer], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer")
}

func (t *CustomerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Customer], *http.Response, error) {
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

func (t *CustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer")
}

func (t *CustomerRepository) Upsert(ctx ApiContext, entity []Customer) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer")
}

func (t *CustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer")
}

type Customer struct {

	RequestedGroupId      string  `json:"requestedGroupId,omitempty"`

	Company      string  `json:"company,omitempty"`

	Email      string  `json:"email,omitempty"`

	VatIds      interface{}  `json:"vatIds,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Id      string  `json:"id,omitempty"`

	CampaignCode      string  `json:"campaignCode,omitempty"`

	ActiveBillingAddress      *CustomerAddress  `json:"activeBillingAddress,omitempty"`

	RemoteAddress      interface{}  `json:"remoteAddress,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	OrderTotalAmount      float64  `json:"orderTotalAmount,omitempty"`

	LastPaymentMethod      *PaymentMethod  `json:"lastPaymentMethod,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	UpdatedById      string  `json:"updatedById,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	AffiliateCode      string  `json:"affiliateCode,omitempty"`

	NewsletterSalesChannelIds      interface{}  `json:"newsletterSalesChannelIds,omitempty"`

	RecoveryCustomer      *CustomerRecovery  `json:"recoveryCustomer,omitempty"`

	Wishlists      []CustomerWishlist  `json:"wishlists,omitempty"`

	UpdatedBy      *User  `json:"updatedBy,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	DefaultBillingAddressId      string  `json:"defaultBillingAddressId,omitempty"`

	CreatedById      string  `json:"createdById,omitempty"`

	DefaultShippingAddressId      string  `json:"defaultShippingAddressId,omitempty"`

	DoubleOptInEmailSentDate      time.Time  `json:"doubleOptInEmailSentDate,omitempty"`

	Birthday      time.Time  `json:"birthday,omitempty"`

	OrderCount      float64  `json:"orderCount,omitempty"`

	DefaultPaymentMethodId      string  `json:"defaultPaymentMethodId,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Addresses      []CustomerAddress  `json:"addresses,omitempty"`

	RequestedGroup      *CustomerGroup  `json:"requestedGroup,omitempty"`

	CreatedBy      *User  `json:"createdBy,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	TagIds      interface{}  `json:"tagIds,omitempty"`

	BoundSalesChannelId      string  `json:"boundSalesChannelId,omitempty"`

	LastPaymentMethodId      string  `json:"lastPaymentMethodId,omitempty"`

	CustomerNumber      string  `json:"customerNumber,omitempty"`

	ReviewCount      float64  `json:"reviewCount,omitempty"`

	LegacyPassword      string  `json:"legacyPassword,omitempty"`

	BoundSalesChannel      *SalesChannel  `json:"boundSalesChannel,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	DoubleOptInRegistration      bool  `json:"doubleOptInRegistration,omitempty"`

	LastOrderDate      time.Time  `json:"lastOrderDate,omitempty"`

	DefaultShippingAddress      *CustomerAddress  `json:"defaultShippingAddress,omitempty"`

	OrderCustomers      []OrderCustomer  `json:"orderCustomers,omitempty"`

	GroupId      string  `json:"groupId,omitempty"`

	Guest      bool  `json:"guest,omitempty"`

	FirstLogin      time.Time  `json:"firstLogin,omitempty"`

	LastLogin      time.Time  `json:"lastLogin,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	LegacyEncoder      string  `json:"legacyEncoder,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	ActiveShippingAddress      *CustomerAddress  `json:"activeShippingAddress,omitempty"`

	DefaultBillingAddress      *CustomerAddress  `json:"defaultBillingAddress,omitempty"`

	DefaultPaymentMethod      *PaymentMethod  `json:"defaultPaymentMethod,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Password      interface{}  `json:"password,omitempty"`

	Promotions      []Promotion  `json:"promotions,omitempty"`

	AccountType      string  `json:"accountType,omitempty"`

	Title      string  `json:"title,omitempty"`

	DoubleOptInConfirmDate      time.Time  `json:"doubleOptInConfirmDate,omitempty"`

	Group      *CustomerGroup  `json:"group,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

}
