package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type LanguageRepository ClientService

func (t LanguageRepository) Search(ctx ApiContext, criteria Criteria) (*LanguageCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/language", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LanguageCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LanguageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*LanguageCollection, *http.Response, error) {
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

func (t LanguageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/language", criteria)

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

func (t LanguageRepository) Upsert(ctx ApiContext, entity []Language) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"language": {
		Entity:  "language",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LanguageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"language": {
		Entity:  "language",
		Action:  "delete",
		Payload: payload,
	}})
}

type Language struct {
	ActionButtonTranslations []AppActionButtonTranslation `json:"actionButtonTranslations,omitempty"`

	Id string `json:"id,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	MediaTranslations []MediaTranslation `json:"mediaTranslations,omitempty"`

	MailTemplateTypeTranslations []MailTemplateTypeTranslation `json:"mailTemplateTypeTranslations,omitempty"`

	ProductTranslations []ProductTranslation `json:"productTranslations,omitempty"`

	SalutationTranslations []SalutationTranslation `json:"salutationTranslations,omitempty"`

	ImportExportProfileTranslations []ImportExportProfileTranslation `json:"importExportProfileTranslations,omitempty"`

	ThemeTranslations []ThemeTranslation `json:"themeTranslations,omitempty"`

	PropertyGroupOptionTranslations []PropertyGroupOptionTranslation `json:"propertyGroupOptionTranslations,omitempty"`

	StateMachineStateTranslations []StateMachineStateTranslation `json:"stateMachineStateTranslations,omitempty"`

	ProductSearchKeywords []ProductSearchKeyword `json:"productSearchKeywords,omitempty"`

	LandingPageTranslations []LandingPageTranslation `json:"landingPageTranslations,omitempty"`

	TranslationCode *Locale `json:"translationCode,omitempty"`

	CountryTranslations []CountryTranslation `json:"countryTranslations,omitempty"`

	ShippingMethodTranslations []ShippingMethodTranslation `json:"shippingMethodTranslations,omitempty"`

	DeliveryTimeTranslations []DeliveryTimeTranslation `json:"deliveryTimeTranslations,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	PluginTranslations []PluginTranslation `json:"pluginTranslations,omitempty"`

	ProductStreamTranslations []ProductStreamTranslation `json:"productStreamTranslations,omitempty"`

	CmsPageTranslations []CmsPageTranslation `json:"cmsPageTranslations,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SeoUrlTranslations []SeoUrl `json:"seoUrlTranslations,omitempty"`

	ProductCrossSellingTranslations []ProductCrossSellingTranslation `json:"productCrossSellingTranslations,omitempty"`

	StateMachineTranslations []StateMachineTranslation `json:"stateMachineTranslations,omitempty"`

	PromotionTranslations []PromotionTranslation `json:"promotionTranslations,omitempty"`

	LocaleId string `json:"localeId,omitempty"`

	Locale *Locale `json:"locale,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	SalesChannelTranslations []SalesChannelTranslation `json:"salesChannelTranslations,omitempty"`

	AppCmsBlockTranslations []AppCmsBlockTranslation `json:"appCmsBlockTranslations,omitempty"`

	TranslationCodeId string `json:"translationCodeId,omitempty"`

	CmsSlotTranslations []CmsSlotTranslation `json:"cmsSlotTranslations,omitempty"`

	TaxRuleTypeTranslations []TaxRuleTypeTranslation `json:"taxRuleTypeTranslations,omitempty"`

	ProductSearchConfig *ProductSearchConfig `json:"productSearchConfig,omitempty"`

	ProductSortingTranslations []ProductSortingTranslation `json:"productSortingTranslations,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	CurrencyTranslations []CurrencyTranslation `json:"currencyTranslations,omitempty"`

	UnitTranslations []UnitTranslation `json:"unitTranslations,omitempty"`

	MailTemplateTranslations []MailTemplateTranslation `json:"mailTemplateTranslations,omitempty"`

	Children []Language `json:"children,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	CategoryTranslations []CategoryTranslation `json:"categoryTranslations,omitempty"`

	SalesChannelTypeTranslations []SalesChannelTypeTranslation `json:"salesChannelTypeTranslations,omitempty"`

	CustomerGroupTranslations []CustomerGroupTranslation `json:"customerGroupTranslations,omitempty"`

	CountryStateTranslations []CountryStateTranslation `json:"countryStateTranslations,omitempty"`

	PropertyGroupTranslations []PropertyGroupTranslation `json:"propertyGroupTranslations,omitempty"`

	DocumentTypeTranslations []DocumentTypeTranslation `json:"documentTypeTranslations,omitempty"`

	ProductFeatureSetTranslations []ProductFeatureSetTranslation `json:"productFeatureSetTranslations,omitempty"`

	ProductManufacturerTranslations []ProductManufacturerTranslation `json:"productManufacturerTranslations,omitempty"`

	MailHeaderFooterTranslations []MailHeaderFooterTranslation `json:"mailHeaderFooterTranslations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Name string `json:"name,omitempty"`

	Parent *Language `json:"parent,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	NumberRangeTranslations []NumberRangeTranslation `json:"numberRangeTranslations,omitempty"`

	AppTranslations []AppTranslation `json:"appTranslations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LocaleTranslations []LocaleTranslation `json:"localeTranslations,omitempty"`

	PaymentMethodTranslations []PaymentMethodTranslation `json:"paymentMethodTranslations,omitempty"`

	NumberRangeTypeTranslations []NumberRangeTypeTranslation `json:"numberRangeTypeTranslations,omitempty"`

	ProductKeywordDictionaries []ProductKeywordDictionary `json:"productKeywordDictionaries,omitempty"`
}

type LanguageCollection struct {
	EntityCollection

	Data []Language `json:"data"`
}
