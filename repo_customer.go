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

	AccountType      string  `json:"accountType,omitempty"`

	Active      bool  `json:"active,omitempty"`

	ActiveBillingAddress      *CustomerAddress  `json:"activeBillingAddress,omitempty"`

	ActiveShippingAddress      *CustomerAddress  `json:"activeShippingAddress,omitempty"`

	Addresses      []CustomerAddress  `json:"addresses,omitempty"`

	AffiliateCode      string  `json:"affiliateCode,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	Birthday      time.Time  `json:"birthday,omitempty"`

	BoundSalesChannel      *SalesChannel  `json:"boundSalesChannel,omitempty"`

	BoundSalesChannelId      string  `json:"boundSalesChannelId,omitempty"`

	CampaignCode      string  `json:"campaignCode,omitempty"`

	Company      string  `json:"company,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CreatedBy      *User  `json:"createdBy,omitempty"`

	CreatedById      string  `json:"createdById,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomerNumber      string  `json:"customerNumber,omitempty"`

	DefaultBillingAddress      *CustomerAddress  `json:"defaultBillingAddress,omitempty"`

	DefaultBillingAddressId      string  `json:"defaultBillingAddressId,omitempty"`

	DefaultPaymentMethod      *PaymentMethod  `json:"defaultPaymentMethod,omitempty"`

	DefaultPaymentMethodId      string  `json:"defaultPaymentMethodId,omitempty"`

	DefaultShippingAddress      *CustomerAddress  `json:"defaultShippingAddress,omitempty"`

	DefaultShippingAddressId      string  `json:"defaultShippingAddressId,omitempty"`

	DoubleOptInConfirmDate      time.Time  `json:"doubleOptInConfirmDate,omitempty"`

	DoubleOptInEmailSentDate      time.Time  `json:"doubleOptInEmailSentDate,omitempty"`

	DoubleOptInRegistration      bool  `json:"doubleOptInRegistration,omitempty"`

	Email      string  `json:"email,omitempty"`

	FirstLogin      time.Time  `json:"firstLogin,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Group      *CustomerGroup  `json:"group,omitempty"`

	GroupId      string  `json:"groupId,omitempty"`

	Guest      bool  `json:"guest,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	Id      string  `json:"id,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	LastLogin      time.Time  `json:"lastLogin,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	LastOrderDate      time.Time  `json:"lastOrderDate,omitempty"`

	LastPaymentMethod      *PaymentMethod  `json:"lastPaymentMethod,omitempty"`

	LastPaymentMethodId      string  `json:"lastPaymentMethodId,omitempty"`

	LegacyEncoder      string  `json:"legacyEncoder,omitempty"`

	LegacyPassword      string  `json:"legacyPassword,omitempty"`

	NewsletterSalesChannelIds      interface{}  `json:"newsletterSalesChannelIds,omitempty"`

	OrderCount      float64  `json:"orderCount,omitempty"`

	OrderCustomers      []OrderCustomer  `json:"orderCustomers,omitempty"`

	OrderTotalAmount      float64  `json:"orderTotalAmount,omitempty"`

	Password      interface{}  `json:"password,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	Promotions      []Promotion  `json:"promotions,omitempty"`

	RecoveryCustomer      *CustomerRecovery  `json:"recoveryCustomer,omitempty"`

	RemoteAddress      interface{}  `json:"remoteAddress,omitempty"`

	RequestedGroup      *CustomerGroup  `json:"requestedGroup,omitempty"`

	RequestedGroupId      string  `json:"requestedGroupId,omitempty"`

	ReviewCount      float64  `json:"reviewCount,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	TagIds      interface{}  `json:"tagIds,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Title      string  `json:"title,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	UpdatedBy      *User  `json:"updatedBy,omitempty"`

	UpdatedById      string  `json:"updatedById,omitempty"`

	VatIds      interface{}  `json:"vatIds,omitempty"`

	Wishlists      []CustomerWishlist  `json:"wishlists,omitempty"`

}
