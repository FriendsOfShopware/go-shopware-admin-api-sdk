package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SalesChannelRepository ClientService

func (t SalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelCollection, *http.Response, error) {
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

func (t SalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel", criteria)

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

func (t SalesChannelRepository) Upsert(ctx ApiContext, entity []SalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel": {
		Entity:  "sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel": {
		Entity:  "sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannel struct {

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	LandingPages      []LandingPage  `json:"landingPages,omitempty"`

	Wishlists      []CustomerWishlist  `json:"wishlists,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	ShortName      string  `json:"shortName,omitempty"`

	Languages      []Language  `json:"languages,omitempty"`

	PaymentMethodIds      interface{}  `json:"paymentMethodIds,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	SystemConfigs      []SystemConfig  `json:"systemConfigs,omitempty"`

	MailHeaderFooter      *MailHeaderFooter  `json:"mailHeaderFooter,omitempty"`

	SeoUrlTemplates      []SeoUrlTemplate  `json:"seoUrlTemplates,omitempty"`

	FooterCategoryId      string  `json:"footerCategoryId,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	NavigationCategoryVersionId      string  `json:"navigationCategoryVersionId,omitempty"`

	ServiceCategoryId      string  `json:"serviceCategoryId,omitempty"`

	ServiceCategoryVersionId      string  `json:"serviceCategoryVersionId,omitempty"`

	Countries      []Country  `json:"countries,omitempty"`

	NewsletterRecipients      []NewsletterRecipient  `json:"newsletterRecipients,omitempty"`

	BoundCustomers      []Customer  `json:"boundCustomers,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	Themes      []Theme  `json:"themes,omitempty"`

	NavigationCategoryDepth      float64  `json:"navigationCategoryDepth,omitempty"`

	Translations      []SalesChannelTranslation  `json:"translations,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	AnalyticsId      string  `json:"analyticsId,omitempty"`

	HreflangDefaultDomainId      string  `json:"hreflangDefaultDomainId,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	HomeCmsPageId      string  `json:"homeCmsPageId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	FooterCategoryVersionId      string  `json:"footerCategoryVersionId,omitempty"`

	Domains      []SalesChannelDomain  `json:"domains,omitempty"`

	DocumentBaseConfigSalesChannels      []DocumentBaseConfigSalesChannel  `json:"documentBaseConfigSalesChannels,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	HomeMetaTitle      string  `json:"homeMetaTitle,omitempty"`

	HomeKeywords      string  `json:"homeKeywords,omitempty"`

	HreflangDefaultDomain      *SalesChannelDomain  `json:"hreflangDefaultDomain,omitempty"`

	CustomerGroupId      string  `json:"customerGroupId,omitempty"`

	Active      bool  `json:"active,omitempty"`

	TaxCalculationType      string  `json:"taxCalculationType,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	CustomerGroupsRegistrations      []CustomerGroup  `json:"customerGroupsRegistrations,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	HreflangActive      bool  `json:"hreflangActive,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	Analytics      *SalesChannelAnalytics  `json:"analytics,omitempty"`

	MailHeaderFooterId      string  `json:"mailHeaderFooterId,omitempty"`

	Configuration      interface{}  `json:"configuration,omitempty"`

	MaintenanceIpWhitelist      interface{}  `json:"maintenanceIpWhitelist,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	HomeSlotConfig      interface{}  `json:"homeSlotConfig,omitempty"`

	HomeName      string  `json:"homeName,omitempty"`

	NumberRangeSalesChannels      []NumberRangeSalesChannel  `json:"numberRangeSalesChannels,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	NavigationCategoryId      string  `json:"navigationCategoryId,omitempty"`

	Type      *SalesChannelType  `json:"type,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	HomeMetaDescription      string  `json:"homeMetaDescription,omitempty"`

	Id      string  `json:"id,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Currencies      []Currency  `json:"currencies,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	CustomerGroup      *CustomerGroup  `json:"customerGroup,omitempty"`

	HomeCmsPageVersionId      string  `json:"homeCmsPageVersionId,omitempty"`

	HomeCmsPage      *CmsPage  `json:"homeCmsPage,omitempty"`

	HomeEnabled      bool  `json:"homeEnabled,omitempty"`

	NavigationCategory      *Category  `json:"navigationCategory,omitempty"`

	Name      string  `json:"name,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	ProductVisibilities      []ProductVisibility  `json:"productVisibilities,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	PromotionSalesChannels      []PromotionSalesChannel  `json:"promotionSalesChannels,omitempty"`

	AccessKey      string  `json:"accessKey,omitempty"`

	Maintenance      bool  `json:"maintenance,omitempty"`

	FooterCategory      *Category  `json:"footerCategory,omitempty"`

	ServiceCategory      *Category  `json:"serviceCategory,omitempty"`

	TypeId      string  `json:"typeId,omitempty"`

}

type SalesChannelCollection struct {
	EntityCollection

	Data []SalesChannel `json:"data"`
}
