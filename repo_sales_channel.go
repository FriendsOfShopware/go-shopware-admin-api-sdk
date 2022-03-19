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
	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	MailHeaderFooterId string `json:"mailHeaderFooterId,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	HomeCmsPageId string `json:"homeCmsPageId,omitempty"`

	MailHeaderFooter *MailHeaderFooter `json:"mailHeaderFooter,omitempty"`

	ProductExports []ProductExport `json:"productExports,omitempty"`

	ProductVisibilities []ProductVisibility `json:"productVisibilities,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	Id string `json:"id,omitempty"`

	NavigationCategoryId string `json:"navigationCategoryId,omitempty"`

	NavigationCategoryVersionId string `json:"navigationCategoryVersionId,omitempty"`

	FooterCategoryVersionId string `json:"footerCategoryVersionId,omitempty"`

	Currencies []Currency `json:"currencies,omitempty"`

	Country *Country `json:"country,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	BoundCustomers []Customer `json:"boundCustomers,omitempty"`

	Languages []Language `json:"languages,omitempty"`

	CustomerGroupsRegistrations []CustomerGroup `json:"customerGroupsRegistrations,omitempty"`

	CustomerGroupId string `json:"customerGroupId,omitempty"`

	Countries []Country `json:"countries,omitempty"`

	Themes []Theme `json:"themes,omitempty"`

	TypeId string `json:"typeId,omitempty"`

	Type *SalesChannelType `json:"type,omitempty"`

	HomeMetaTitle string `json:"homeMetaTitle,omitempty"`

	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`

	Analytics *SalesChannelAnalytics `json:"analytics,omitempty"`

	EventActions []EventAction `json:"eventActions,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	FooterCategoryId string `json:"footerCategoryId,omitempty"`

	Language *Language `json:"language,omitempty"`

	ServiceCategory *Category `json:"serviceCategory,omitempty"`

	HreflangDefaultDomain *SalesChannelDomain `json:"hreflangDefaultDomain,omitempty"`

	Translations []SalesChannelTranslation `json:"translations,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	CustomerGroup *CustomerGroup `json:"customerGroup,omitempty"`

	HomeCmsPage *CmsPage `json:"homeCmsPage,omitempty"`

	HomeName string `json:"homeName,omitempty"`

	Domains []SalesChannelDomain `json:"domains,omitempty"`

	AnalyticsId string `json:"analyticsId,omitempty"`

	PaymentMethodIds interface{} `json:"paymentMethodIds,omitempty"`

	NavigationCategory *Category `json:"navigationCategory,omitempty"`

	FooterCategory *Category `json:"footerCategory,omitempty"`

	TaxCalculationType string `json:"taxCalculationType,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	DocumentBaseConfigSalesChannels []DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`

	Wishlists []CustomerWishlist `json:"wishlists,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	ShortName string `json:"shortName,omitempty"`

	HomeMetaDescription string `json:"homeMetaDescription,omitempty"`

	HomeKeywords string `json:"homeKeywords,omitempty"`

	NavigationCategoryDepth float64 `json:"navigationCategoryDepth,omitempty"`

	HomeCmsPageVersionId string `json:"homeCmsPageVersionId,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	PromotionSalesChannels []PromotionSalesChannel `json:"promotionSalesChannels,omitempty"`

	SeoUrlTemplates []SeoUrlTemplate `json:"seoUrlTemplates,omitempty"`

	ServiceCategoryVersionId string `json:"serviceCategoryVersionId,omitempty"`

	Name string `json:"name,omitempty"`

	HreflangActive bool `json:"hreflangActive,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	LandingPages []LandingPage `json:"landingPages,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	HreflangDefaultDomainId string `json:"hreflangDefaultDomainId,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ServiceCategoryId string `json:"serviceCategoryId,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	Active bool `json:"active,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	HomeSlotConfig interface{} `json:"homeSlotConfig,omitempty"`

	Configuration interface{} `json:"configuration,omitempty"`

	MaintenanceIpWhitelist interface{} `json:"maintenanceIpWhitelist,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	SystemConfigs []SystemConfig `json:"systemConfigs,omitempty"`

	Maintenance bool `json:"maintenance,omitempty"`

	HomeEnabled bool `json:"homeEnabled,omitempty"`
}

type SalesChannelCollection struct {
	EntityCollection

	Data []SalesChannel `json:"data"`
}
