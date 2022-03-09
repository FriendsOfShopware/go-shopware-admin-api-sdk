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
	ProductCrossSellingTranslations []ProductCrossSellingTranslation `json:"productCrossSellingTranslations,omitempty"`

	Parent *Language `json:"parent,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	DeliveryTimeTranslations []DeliveryTimeTranslation `json:"deliveryTimeTranslations,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CustomerGroupTranslations []CustomerGroupTranslation `json:"customerGroupTranslations,omitempty"`

	LocaleTranslations []LocaleTranslation `json:"localeTranslations,omitempty"`

	SeoUrlTranslations []SeoUrl `json:"seoUrlTranslations,omitempty"`

	ThemeTranslations []ThemeTranslation `json:"themeTranslations,omitempty"`

	ProductTranslations []ProductTranslation `json:"productTranslations,omitempty"`

	PluginTranslations []PluginTranslation `json:"pluginTranslations,omitempty"`

	MailTemplateTypeTranslations []MailTemplateTypeTranslation `json:"mailTemplateTypeTranslations,omitempty"`

	LandingPageTranslations []LandingPageTranslation `json:"landingPageTranslations,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	CategoryTranslations []CategoryTranslation `json:"categoryTranslations,omitempty"`

	CountryTranslations []CountryTranslation `json:"countryTranslations,omitempty"`

	AppCmsBlockTranslations []AppCmsBlockTranslation `json:"appCmsBlockTranslations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Name string `json:"name,omitempty"`

	CurrencyTranslations []CurrencyTranslation `json:"currencyTranslations,omitempty"`

	ShippingMethodTranslations []ShippingMethodTranslation `json:"shippingMethodTranslations,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	LocaleId string `json:"localeId,omitempty"`

	TranslationCodeId string `json:"translationCodeId,omitempty"`

	ProductStreamTranslations []ProductStreamTranslation `json:"productStreamTranslations,omitempty"`

	AppTranslations []AppTranslation `json:"appTranslations,omitempty"`

	Id string `json:"id,omitempty"`

	Locale *Locale `json:"locale,omitempty"`

	ImportExportProfileTranslations []ImportExportProfileTranslation `json:"importExportProfileTranslations,omitempty"`

	ProductSearchKeywords []ProductSearchKeyword `json:"productSearchKeywords,omitempty"`

	NumberRangeTranslations []NumberRangeTranslation `json:"numberRangeTranslations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	TranslationCode *Locale `json:"translationCode,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	PropertyGroupOptionTranslations []PropertyGroupOptionTranslation `json:"propertyGroupOptionTranslations,omitempty"`

	MailTemplateTranslations []MailTemplateTranslation `json:"mailTemplateTranslations,omitempty"`

	ActionButtonTranslations []AppActionButtonTranslation `json:"actionButtonTranslations,omitempty"`

	MediaTranslations []MediaTranslation `json:"mediaTranslations,omitempty"`

	ProductManufacturerTranslations []ProductManufacturerTranslation `json:"productManufacturerTranslations,omitempty"`

	SalesChannelTypeTranslations []SalesChannelTypeTranslation `json:"salesChannelTypeTranslations,omitempty"`

	CmsSlotTranslations []CmsSlotTranslation `json:"cmsSlotTranslations,omitempty"`

	PromotionTranslations []PromotionTranslation `json:"promotionTranslations,omitempty"`

	PropertyGroupTranslations []PropertyGroupTranslation `json:"propertyGroupTranslations,omitempty"`

	StateMachineStateTranslations []StateMachineStateTranslation `json:"stateMachineStateTranslations,omitempty"`

	CmsPageTranslations []CmsPageTranslation `json:"cmsPageTranslations,omitempty"`

	NumberRangeTypeTranslations []NumberRangeTypeTranslation `json:"numberRangeTypeTranslations,omitempty"`

	ProductSearchConfig *ProductSearchConfig `json:"productSearchConfig,omitempty"`

	Children []Language `json:"children,omitempty"`

	CountryStateTranslations []CountryStateTranslation `json:"countryStateTranslations,omitempty"`

	SalutationTranslations []SalutationTranslation `json:"salutationTranslations,omitempty"`

	TaxRuleTypeTranslations []TaxRuleTypeTranslation `json:"taxRuleTypeTranslations,omitempty"`

	ProductSortingTranslations []ProductSortingTranslation `json:"productSortingTranslations,omitempty"`

	PaymentMethodTranslations []PaymentMethodTranslation `json:"paymentMethodTranslations,omitempty"`

	MailHeaderFooterTranslations []MailHeaderFooterTranslation `json:"mailHeaderFooterTranslations,omitempty"`

	ProductKeywordDictionaries []ProductKeywordDictionary `json:"productKeywordDictionaries,omitempty"`

	ProductFeatureSetTranslations []ProductFeatureSetTranslation `json:"productFeatureSetTranslations,omitempty"`

	SalesChannelTranslations []SalesChannelTranslation `json:"salesChannelTranslations,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	UnitTranslations []UnitTranslation `json:"unitTranslations,omitempty"`

	StateMachineTranslations []StateMachineTranslation `json:"stateMachineTranslations,omitempty"`

	DocumentTypeTranslations []DocumentTypeTranslation `json:"documentTypeTranslations,omitempty"`
}

type LanguageCollection struct {
	EntityCollection

	Data []Language `json:"data"`
}
