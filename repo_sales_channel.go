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
	ServiceCategory *Category `json:"serviceCategory,omitempty"`

	MailHeaderFooter *MailHeaderFooter `json:"mailHeaderFooter,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	ServiceCategoryId string `json:"serviceCategoryId,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	Maintenance bool `json:"maintenance,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	HomeCmsPageId string `json:"homeCmsPageId,omitempty"`

	Wishlists []CustomerWishlist `json:"wishlists,omitempty"`

	TypeId string `json:"typeId,omitempty"`

	MailHeaderFooterId string `json:"mailHeaderFooterId,omitempty"`

	HreflangDefaultDomainId string `json:"hreflangDefaultDomainId,omitempty"`

	BoundCustomers []Customer `json:"boundCustomers,omitempty"`

	Themes []Theme `json:"themes,omitempty"`

	NavigationCategoryDepth float64 `json:"navigationCategoryDepth,omitempty"`

	ShortName string `json:"shortName,omitempty"`

	MaintenanceIpWhitelist interface{} `json:"maintenanceIpWhitelist,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	FooterCategoryId string `json:"footerCategoryId,omitempty"`

	HreflangDefaultDomain *SalesChannelDomain `json:"hreflangDefaultDomain,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	Configuration interface{} `json:"configuration,omitempty"`

	Translations []SalesChannelTranslation `json:"translations,omitempty"`

	HomeMetaDescription string `json:"homeMetaDescription,omitempty"`

	CustomerGroupsRegistrations []CustomerGroup `json:"customerGroupsRegistrations,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	Countries []Country `json:"countries,omitempty"`

	Country *Country `json:"country,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	PromotionSalesChannels []PromotionSalesChannel `json:"promotionSalesChannels,omitempty"`

	ProductExports []ProductExport `json:"productExports,omitempty"`

	Domains []SalesChannelDomain `json:"domains,omitempty"`

	ProductVisibilities []ProductVisibility `json:"productVisibilities,omitempty"`

	CustomerGroupId string `json:"customerGroupId,omitempty"`

	ServiceCategoryVersionId string `json:"serviceCategoryVersionId,omitempty"`

	Name string `json:"name,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	HomeCmsPageVersionId string `json:"homeCmsPageVersionId,omitempty"`

	HomeCmsPage *CmsPage `json:"homeCmsPage,omitempty"`

	DocumentBaseConfigSalesChannels []DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`

	SeoUrlTemplates []SeoUrlTemplate `json:"seoUrlTemplates,omitempty"`

	FooterCategoryVersionId string `json:"footerCategoryVersionId,omitempty"`

	HreflangActive bool `json:"hreflangActive,omitempty"`

	FooterCategory *Category `json:"footerCategory,omitempty"`

	Id string `json:"id,omitempty"`

	NavigationCategoryId string `json:"navigationCategoryId,omitempty"`

	Currencies []Currency `json:"currencies,omitempty"`

	HomeSlotConfig interface{} `json:"homeSlotConfig,omitempty"`

	HomeEnabled bool `json:"homeEnabled,omitempty"`

	Analytics *SalesChannelAnalytics `json:"analytics,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	PaymentMethodIds interface{} `json:"paymentMethodIds,omitempty"`

	HomeMetaTitle string `json:"homeMetaTitle,omitempty"`

	HomeKeywords string `json:"homeKeywords,omitempty"`

	TaxCalculationType string `json:"taxCalculationType,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	NavigationCategory *Category `json:"navigationCategory,omitempty"`

	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AnalyticsId string `json:"analyticsId,omitempty"`

	Active bool `json:"active,omitempty"`

	Languages []Language `json:"languages,omitempty"`

	Type *SalesChannelType `json:"type,omitempty"`

	Language *Language `json:"language,omitempty"`

	CustomerGroup *CustomerGroup `json:"customerGroup,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	LandingPages []LandingPage `json:"landingPages,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	NavigationCategoryVersionId string `json:"navigationCategoryVersionId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	HomeName string `json:"homeName,omitempty"`

	SystemConfigs []SystemConfig `json:"systemConfigs,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`
}

type SalesChannelCollection struct {
	EntityCollection

	Data []SalesChannel `json:"data"`
}
