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
	Translations []SalesChannelTranslation `json:"translations,omitempty"`

	Wishlists []CustomerWishlist `json:"wishlists,omitempty"`

	ServiceCategoryVersionId string `json:"serviceCategoryVersionId,omitempty"`

	HreflangDefaultDomain *SalesChannelDomain `json:"hreflangDefaultDomain,omitempty"`

	Themes []Theme `json:"themes,omitempty"`

	TypeId string `json:"typeId,omitempty"`

	NavigationCategoryId string `json:"navigationCategoryId,omitempty"`

	Languages []Language `json:"languages,omitempty"`

	EventActions []EventAction `json:"eventActions,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	ServiceCategoryId string `json:"serviceCategoryId,omitempty"`

	HomeKeywords string `json:"homeKeywords,omitempty"`

	PromotionSalesChannels []PromotionSalesChannel `json:"promotionSalesChannels,omitempty"`

	SeoUrlTemplates []SeoUrlTemplate `json:"seoUrlTemplates,omitempty"`

	HomeMetaTitle string `json:"homeMetaTitle,omitempty"`

	LandingPages []LandingPage `json:"landingPages,omitempty"`

	NavigationCategoryVersionId string `json:"navigationCategoryVersionId,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	HomeCmsPageVersionId string `json:"homeCmsPageVersionId,omitempty"`

	FooterCategory *Category `json:"footerCategory,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	FooterCategoryVersionId string `json:"footerCategoryVersionId,omitempty"`

	Type *SalesChannelType `json:"type,omitempty"`

	HomeSlotConfig interface{} `json:"homeSlotConfig,omitempty"`

	Language *Language `json:"language,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	HomeEnabled bool `json:"homeEnabled,omitempty"`

	NavigationCategory *Category `json:"navigationCategory,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	TaxCalculationType string `json:"taxCalculationType,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CustomerGroupsRegistrations []CustomerGroup `json:"customerGroupsRegistrations,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	Id string `json:"id,omitempty"`

	NavigationCategoryDepth float64 `json:"navigationCategoryDepth,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	HomeMetaDescription string `json:"homeMetaDescription,omitempty"`

	Countries []Country `json:"countries,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	HomeName string `json:"homeName,omitempty"`

	ServiceCategory *Category `json:"serviceCategory,omitempty"`

	SystemConfigs []SystemConfig `json:"systemConfigs,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	BoundCustomers []Customer `json:"boundCustomers,omitempty"`

	PaymentMethodIds interface{} `json:"paymentMethodIds,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	MailHeaderFooter *MailHeaderFooter `json:"mailHeaderFooter,omitempty"`

	ProductExports []ProductExport `json:"productExports,omitempty"`

	MaintenanceIpWhitelist interface{} `json:"maintenanceIpWhitelist,omitempty"`

	Currencies []Currency `json:"currencies,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	CustomerGroup *CustomerGroup `json:"customerGroup,omitempty"`

	HreflangActive bool `json:"hreflangActive,omitempty"`

	HomeCmsPage *CmsPage `json:"homeCmsPage,omitempty"`

	ProductVisibilities []ProductVisibility `json:"productVisibilities,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	Maintenance bool `json:"maintenance,omitempty"`

	Country *Country `json:"country,omitempty"`

	HomeCmsPageId string `json:"homeCmsPageId,omitempty"`

	Domains []SalesChannelDomain `json:"domains,omitempty"`

	AnalyticsId string `json:"analyticsId,omitempty"`

	MailHeaderFooterId string `json:"mailHeaderFooterId,omitempty"`

	Name string `json:"name,omitempty"`

	Configuration interface{} `json:"configuration,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	DocumentBaseConfigSalesChannels []DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`

	Analytics *SalesChannelAnalytics `json:"analytics,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	FooterCategoryId string `json:"footerCategoryId,omitempty"`

	ShortName string `json:"shortName,omitempty"`

	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomerGroupId string `json:"customerGroupId,omitempty"`

	HreflangDefaultDomainId string `json:"hreflangDefaultDomainId,omitempty"`

	Active bool `json:"active,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`
}

type SalesChannelCollection struct {
	EntityCollection

	Data []SalesChannel `json:"data"`
}
