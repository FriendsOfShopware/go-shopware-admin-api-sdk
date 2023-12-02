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
	ProductSearchKeywords []ProductSearchKeyword `json:"productSearchKeywords,omitempty"`

	AppTranslations []AppTranslation `json:"appTranslations,omitempty"`

	AppCmsBlockTranslations []AppCmsBlockTranslation `json:"appCmsBlockTranslations,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	ActionButtonTranslations []AppActionButtonTranslation `json:"actionButtonTranslations,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	Parent *Language `json:"parent,omitempty"`

	PropertyGroupOptionTranslations []PropertyGroupOptionTranslation `json:"propertyGroupOptionTranslations,omitempty"`

	CmsPageTranslations []CmsPageTranslation `json:"cmsPageTranslations,omitempty"`

	ProductSortingTranslations []ProductSortingTranslation `json:"productSortingTranslations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	TranslationCode *Locale `json:"translationCode,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	CountryTranslations []CountryTranslation `json:"countryTranslations,omitempty"`

	UnitTranslations []UnitTranslation `json:"unitTranslations,omitempty"`

	ProductKeywordDictionaries []ProductKeywordDictionary `json:"productKeywordDictionaries,omitempty"`

	Id string `json:"id,omitempty"`

	NumberRangeTranslations []NumberRangeTranslation `json:"numberRangeTranslations,omitempty"`

	Children []Language `json:"children,omitempty"`

	StateMachineStateTranslations []StateMachineStateTranslation `json:"stateMachineStateTranslations,omitempty"`

	CmsSlotTranslations []CmsSlotTranslation `json:"cmsSlotTranslations,omitempty"`

	LocaleTranslations []LocaleTranslation `json:"localeTranslations,omitempty"`

	ProductSearchConfig *ProductSearchConfig `json:"productSearchConfig,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PromotionTranslations []PromotionTranslation `json:"promotionTranslations,omitempty"`

	ProductManufacturerTranslations []ProductManufacturerTranslation `json:"productManufacturerTranslations,omitempty"`

	StateMachineTranslations []StateMachineTranslation `json:"stateMachineTranslations,omitempty"`

	MailHeaderFooterTranslations []MailHeaderFooterTranslation `json:"mailHeaderFooterTranslations,omitempty"`

	NumberRangeTypeTranslations []NumberRangeTypeTranslation `json:"numberRangeTypeTranslations,omitempty"`

	MailTemplateTypeTranslations []MailTemplateTypeTranslation `json:"mailTemplateTypeTranslations,omitempty"`

	LocaleId string `json:"localeId,omitempty"`

	CategoryTranslations []CategoryTranslation `json:"categoryTranslations,omitempty"`

	CountryStateTranslations []CountryStateTranslation `json:"countryStateTranslations,omitempty"`

	MediaTranslations []MediaTranslation `json:"mediaTranslations,omitempty"`

	PaymentMethodTranslations []PaymentMethodTranslation `json:"paymentMethodTranslations,omitempty"`

	ProductStreamTranslations []ProductStreamTranslation `json:"productStreamTranslations,omitempty"`

	TaxRuleTypeTranslations []TaxRuleTypeTranslation `json:"taxRuleTypeTranslations,omitempty"`

	ImportExportProfileTranslations []ImportExportProfileTranslation `json:"importExportProfileTranslations,omitempty"`

	TranslationCodeId string `json:"translationCodeId,omitempty"`

	ProductFeatureSetTranslations []ProductFeatureSetTranslation `json:"productFeatureSetTranslations,omitempty"`

	MailTemplateTranslations []MailTemplateTranslation `json:"mailTemplateTranslations,omitempty"`

	ProductCrossSellingTranslations []ProductCrossSellingTranslation `json:"productCrossSellingTranslations,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	ShippingMethodTranslations []ShippingMethodTranslation `json:"shippingMethodTranslations,omitempty"`

	SalesChannelTypeTranslations []SalesChannelTypeTranslation `json:"salesChannelTypeTranslations,omitempty"`

	SalutationTranslations []SalutationTranslation `json:"salutationTranslations,omitempty"`

	TaxProviderTranslations []TaxProviderTranslation `json:"taxProviderTranslations,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	ThemeTranslations []ThemeTranslation `json:"themeTranslations,omitempty"`

	DeliveryTimeTranslations []DeliveryTimeTranslation `json:"deliveryTimeTranslations,omitempty"`

	ProductTranslations []ProductTranslation `json:"productTranslations,omitempty"`

	DocumentTypeTranslations []DocumentTypeTranslation `json:"documentTypeTranslations,omitempty"`

	LandingPageTranslations []LandingPageTranslation `json:"landingPageTranslations,omitempty"`

	AppFlowActionTranslations []AppFlowActionTranslation `json:"appFlowActionTranslations,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	PropertyGroupTranslations []PropertyGroupTranslation `json:"propertyGroupTranslations,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	SalesChannelTranslations []SalesChannelTranslation `json:"salesChannelTranslations,omitempty"`

	SeoUrlTranslations []SeoUrl `json:"seoUrlTranslations,omitempty"`

	AppScriptConditionTranslations []AppScriptConditionTranslation `json:"appScriptConditionTranslations,omitempty"`

	Name string `json:"name,omitempty"`

	CurrencyTranslations []CurrencyTranslation `json:"currencyTranslations,omitempty"`

	CustomerGroupTranslations []CustomerGroupTranslation `json:"customerGroupTranslations,omitempty"`

	PluginTranslations []PluginTranslation `json:"pluginTranslations,omitempty"`

	Locale *Locale `json:"locale,omitempty"`
}

type LanguageCollection struct {
	EntityCollection

	Data []Language `json:"data"`
}
