package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type LanguageRepository struct {
	*GenericRepository[Language]
}

func NewLanguageRepository(client *Client) *LanguageRepository {
	return &LanguageRepository{
		GenericRepository: NewGenericRepository[Language](client),
	}
}

func (t *LanguageRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Language], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "language")
}

func (t *LanguageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Language], *http.Response, error) {
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

func (t *LanguageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "language")
}

func (t *LanguageRepository) Upsert(ctx ApiContext, entity []Language) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "language")
}

func (t *LanguageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "language")
}

type Language struct {

	StateMachineStateTranslations      []StateMachineStateTranslation  `json:"stateMachineStateTranslations,omitempty"`

	ThemeTranslations      []ThemeTranslation  `json:"themeTranslations,omitempty"`

	LocaleId      string  `json:"localeId,omitempty"`

	CurrencyTranslations      []CurrencyTranslation  `json:"currencyTranslations,omitempty"`

	MailHeaderFooterTranslations      []MailHeaderFooterTranslation  `json:"mailHeaderFooterTranslations,omitempty"`

	TaxRuleTypeTranslations      []TaxRuleTypeTranslation  `json:"taxRuleTypeTranslations,omitempty"`

	ProductSortingTranslations      []ProductSortingTranslation  `json:"productSortingTranslations,omitempty"`

	AppScriptConditionTranslations      []AppScriptConditionTranslation  `json:"appScriptConditionTranslations,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ProductStreamTranslations      []ProductStreamTranslation  `json:"productStreamTranslations,omitempty"`

	CmsPageTranslations      []CmsPageTranslation  `json:"cmsPageTranslations,omitempty"`

	MailTemplateTypeTranslations      []MailTemplateTypeTranslation  `json:"mailTemplateTypeTranslations,omitempty"`

	ActionButtonTranslations      []AppActionButtonTranslation  `json:"actionButtonTranslations,omitempty"`

	ProductManufacturerTranslations      []ProductManufacturerTranslation  `json:"productManufacturerTranslations,omitempty"`

	SalesChannelTranslations      []SalesChannelTranslation  `json:"salesChannelTranslations,omitempty"`

	PromotionTranslations      []PromotionTranslation  `json:"promotionTranslations,omitempty"`

	SeoUrlTranslations      []SeoUrl  `json:"seoUrlTranslations,omitempty"`

	ProductCrossSellingTranslations      []ProductCrossSellingTranslation  `json:"productCrossSellingTranslations,omitempty"`

	AppCmsBlockTranslations      []AppCmsBlockTranslation  `json:"appCmsBlockTranslations,omitempty"`

	TaxProviderTranslations      []TaxProviderTranslation  `json:"taxProviderTranslations,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	Children      []Language  `json:"children,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	NewsletterRecipients      []NewsletterRecipient  `json:"newsletterRecipients,omitempty"`

	CategoryTranslations      []CategoryTranslation  `json:"categoryTranslations,omitempty"`

	PaymentMethodTranslations      []PaymentMethodTranslation  `json:"paymentMethodTranslations,omitempty"`

	AppFlowActionTranslations      []AppFlowActionTranslation  `json:"appFlowActionTranslations,omitempty"`

	CountryStateTranslations      []CountryStateTranslation  `json:"countryStateTranslations,omitempty"`

	PropertyGroupTranslations      []PropertyGroupTranslation  `json:"propertyGroupTranslations,omitempty"`

	SalutationTranslations      []SalutationTranslation  `json:"salutationTranslations,omitempty"`

	PluginTranslations      []PluginTranslation  `json:"pluginTranslations,omitempty"`

	LandingPageTranslations      []LandingPageTranslation  `json:"landingPageTranslations,omitempty"`

	TranslationCodeId      string  `json:"translationCodeId,omitempty"`

	TranslationCode      *Locale  `json:"translationCode,omitempty"`

	CustomerGroupTranslations      []CustomerGroupTranslation  `json:"customerGroupTranslations,omitempty"`

	DeliveryTimeTranslations      []DeliveryTimeTranslation  `json:"deliveryTimeTranslations,omitempty"`

	ImportExportProfileTranslations      []ImportExportProfileTranslation  `json:"importExportProfileTranslations,omitempty"`

	SalesChannelDomains      []SalesChannelDomain  `json:"salesChannelDomains,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	StateMachineTranslations      []StateMachineTranslation  `json:"stateMachineTranslations,omitempty"`

	MailTemplateTranslations      []MailTemplateTranslation  `json:"mailTemplateTranslations,omitempty"`

	Id      string  `json:"id,omitempty"`

	MediaTranslations      []MediaTranslation  `json:"mediaTranslations,omitempty"`

	PropertyGroupOptionTranslations      []PropertyGroupOptionTranslation  `json:"propertyGroupOptionTranslations,omitempty"`

	CmsSlotTranslations      []CmsSlotTranslation  `json:"cmsSlotTranslations,omitempty"`

	DocumentTypeTranslations      []DocumentTypeTranslation  `json:"documentTypeTranslations,omitempty"`

	AppTranslations      []AppTranslation  `json:"appTranslations,omitempty"`

	LocaleTranslations      []LocaleTranslation  `json:"localeTranslations,omitempty"`

	ProductFeatureSetTranslations      []ProductFeatureSetTranslation  `json:"productFeatureSetTranslations,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Parent      *Language  `json:"parent,omitempty"`

	ProductSearchConfig      *ProductSearchConfig  `json:"productSearchConfig,omitempty"`

	NumberRangeTranslations      []NumberRangeTranslation  `json:"numberRangeTranslations,omitempty"`

	UnitTranslations      []UnitTranslation  `json:"unitTranslations,omitempty"`

	Locale      *Locale  `json:"locale,omitempty"`

	ShippingMethodTranslations      []ShippingMethodTranslation  `json:"shippingMethodTranslations,omitempty"`

	SalesChannelTypeTranslations      []SalesChannelTypeTranslation  `json:"salesChannelTypeTranslations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	NumberRangeTypeTranslations      []NumberRangeTypeTranslation  `json:"numberRangeTypeTranslations,omitempty"`

	ProductSearchKeywords      []ProductSearchKeyword  `json:"productSearchKeywords,omitempty"`

	ProductKeywordDictionaries      []ProductKeywordDictionary  `json:"productKeywordDictionaries,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	Name      string  `json:"name,omitempty"`

	CountryTranslations      []CountryTranslation  `json:"countryTranslations,omitempty"`

	ProductTranslations      []ProductTranslation  `json:"productTranslations,omitempty"`

}
