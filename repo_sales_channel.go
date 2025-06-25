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

	CurrencyId      string  `json:"currencyId,omitempty"`

	FooterCategoryVersionId      string  `json:"footerCategoryVersionId,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	NavigationCategory      *Category  `json:"navigationCategory,omitempty"`

	CustomerGroupId      string  `json:"customerGroupId,omitempty"`

	ServiceCategoryVersionId      string  `json:"serviceCategoryVersionId,omitempty"`

	HomeName      string  `json:"homeName,omitempty"`

	HomeMetaDescription      string  `json:"homeMetaDescription,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	AnalyticsId      string  `json:"analyticsId,omitempty"`

	HreflangDefaultDomainId      string  `json:"hreflangDefaultDomainId,omitempty"`

	Name      string  `json:"name,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	NumberRangeSalesChannels      []NumberRangeSalesChannel  `json:"numberRangeSalesChannels,omitempty"`

	PromotionSalesChannels      []PromotionSalesChannel  `json:"promotionSalesChannels,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	HomeCmsPageVersionId      string  `json:"homeCmsPageVersionId,omitempty"`

	Domains      []SalesChannelDomain  `json:"domains,omitempty"`

	DocumentBaseConfigSalesChannels      []DocumentBaseConfigSalesChannel  `json:"documentBaseConfigSalesChannels,omitempty"`

	CustomerGroupsRegistrations      []CustomerGroup  `json:"customerGroupsRegistrations,omitempty"`

	FooterCategoryId      string  `json:"footerCategoryId,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	ServiceCategory      *Category  `json:"serviceCategory,omitempty"`

	Analytics      *SalesChannelAnalytics  `json:"analytics,omitempty"`

	Wishlists      []CustomerWishlist  `json:"wishlists,omitempty"`

	MailHeaderFooterId      string  `json:"mailHeaderFooterId,omitempty"`

	Active      bool  `json:"active,omitempty"`

	HomeEnabled      bool  `json:"homeEnabled,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	NavigationCategoryId      string  `json:"navigationCategoryId,omitempty"`

	HreflangActive      bool  `json:"hreflangActive,omitempty"`

	Type      *SalesChannelType  `json:"type,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	BoundCustomers      []Customer  `json:"boundCustomers,omitempty"`

	Countries      []Country  `json:"countries,omitempty"`

	HomeKeywords      string  `json:"homeKeywords,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	LandingPages      []LandingPage  `json:"landingPages,omitempty"`

	TaxCalculationType      string  `json:"taxCalculationType,omitempty"`

	Configuration      interface{}  `json:"configuration,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	HomeSlotConfig      interface{}  `json:"homeSlotConfig,omitempty"`

	HomeMetaTitle      string  `json:"homeMetaTitle,omitempty"`

	ProductVisibilities      []ProductVisibility  `json:"productVisibilities,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	ServiceCategoryId      string  `json:"serviceCategoryId,omitempty"`

	Maintenance      bool  `json:"maintenance,omitempty"`

	ShortName      string  `json:"shortName,omitempty"`

	HomeCmsPageId      string  `json:"homeCmsPageId,omitempty"`

	SeoUrlTemplates      []SeoUrlTemplate  `json:"seoUrlTemplates,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	NavigationCategoryVersionId      string  `json:"navigationCategoryVersionId,omitempty"`

	Translations      []SalesChannelTranslation  `json:"translations,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	MaintenanceIpWhitelist      interface{}  `json:"maintenanceIpWhitelist,omitempty"`

	Languages      []Language  `json:"languages,omitempty"`

	FooterCategory      *Category  `json:"footerCategory,omitempty"`

	MailHeaderFooter      *MailHeaderFooter  `json:"mailHeaderFooter,omitempty"`

	CustomerGroup      *CustomerGroup  `json:"customerGroup,omitempty"`

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	HomeCmsPage      *CmsPage  `json:"homeCmsPage,omitempty"`

	HreflangDefaultDomain      *SalesChannelDomain  `json:"hreflangDefaultDomain,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	TypeId      string  `json:"typeId,omitempty"`

	Currencies      []Currency  `json:"currencies,omitempty"`

	PaymentMethodIds      interface{}  `json:"paymentMethodIds,omitempty"`

	SystemConfigs      []SystemConfig  `json:"systemConfigs,omitempty"`

	Themes      []Theme  `json:"themes,omitempty"`

	Id      string  `json:"id,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	NavigationCategoryDepth      float64  `json:"navigationCategoryDepth,omitempty"`

	AccessKey      string  `json:"accessKey,omitempty"`

	NewsletterRecipients      []NewsletterRecipient  `json:"newsletterRecipients,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

}
