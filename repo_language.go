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

	TranslationCodeId      string  `json:"translationCodeId,omitempty"`

	TranslationCode      *Locale  `json:"translationCode,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ProductKeywordDictionaries      []ProductKeywordDictionary  `json:"productKeywordDictionaries,omitempty"`

	LandingPageTranslations      []LandingPageTranslation  `json:"landingPageTranslations,omitempty"`

	ProductSearchConfig      *ProductSearchConfig  `json:"productSearchConfig,omitempty"`

	LocaleId      string  `json:"localeId,omitempty"`

	PaymentMethodTranslations      []PaymentMethodTranslation  `json:"paymentMethodTranslations,omitempty"`

	MailHeaderFooterTranslations      []MailHeaderFooterTranslation  `json:"mailHeaderFooterTranslations,omitempty"`

	SalesChannelTypeTranslations      []SalesChannelTypeTranslation  `json:"salesChannelTypeTranslations,omitempty"`

	StateMachineStateTranslations      []StateMachineStateTranslation  `json:"stateMachineStateTranslations,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	CountryStateTranslations      []CountryStateTranslation  `json:"countryStateTranslations,omitempty"`

	PropertyGroupTranslations      []PropertyGroupTranslation  `json:"propertyGroupTranslations,omitempty"`

	CmsPageTranslations      []CmsPageTranslation  `json:"cmsPageTranslations,omitempty"`

	NumberRangeTypeTranslations      []NumberRangeTypeTranslation  `json:"numberRangeTypeTranslations,omitempty"`

	MailTemplateTypeTranslations      []MailTemplateTypeTranslation  `json:"mailTemplateTypeTranslations,omitempty"`

	PromotionTranslations      []PromotionTranslation  `json:"promotionTranslations,omitempty"`

	ImportExportProfileTranslations      []ImportExportProfileTranslation  `json:"importExportProfileTranslations,omitempty"`

	Children      []Language  `json:"children,omitempty"`

	ShippingMethodTranslations      []ShippingMethodTranslation  `json:"shippingMethodTranslations,omitempty"`

	PluginTranslations      []PluginTranslation  `json:"pluginTranslations,omitempty"`

	ActionButtonTranslations      []AppActionButtonTranslation  `json:"actionButtonTranslations,omitempty"`

	ThemeTranslations      []ThemeTranslation  `json:"themeTranslations,omitempty"`

	CategoryTranslations      []CategoryTranslation  `json:"categoryTranslations,omitempty"`

	MediaTranslations      []MediaTranslation  `json:"mediaTranslations,omitempty"`

	StateMachineTranslations      []StateMachineTranslation  `json:"stateMachineTranslations,omitempty"`

	DocumentTypeTranslations      []DocumentTypeTranslation  `json:"documentTypeTranslations,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	MailTemplateTranslations      []MailTemplateTranslation  `json:"mailTemplateTranslations,omitempty"`

	AppScriptConditionTranslations      []AppScriptConditionTranslation  `json:"appScriptConditionTranslations,omitempty"`

	DeliveryTimeTranslations      []DeliveryTimeTranslation  `json:"deliveryTimeTranslations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	SalesChannelDomains      []SalesChannelDomain  `json:"salesChannelDomains,omitempty"`

	SalutationTranslations      []SalutationTranslation  `json:"salutationTranslations,omitempty"`

	CurrencyTranslations      []CurrencyTranslation  `json:"currencyTranslations,omitempty"`

	UnitTranslations      []UnitTranslation  `json:"unitTranslations,omitempty"`

	Locale      *Locale  `json:"locale,omitempty"`

	ProductSortingTranslations      []ProductSortingTranslation  `json:"productSortingTranslations,omitempty"`

	ProductFeatureSetTranslations      []ProductFeatureSetTranslation  `json:"productFeatureSetTranslations,omitempty"`

	AppCmsBlockTranslations      []AppCmsBlockTranslation  `json:"appCmsBlockTranslations,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	SalesChannelTranslations      []SalesChannelTranslation  `json:"salesChannelTranslations,omitempty"`

	ProductCrossSellingTranslations      []ProductCrossSellingTranslation  `json:"productCrossSellingTranslations,omitempty"`

	NumberRangeTranslations      []NumberRangeTranslation  `json:"numberRangeTranslations,omitempty"`

	TaxRuleTypeTranslations      []TaxRuleTypeTranslation  `json:"taxRuleTypeTranslations,omitempty"`

	Name      string  `json:"name,omitempty"`

	ProductManufacturerTranslations      []ProductManufacturerTranslation  `json:"productManufacturerTranslations,omitempty"`

	ProductSearchKeywords      []ProductSearchKeyword  `json:"productSearchKeywords,omitempty"`

	AppFlowActionTranslations      []AppFlowActionTranslation  `json:"appFlowActionTranslations,omitempty"`

	NewsletterRecipients      []NewsletterRecipient  `json:"newsletterRecipients,omitempty"`

	LocaleTranslations      []LocaleTranslation  `json:"localeTranslations,omitempty"`

	AppTranslations      []AppTranslation  `json:"appTranslations,omitempty"`

	PropertyGroupOptionTranslations      []PropertyGroupOptionTranslation  `json:"propertyGroupOptionTranslations,omitempty"`

	CmsSlotTranslations      []CmsSlotTranslation  `json:"cmsSlotTranslations,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	Parent      *Language  `json:"parent,omitempty"`

	CountryTranslations      []CountryTranslation  `json:"countryTranslations,omitempty"`

	ProductTranslations      []ProductTranslation  `json:"productTranslations,omitempty"`

	TaxProviderTranslations      []TaxProviderTranslation  `json:"taxProviderTranslations,omitempty"`

	CustomerGroupTranslations      []CustomerGroupTranslation  `json:"customerGroupTranslations,omitempty"`

	ProductStreamTranslations      []ProductStreamTranslation  `json:"productStreamTranslations,omitempty"`

	SeoUrlTranslations      []SeoUrl  `json:"seoUrlTranslations,omitempty"`

}

type LanguageCollection struct {
	EntityCollection

	Data []Language `json:"data"`
}
