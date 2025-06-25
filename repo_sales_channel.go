package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalesChannelRepository struct {
	*GenericRepository[SalesChannel]
}

func NewSalesChannelRepository(client *Client) *SalesChannelRepository {
	return &SalesChannelRepository{
		GenericRepository: NewGenericRepository[SalesChannel](client),
	}
}

func (t *SalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel")
}

func (t *SalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannel], *http.Response, error) {
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

func (t *SalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel")
}

func (t *SalesChannelRepository) Upsert(ctx ApiContext, entity []SalesChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel")
}

func (t *SalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel")
}

type SalesChannel struct {

	AccessKey      string  `json:"accessKey,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Analytics      *SalesChannelAnalytics  `json:"analytics,omitempty"`

	AnalyticsId      string  `json:"analyticsId,omitempty"`

	BoundCustomers      []Customer  `json:"boundCustomers,omitempty"`

	Configuration      interface{}  `json:"configuration,omitempty"`

	Countries      []Country  `json:"countries,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Currencies      []Currency  `json:"currencies,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomerGroup      *CustomerGroup  `json:"customerGroup,omitempty"`

	CustomerGroupId      string  `json:"customerGroupId,omitempty"`

	CustomerGroupsRegistrations      []CustomerGroup  `json:"customerGroupsRegistrations,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	DocumentBaseConfigSalesChannels      []DocumentBaseConfigSalesChannel  `json:"documentBaseConfigSalesChannels,omitempty"`

	Domains      []SalesChannelDomain  `json:"domains,omitempty"`

	FooterCategory      *Category  `json:"footerCategory,omitempty"`

	FooterCategoryId      string  `json:"footerCategoryId,omitempty"`

	FooterCategoryVersionId      string  `json:"footerCategoryVersionId,omitempty"`

	HomeCmsPage      *CmsPage  `json:"homeCmsPage,omitempty"`

	HomeCmsPageId      string  `json:"homeCmsPageId,omitempty"`

	HomeCmsPageVersionId      string  `json:"homeCmsPageVersionId,omitempty"`

	HomeEnabled      bool  `json:"homeEnabled,omitempty"`

	HomeKeywords      string  `json:"homeKeywords,omitempty"`

	HomeMetaDescription      string  `json:"homeMetaDescription,omitempty"`

	HomeMetaTitle      string  `json:"homeMetaTitle,omitempty"`

	HomeName      string  `json:"homeName,omitempty"`

	HomeSlotConfig      interface{}  `json:"homeSlotConfig,omitempty"`

	HreflangActive      bool  `json:"hreflangActive,omitempty"`

	HreflangDefaultDomain      *SalesChannelDomain  `json:"hreflangDefaultDomain,omitempty"`

	HreflangDefaultDomainId      string  `json:"hreflangDefaultDomainId,omitempty"`

	Id      string  `json:"id,omitempty"`

	LandingPages      []LandingPage  `json:"landingPages,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Languages      []Language  `json:"languages,omitempty"`

	MailHeaderFooter      *MailHeaderFooter  `json:"mailHeaderFooter,omitempty"`

	MailHeaderFooterId      string  `json:"mailHeaderFooterId,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	Maintenance      bool  `json:"maintenance,omitempty"`

	MaintenanceIpWhitelist      interface{}  `json:"maintenanceIpWhitelist,omitempty"`

	Name      string  `json:"name,omitempty"`

	NavigationCategory      *Category  `json:"navigationCategory,omitempty"`

	NavigationCategoryDepth      float64  `json:"navigationCategoryDepth,omitempty"`

	NavigationCategoryId      string  `json:"navigationCategoryId,omitempty"`

	NavigationCategoryVersionId      string  `json:"navigationCategoryVersionId,omitempty"`

	NewsletterRecipients      []NewsletterRecipient  `json:"newsletterRecipients,omitempty"`

	NumberRangeSalesChannels      []NumberRangeSalesChannel  `json:"numberRangeSalesChannels,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	PaymentMethodIds      interface{}  `json:"paymentMethodIds,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	ProductVisibilities      []ProductVisibility  `json:"productVisibilities,omitempty"`

	PromotionSalesChannels      []PromotionSalesChannel  `json:"promotionSalesChannels,omitempty"`

	SeoUrlTemplates      []SeoUrlTemplate  `json:"seoUrlTemplates,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	ServiceCategory      *Category  `json:"serviceCategory,omitempty"`

	ServiceCategoryId      string  `json:"serviceCategoryId,omitempty"`

	ServiceCategoryVersionId      string  `json:"serviceCategoryVersionId,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	ShortName      string  `json:"shortName,omitempty"`

	SystemConfigs      []SystemConfig  `json:"systemConfigs,omitempty"`

	TaxCalculationType      string  `json:"taxCalculationType,omitempty"`

	Themes      []Theme  `json:"themes,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []SalesChannelTranslation  `json:"translations,omitempty"`

	Type      *SalesChannelType  `json:"type,omitempty"`

	TypeId      string  `json:"typeId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Wishlists      []CustomerWishlist  `json:"wishlists,omitempty"`

}
