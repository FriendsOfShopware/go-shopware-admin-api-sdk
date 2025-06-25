package go_shopware_admin_sdk

type Repository struct {
	ClientService


	AclRole *AclRoleRepository

	AclUserRole *AclUserRoleRepository

	App *AppRepository

	AppActionButton *AppActionButtonRepository

	AppActionButtonTranslation *AppActionButtonTranslationRepository

	AppAdministrationSnippet *AppAdministrationSnippetRepository

	AppCmsBlock *AppCmsBlockRepository

	AppCmsBlockTranslation *AppCmsBlockTranslationRepository

	AppFlowAction *AppFlowActionRepository

	AppFlowActionTranslation *AppFlowActionTranslationRepository

	AppFlowEvent *AppFlowEventRepository

	AppPaymentMethod *AppPaymentMethodRepository

	AppScriptCondition *AppScriptConditionRepository

	AppScriptConditionTranslation *AppScriptConditionTranslationRepository

	AppShippingMethod *AppShippingMethodRepository

	AppTemplate *AppTemplateRepository

	AppTranslation *AppTranslationRepository

	Category *CategoryRepository

	CategoryTag *CategoryTagRepository

	CategoryTranslation *CategoryTranslationRepository

	CmsBlock *CmsBlockRepository

	CmsPage *CmsPageRepository

	CmsPageTranslation *CmsPageTranslationRepository

	CmsSection *CmsSectionRepository

	CmsSlot *CmsSlotRepository

	CmsSlotTranslation *CmsSlotTranslationRepository

	Country *CountryRepository

	CountryState *CountryStateRepository

	CountryStateTranslation *CountryStateTranslationRepository

	CountryTranslation *CountryTranslationRepository

	Currency *CurrencyRepository

	CurrencyCountryRounding *CurrencyCountryRoundingRepository

	CurrencyTranslation *CurrencyTranslationRepository

	CustomEntity *CustomEntityRepository

	CustomField *CustomFieldRepository

	CustomFieldSet *CustomFieldSetRepository

	CustomFieldSetRelation *CustomFieldSetRelationRepository

	Customer *CustomerRepository

	CustomerAddress *CustomerAddressRepository

	CustomerGroup *CustomerGroupRepository

	CustomerGroupRegistrationSalesChannels *CustomerGroupRegistrationSalesChannelsRepository

	CustomerGroupTranslation *CustomerGroupTranslationRepository

	CustomerRecovery *CustomerRecoveryRepository

	CustomerTag *CustomerTagRepository

	CustomerWishlist *CustomerWishlistRepository

	CustomerWishlistProduct *CustomerWishlistProductRepository

	DeliveryTime *DeliveryTimeRepository

	DeliveryTimeTranslation *DeliveryTimeTranslationRepository

	Document *DocumentRepository

	DocumentBaseConfig *DocumentBaseConfigRepository

	DocumentBaseConfigSalesChannel *DocumentBaseConfigSalesChannelRepository

	DocumentType *DocumentTypeRepository

	DocumentTypeTranslation *DocumentTypeTranslationRepository

	Flow *FlowRepository

	FlowSequence *FlowSequenceRepository

	FlowTemplate *FlowTemplateRepository

	ImportExportFile *ImportExportFileRepository

	ImportExportLog *ImportExportLogRepository

	ImportExportProfile *ImportExportProfileRepository

	ImportExportProfileTranslation *ImportExportProfileTranslationRepository

	Integration *IntegrationRepository

	IntegrationRole *IntegrationRoleRepository

	LandingPage *LandingPageRepository

	LandingPageSalesChannel *LandingPageSalesChannelRepository

	LandingPageTag *LandingPageTagRepository

	LandingPageTranslation *LandingPageTranslationRepository

	Language *LanguageRepository

	Locale *LocaleRepository

	LocaleTranslation *LocaleTranslationRepository

	LogEntry *LogEntryRepository

	MailHeaderFooter *MailHeaderFooterRepository

	MailHeaderFooterTranslation *MailHeaderFooterTranslationRepository

	MailTemplate *MailTemplateRepository

	MailTemplateMedia *MailTemplateMediaRepository

	MailTemplateTranslation *MailTemplateTranslationRepository

	MailTemplateType *MailTemplateTypeRepository

	MailTemplateTypeTranslation *MailTemplateTypeTranslationRepository

	MainCategory *MainCategoryRepository

	Media *MediaRepository

	MediaDefaultFolder *MediaDefaultFolderRepository

	MediaFolder *MediaFolderRepository

	MediaFolderConfiguration *MediaFolderConfigurationRepository

	MediaFolderConfigurationMediaThumbnailSize *MediaFolderConfigurationMediaThumbnailSizeRepository

	MediaTag *MediaTagRepository

	MediaThumbnail *MediaThumbnailRepository

	MediaThumbnailSize *MediaThumbnailSizeRepository

	MediaTranslation *MediaTranslationRepository

	NewsletterRecipient *NewsletterRecipientRepository

	NewsletterRecipientTag *NewsletterRecipientTagRepository

	NumberRange *NumberRangeRepository

	NumberRangeSalesChannel *NumberRangeSalesChannelRepository

	NumberRangeState *NumberRangeStateRepository

	NumberRangeTranslation *NumberRangeTranslationRepository

	NumberRangeType *NumberRangeTypeRepository

	NumberRangeTypeTranslation *NumberRangeTypeTranslationRepository

	Order *OrderRepository

	OrderAddress *OrderAddressRepository

	OrderCustomer *OrderCustomerRepository

	OrderDelivery *OrderDeliveryRepository

	OrderDeliveryPosition *OrderDeliveryPositionRepository

	OrderLineItem *OrderLineItemRepository

	OrderLineItemDownload *OrderLineItemDownloadRepository

	OrderTag *OrderTagRepository

	OrderTransaction *OrderTransactionRepository

	OrderTransactionCapture *OrderTransactionCaptureRepository

	OrderTransactionCaptureRefund *OrderTransactionCaptureRefundRepository

	OrderTransactionCaptureRefundPosition *OrderTransactionCaptureRefundPositionRepository

	PaymentMethod *PaymentMethodRepository

	PaymentMethodTranslation *PaymentMethodTranslationRepository

	Plugin *PluginRepository

	PluginTranslation *PluginTranslationRepository

	Product *ProductRepository

	ProductCategory *ProductCategoryRepository

	ProductCategoryTree *ProductCategoryTreeRepository

	ProductConfiguratorSetting *ProductConfiguratorSettingRepository

	ProductCrossSelling *ProductCrossSellingRepository

	ProductCrossSellingAssignedProducts *ProductCrossSellingAssignedProductsRepository

	ProductCrossSellingTranslation *ProductCrossSellingTranslationRepository

	ProductCustomFieldSet *ProductCustomFieldSetRepository

	ProductDownload *ProductDownloadRepository

	ProductExport *ProductExportRepository

	ProductFeatureSet *ProductFeatureSetRepository

	ProductFeatureSetTranslation *ProductFeatureSetTranslationRepository

	ProductKeywordDictionary *ProductKeywordDictionaryRepository

	ProductManufacturer *ProductManufacturerRepository

	ProductManufacturerTranslation *ProductManufacturerTranslationRepository

	ProductMedia *ProductMediaRepository

	ProductOption *ProductOptionRepository

	ProductPrice *ProductPriceRepository

	ProductProperty *ProductPropertyRepository

	ProductReview *ProductReviewRepository

	ProductSearchConfig *ProductSearchConfigRepository

	ProductSearchConfigField *ProductSearchConfigFieldRepository

	ProductSearchKeyword *ProductSearchKeywordRepository

	ProductSorting *ProductSortingRepository

	ProductSortingTranslation *ProductSortingTranslationRepository

	ProductStream *ProductStreamRepository

	ProductStreamFilter *ProductStreamFilterRepository

	ProductStreamMapping *ProductStreamMappingRepository

	ProductStreamTranslation *ProductStreamTranslationRepository

	ProductTag *ProductTagRepository

	ProductTranslation *ProductTranslationRepository

	ProductVisibility *ProductVisibilityRepository

	Promotion *PromotionRepository

	PromotionCartRule *PromotionCartRuleRepository

	PromotionDiscount *PromotionDiscountRepository

	PromotionDiscountPrices *PromotionDiscountPricesRepository

	PromotionDiscountRule *PromotionDiscountRuleRepository

	PromotionIndividualCode *PromotionIndividualCodeRepository

	PromotionOrderRule *PromotionOrderRuleRepository

	PromotionPersonaCustomer *PromotionPersonaCustomerRepository

	PromotionPersonaRule *PromotionPersonaRuleRepository

	PromotionSalesChannel *PromotionSalesChannelRepository

	PromotionSetgroup *PromotionSetgroupRepository

	PromotionSetgroupRule *PromotionSetgroupRuleRepository

	PromotionTranslation *PromotionTranslationRepository

	PropertyGroup *PropertyGroupRepository

	PropertyGroupOption *PropertyGroupOptionRepository

	PropertyGroupOptionTranslation *PropertyGroupOptionTranslationRepository

	PropertyGroupTranslation *PropertyGroupTranslationRepository

	Rule *RuleRepository

	RuleCondition *RuleConditionRepository

	RuleTag *RuleTagRepository

	SalesChannel *SalesChannelRepository

	SalesChannelAnalytics *SalesChannelAnalyticsRepository

	SalesChannelCountry *SalesChannelCountryRepository

	SalesChannelCurrency *SalesChannelCurrencyRepository

	SalesChannelDomain *SalesChannelDomainRepository

	SalesChannelLanguage *SalesChannelLanguageRepository

	SalesChannelPaymentMethod *SalesChannelPaymentMethodRepository

	SalesChannelShippingMethod *SalesChannelShippingMethodRepository

	SalesChannelTranslation *SalesChannelTranslationRepository

	SalesChannelType *SalesChannelTypeRepository

	SalesChannelTypeTranslation *SalesChannelTypeTranslationRepository

	Salutation *SalutationRepository

	SalutationTranslation *SalutationTranslationRepository

	ScheduledTask *ScheduledTaskRepository

	Script *ScriptRepository

	SeoUrl *SeoUrlRepository

	SeoUrlTemplate *SeoUrlTemplateRepository

	ShippingMethod *ShippingMethodRepository

	ShippingMethodPrice *ShippingMethodPriceRepository

	ShippingMethodTag *ShippingMethodTagRepository

	ShippingMethodTranslation *ShippingMethodTranslationRepository

	Snippet *SnippetRepository

	SnippetSet *SnippetSetRepository

	StateMachine *StateMachineRepository

	StateMachineHistory *StateMachineHistoryRepository

	StateMachineState *StateMachineStateRepository

	StateMachineStateTranslation *StateMachineStateTranslationRepository

	StateMachineTransition *StateMachineTransitionRepository

	StateMachineTranslation *StateMachineTranslationRepository

	SystemConfig *SystemConfigRepository

	Tag *TagRepository

	Tax *TaxRepository

	TaxProvider *TaxProviderRepository

	TaxProviderTranslation *TaxProviderTranslationRepository

	TaxRule *TaxRuleRepository

	TaxRuleType *TaxRuleTypeRepository

	TaxRuleTypeTranslation *TaxRuleTypeTranslationRepository

	Theme *ThemeRepository

	ThemeChild *ThemeChildRepository

	ThemeMedia *ThemeMediaRepository

	ThemeSalesChannel *ThemeSalesChannelRepository

	ThemeTranslation *ThemeTranslationRepository

	Unit *UnitRepository

	UnitTranslation *UnitTranslationRepository

	User *UserRepository

	UserAccessKey *UserAccessKeyRepository

	UserConfig *UserConfigRepository

	UserRecovery *UserRecoveryRepository

	Version *VersionRepository

	VersionCommit *VersionCommitRepository

	VersionCommitData *VersionCommitDataRepository

	Webhook *WebhookRepository

	WebhookEventLog *WebhookEventLogRepository

}

func NewRepository(client ClientService) Repository {
	repo := Repository{
		ClientService: client,
	}

	repo.AclRole = NewAclRoleRepository(client.Client)

	repo.AclUserRole = NewAclUserRoleRepository(client.Client)

	repo.App = NewAppRepository(client.Client)

	repo.AppActionButton = NewAppActionButtonRepository(client.Client)

	repo.AppActionButtonTranslation = NewAppActionButtonTranslationRepository(client.Client)

	repo.AppAdministrationSnippet = NewAppAdministrationSnippetRepository(client.Client)

	repo.AppCmsBlock = NewAppCmsBlockRepository(client.Client)

	repo.AppCmsBlockTranslation = NewAppCmsBlockTranslationRepository(client.Client)

	repo.AppFlowAction = NewAppFlowActionRepository(client.Client)

	repo.AppFlowActionTranslation = NewAppFlowActionTranslationRepository(client.Client)

	repo.AppFlowEvent = NewAppFlowEventRepository(client.Client)

	repo.AppPaymentMethod = NewAppPaymentMethodRepository(client.Client)

	repo.AppScriptCondition = NewAppScriptConditionRepository(client.Client)

	repo.AppScriptConditionTranslation = NewAppScriptConditionTranslationRepository(client.Client)

	repo.AppShippingMethod = NewAppShippingMethodRepository(client.Client)

	repo.AppTemplate = NewAppTemplateRepository(client.Client)

	repo.AppTranslation = NewAppTranslationRepository(client.Client)

	repo.Category = NewCategoryRepository(client.Client)

	repo.CategoryTag = NewCategoryTagRepository(client.Client)

	repo.CategoryTranslation = NewCategoryTranslationRepository(client.Client)

	repo.CmsBlock = NewCmsBlockRepository(client.Client)

	repo.CmsPage = NewCmsPageRepository(client.Client)

	repo.CmsPageTranslation = NewCmsPageTranslationRepository(client.Client)

	repo.CmsSection = NewCmsSectionRepository(client.Client)

	repo.CmsSlot = NewCmsSlotRepository(client.Client)

	repo.CmsSlotTranslation = NewCmsSlotTranslationRepository(client.Client)

	repo.Country = NewCountryRepository(client.Client)

	repo.CountryState = NewCountryStateRepository(client.Client)

	repo.CountryStateTranslation = NewCountryStateTranslationRepository(client.Client)

	repo.CountryTranslation = NewCountryTranslationRepository(client.Client)

	repo.Currency = NewCurrencyRepository(client.Client)

	repo.CurrencyCountryRounding = NewCurrencyCountryRoundingRepository(client.Client)

	repo.CurrencyTranslation = NewCurrencyTranslationRepository(client.Client)

	repo.CustomEntity = NewCustomEntityRepository(client.Client)

	repo.CustomField = NewCustomFieldRepository(client.Client)

	repo.CustomFieldSet = NewCustomFieldSetRepository(client.Client)

	repo.CustomFieldSetRelation = NewCustomFieldSetRelationRepository(client.Client)

	repo.Customer = NewCustomerRepository(client.Client)

	repo.CustomerAddress = NewCustomerAddressRepository(client.Client)

	repo.CustomerGroup = NewCustomerGroupRepository(client.Client)

	repo.CustomerGroupRegistrationSalesChannels = NewCustomerGroupRegistrationSalesChannelsRepository(client.Client)

	repo.CustomerGroupTranslation = NewCustomerGroupTranslationRepository(client.Client)

	repo.CustomerRecovery = NewCustomerRecoveryRepository(client.Client)

	repo.CustomerTag = NewCustomerTagRepository(client.Client)

	repo.CustomerWishlist = NewCustomerWishlistRepository(client.Client)

	repo.CustomerWishlistProduct = NewCustomerWishlistProductRepository(client.Client)

	repo.DeliveryTime = NewDeliveryTimeRepository(client.Client)

	repo.DeliveryTimeTranslation = NewDeliveryTimeTranslationRepository(client.Client)

	repo.Document = NewDocumentRepository(client.Client)

	repo.DocumentBaseConfig = NewDocumentBaseConfigRepository(client.Client)

	repo.DocumentBaseConfigSalesChannel = NewDocumentBaseConfigSalesChannelRepository(client.Client)

	repo.DocumentType = NewDocumentTypeRepository(client.Client)

	repo.DocumentTypeTranslation = NewDocumentTypeTranslationRepository(client.Client)

	repo.Flow = NewFlowRepository(client.Client)

	repo.FlowSequence = NewFlowSequenceRepository(client.Client)

	repo.FlowTemplate = NewFlowTemplateRepository(client.Client)

	repo.ImportExportFile = NewImportExportFileRepository(client.Client)

	repo.ImportExportLog = NewImportExportLogRepository(client.Client)

	repo.ImportExportProfile = NewImportExportProfileRepository(client.Client)

	repo.ImportExportProfileTranslation = NewImportExportProfileTranslationRepository(client.Client)

	repo.Integration = NewIntegrationRepository(client.Client)

	repo.IntegrationRole = NewIntegrationRoleRepository(client.Client)

	repo.LandingPage = NewLandingPageRepository(client.Client)

	repo.LandingPageSalesChannel = NewLandingPageSalesChannelRepository(client.Client)

	repo.LandingPageTag = NewLandingPageTagRepository(client.Client)

	repo.LandingPageTranslation = NewLandingPageTranslationRepository(client.Client)

	repo.Language = NewLanguageRepository(client.Client)

	repo.Locale = NewLocaleRepository(client.Client)

	repo.LocaleTranslation = NewLocaleTranslationRepository(client.Client)

	repo.LogEntry = NewLogEntryRepository(client.Client)

	repo.MailHeaderFooter = NewMailHeaderFooterRepository(client.Client)

	repo.MailHeaderFooterTranslation = NewMailHeaderFooterTranslationRepository(client.Client)

	repo.MailTemplate = NewMailTemplateRepository(client.Client)

	repo.MailTemplateMedia = NewMailTemplateMediaRepository(client.Client)

	repo.MailTemplateTranslation = NewMailTemplateTranslationRepository(client.Client)

	repo.MailTemplateType = NewMailTemplateTypeRepository(client.Client)

	repo.MailTemplateTypeTranslation = NewMailTemplateTypeTranslationRepository(client.Client)

	repo.MainCategory = NewMainCategoryRepository(client.Client)

	repo.Media = NewMediaRepository(client.Client)

	repo.MediaDefaultFolder = NewMediaDefaultFolderRepository(client.Client)

	repo.MediaFolder = NewMediaFolderRepository(client.Client)

	repo.MediaFolderConfiguration = NewMediaFolderConfigurationRepository(client.Client)

	repo.MediaFolderConfigurationMediaThumbnailSize = NewMediaFolderConfigurationMediaThumbnailSizeRepository(client.Client)

	repo.MediaTag = NewMediaTagRepository(client.Client)

	repo.MediaThumbnail = NewMediaThumbnailRepository(client.Client)

	repo.MediaThumbnailSize = NewMediaThumbnailSizeRepository(client.Client)

	repo.MediaTranslation = NewMediaTranslationRepository(client.Client)

	repo.NewsletterRecipient = NewNewsletterRecipientRepository(client.Client)

	repo.NewsletterRecipientTag = NewNewsletterRecipientTagRepository(client.Client)

	repo.NumberRange = NewNumberRangeRepository(client.Client)

	repo.NumberRangeSalesChannel = NewNumberRangeSalesChannelRepository(client.Client)

	repo.NumberRangeState = NewNumberRangeStateRepository(client.Client)

	repo.NumberRangeTranslation = NewNumberRangeTranslationRepository(client.Client)

	repo.NumberRangeType = NewNumberRangeTypeRepository(client.Client)

	repo.NumberRangeTypeTranslation = NewNumberRangeTypeTranslationRepository(client.Client)

	repo.Order = NewOrderRepository(client.Client)

	repo.OrderAddress = NewOrderAddressRepository(client.Client)

	repo.OrderCustomer = NewOrderCustomerRepository(client.Client)

	repo.OrderDelivery = NewOrderDeliveryRepository(client.Client)

	repo.OrderDeliveryPosition = NewOrderDeliveryPositionRepository(client.Client)

	repo.OrderLineItem = NewOrderLineItemRepository(client.Client)

	repo.OrderLineItemDownload = NewOrderLineItemDownloadRepository(client.Client)

	repo.OrderTag = NewOrderTagRepository(client.Client)

	repo.OrderTransaction = NewOrderTransactionRepository(client.Client)

	repo.OrderTransactionCapture = NewOrderTransactionCaptureRepository(client.Client)

	repo.OrderTransactionCaptureRefund = NewOrderTransactionCaptureRefundRepository(client.Client)

	repo.OrderTransactionCaptureRefundPosition = NewOrderTransactionCaptureRefundPositionRepository(client.Client)

	repo.PaymentMethod = NewPaymentMethodRepository(client.Client)

	repo.PaymentMethodTranslation = NewPaymentMethodTranslationRepository(client.Client)

	repo.Plugin = NewPluginRepository(client.Client)

	repo.PluginTranslation = NewPluginTranslationRepository(client.Client)

	repo.Product = NewProductRepository(client.Client)

	repo.ProductCategory = NewProductCategoryRepository(client.Client)

	repo.ProductCategoryTree = NewProductCategoryTreeRepository(client.Client)

	repo.ProductConfiguratorSetting = NewProductConfiguratorSettingRepository(client.Client)

	repo.ProductCrossSelling = NewProductCrossSellingRepository(client.Client)

	repo.ProductCrossSellingAssignedProducts = NewProductCrossSellingAssignedProductsRepository(client.Client)

	repo.ProductCrossSellingTranslation = NewProductCrossSellingTranslationRepository(client.Client)

	repo.ProductCustomFieldSet = NewProductCustomFieldSetRepository(client.Client)

	repo.ProductDownload = NewProductDownloadRepository(client.Client)

	repo.ProductExport = NewProductExportRepository(client.Client)

	repo.ProductFeatureSet = NewProductFeatureSetRepository(client.Client)

	repo.ProductFeatureSetTranslation = NewProductFeatureSetTranslationRepository(client.Client)

	repo.ProductKeywordDictionary = NewProductKeywordDictionaryRepository(client.Client)

	repo.ProductManufacturer = NewProductManufacturerRepository(client.Client)

	repo.ProductManufacturerTranslation = NewProductManufacturerTranslationRepository(client.Client)

	repo.ProductMedia = NewProductMediaRepository(client.Client)

	repo.ProductOption = NewProductOptionRepository(client.Client)

	repo.ProductPrice = NewProductPriceRepository(client.Client)

	repo.ProductProperty = NewProductPropertyRepository(client.Client)

	repo.ProductReview = NewProductReviewRepository(client.Client)

	repo.ProductSearchConfig = NewProductSearchConfigRepository(client.Client)

	repo.ProductSearchConfigField = NewProductSearchConfigFieldRepository(client.Client)

	repo.ProductSearchKeyword = NewProductSearchKeywordRepository(client.Client)

	repo.ProductSorting = NewProductSortingRepository(client.Client)

	repo.ProductSortingTranslation = NewProductSortingTranslationRepository(client.Client)

	repo.ProductStream = NewProductStreamRepository(client.Client)

	repo.ProductStreamFilter = NewProductStreamFilterRepository(client.Client)

	repo.ProductStreamMapping = NewProductStreamMappingRepository(client.Client)

	repo.ProductStreamTranslation = NewProductStreamTranslationRepository(client.Client)

	repo.ProductTag = NewProductTagRepository(client.Client)

	repo.ProductTranslation = NewProductTranslationRepository(client.Client)

	repo.ProductVisibility = NewProductVisibilityRepository(client.Client)

	repo.Promotion = NewPromotionRepository(client.Client)

	repo.PromotionCartRule = NewPromotionCartRuleRepository(client.Client)

	repo.PromotionDiscount = NewPromotionDiscountRepository(client.Client)

	repo.PromotionDiscountPrices = NewPromotionDiscountPricesRepository(client.Client)

	repo.PromotionDiscountRule = NewPromotionDiscountRuleRepository(client.Client)

	repo.PromotionIndividualCode = NewPromotionIndividualCodeRepository(client.Client)

	repo.PromotionOrderRule = NewPromotionOrderRuleRepository(client.Client)

	repo.PromotionPersonaCustomer = NewPromotionPersonaCustomerRepository(client.Client)

	repo.PromotionPersonaRule = NewPromotionPersonaRuleRepository(client.Client)

	repo.PromotionSalesChannel = NewPromotionSalesChannelRepository(client.Client)

	repo.PromotionSetgroup = NewPromotionSetgroupRepository(client.Client)

	repo.PromotionSetgroupRule = NewPromotionSetgroupRuleRepository(client.Client)

	repo.PromotionTranslation = NewPromotionTranslationRepository(client.Client)

	repo.PropertyGroup = NewPropertyGroupRepository(client.Client)

	repo.PropertyGroupOption = NewPropertyGroupOptionRepository(client.Client)

	repo.PropertyGroupOptionTranslation = NewPropertyGroupOptionTranslationRepository(client.Client)

	repo.PropertyGroupTranslation = NewPropertyGroupTranslationRepository(client.Client)

	repo.Rule = NewRuleRepository(client.Client)

	repo.RuleCondition = NewRuleConditionRepository(client.Client)

	repo.RuleTag = NewRuleTagRepository(client.Client)

	repo.SalesChannel = NewSalesChannelRepository(client.Client)

	repo.SalesChannelAnalytics = NewSalesChannelAnalyticsRepository(client.Client)

	repo.SalesChannelCountry = NewSalesChannelCountryRepository(client.Client)

	repo.SalesChannelCurrency = NewSalesChannelCurrencyRepository(client.Client)

	repo.SalesChannelDomain = NewSalesChannelDomainRepository(client.Client)

	repo.SalesChannelLanguage = NewSalesChannelLanguageRepository(client.Client)

	repo.SalesChannelPaymentMethod = NewSalesChannelPaymentMethodRepository(client.Client)

	repo.SalesChannelShippingMethod = NewSalesChannelShippingMethodRepository(client.Client)

	repo.SalesChannelTranslation = NewSalesChannelTranslationRepository(client.Client)

	repo.SalesChannelType = NewSalesChannelTypeRepository(client.Client)

	repo.SalesChannelTypeTranslation = NewSalesChannelTypeTranslationRepository(client.Client)

	repo.Salutation = NewSalutationRepository(client.Client)

	repo.SalutationTranslation = NewSalutationTranslationRepository(client.Client)

	repo.ScheduledTask = NewScheduledTaskRepository(client.Client)

	repo.Script = NewScriptRepository(client.Client)

	repo.SeoUrl = NewSeoUrlRepository(client.Client)

	repo.SeoUrlTemplate = NewSeoUrlTemplateRepository(client.Client)

	repo.ShippingMethod = NewShippingMethodRepository(client.Client)

	repo.ShippingMethodPrice = NewShippingMethodPriceRepository(client.Client)

	repo.ShippingMethodTag = NewShippingMethodTagRepository(client.Client)

	repo.ShippingMethodTranslation = NewShippingMethodTranslationRepository(client.Client)

	repo.Snippet = NewSnippetRepository(client.Client)

	repo.SnippetSet = NewSnippetSetRepository(client.Client)

	repo.StateMachine = NewStateMachineRepository(client.Client)

	repo.StateMachineHistory = NewStateMachineHistoryRepository(client.Client)

	repo.StateMachineState = NewStateMachineStateRepository(client.Client)

	repo.StateMachineStateTranslation = NewStateMachineStateTranslationRepository(client.Client)

	repo.StateMachineTransition = NewStateMachineTransitionRepository(client.Client)

	repo.StateMachineTranslation = NewStateMachineTranslationRepository(client.Client)

	repo.SystemConfig = NewSystemConfigRepository(client.Client)

	repo.Tag = NewTagRepository(client.Client)

	repo.Tax = NewTaxRepository(client.Client)

	repo.TaxProvider = NewTaxProviderRepository(client.Client)

	repo.TaxProviderTranslation = NewTaxProviderTranslationRepository(client.Client)

	repo.TaxRule = NewTaxRuleRepository(client.Client)

	repo.TaxRuleType = NewTaxRuleTypeRepository(client.Client)

	repo.TaxRuleTypeTranslation = NewTaxRuleTypeTranslationRepository(client.Client)

	repo.Theme = NewThemeRepository(client.Client)

	repo.ThemeChild = NewThemeChildRepository(client.Client)

	repo.ThemeMedia = NewThemeMediaRepository(client.Client)

	repo.ThemeSalesChannel = NewThemeSalesChannelRepository(client.Client)

	repo.ThemeTranslation = NewThemeTranslationRepository(client.Client)

	repo.Unit = NewUnitRepository(client.Client)

	repo.UnitTranslation = NewUnitTranslationRepository(client.Client)

	repo.User = NewUserRepository(client.Client)

	repo.UserAccessKey = NewUserAccessKeyRepository(client.Client)

	repo.UserConfig = NewUserConfigRepository(client.Client)

	repo.UserRecovery = NewUserRecoveryRepository(client.Client)

	repo.Version = NewVersionRepository(client.Client)

	repo.VersionCommit = NewVersionCommitRepository(client.Client)

	repo.VersionCommitData = NewVersionCommitDataRepository(client.Client)

	repo.Webhook = NewWebhookRepository(client.Client)

	repo.WebhookEventLog = NewWebhookEventLogRepository(client.Client)

	return repo
}
