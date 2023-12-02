package go_shopware_admin_sdk

type Repository struct {
	ClientService

	CustomerWishlist *CustomerWishlistRepository

	SalesChannelAnalytics *SalesChannelAnalyticsRepository

	AppActionButton *AppActionButtonRepository

	AppFlowEvent *AppFlowEventRepository

	CustomerRecovery *CustomerRecoveryRepository

	ProductCategoryTree *ProductCategoryTreeRepository

	ProductStream *ProductStreamRepository

	App *AppRepository

	CategoryTranslation *CategoryTranslationRepository

	LandingPageTranslation *LandingPageTranslationRepository

	OrderTag *OrderTagRepository

	PromotionSetgroup *PromotionSetgroupRepository

	SalesChannelTranslation *SalesChannelTranslationRepository

	ThemeMedia *ThemeMediaRepository

	Theme *ThemeRepository

	ProductSearchConfigField *ProductSearchConfigFieldRepository

	StateMachineHistory *StateMachineHistoryRepository

	UserConfig *UserConfigRepository

	AppScriptConditionTranslation *AppScriptConditionTranslationRepository

	ProductExport *ProductExportRepository

	PromotionDiscountRule *PromotionDiscountRuleRepository

	RuleCondition *RuleConditionRepository

	TaxProviderTranslation *TaxProviderTranslationRepository

	WebhookEventLog *WebhookEventLogRepository

	CountryStateTranslation *CountryStateTranslationRepository

	MailTemplateTranslation *MailTemplateTranslationRepository

	NewsletterRecipientTag *NewsletterRecipientTagRepository

	SnippetSet *SnippetSetRepository

	UserAccessKey *UserAccessKeyRepository

	PropertyGroup *PropertyGroupRepository

	CustomerTag *CustomerTagRepository

	ImportExportProfileTranslation *ImportExportProfileTranslationRepository

	SalesChannelCountry *SalesChannelCountryRepository

	PromotionPersonaCustomer *PromotionPersonaCustomerRepository

	Version *VersionRepository

	Language *LanguageRepository

	PromotionSalesChannel *PromotionSalesChannelRepository

	Salutation *SalutationRepository

	StateMachine *StateMachineRepository

	StateMachineTranslation *StateMachineTranslationRepository

	UnitTranslation *UnitTranslationRepository

	ThemeChild *ThemeChildRepository

	AppScriptCondition *AppScriptConditionRepository

	CategoryTag *CategoryTagRepository

	LogEntry *LogEntryRepository

	MediaThumbnail *MediaThumbnailRepository

	OrderDelivery *OrderDeliveryRepository

	ShippingMethod *ShippingMethodRepository

	CmsPageTranslation *CmsPageTranslationRepository

	CmsSlot *CmsSlotRepository

	OrderTransaction *OrderTransactionRepository

	Plugin *PluginRepository

	OrderLineItemDownload *OrderLineItemDownloadRepository

	ProductOption *ProductOptionRepository

	ProductPrice *ProductPriceRepository

	ProductTranslation *ProductTranslationRepository

	TaxRuleTypeTranslation *TaxRuleTypeTranslationRepository

	ProductStreamTranslation *ProductStreamTranslationRepository

	Rule *RuleRepository

	SalesChannelLanguage *SalesChannelLanguageRepository

	PropertyGroupOption *PropertyGroupOptionRepository

	SalesChannelPaymentMethod *SalesChannelPaymentMethodRepository

	AppAdministrationSnippet *AppAdministrationSnippetRepository

	AppCmsBlockTranslation *AppCmsBlockTranslationRepository

	AppFlowActionTranslation *AppFlowActionTranslationRepository

	NumberRangeState *NumberRangeStateRepository

	PluginTranslation *PluginTranslationRepository

	PromotionOrderRule *PromotionOrderRuleRepository

	MediaThumbnailSize *MediaThumbnailSizeRepository

	ProductDownload *ProductDownloadRepository

	Tax *TaxRepository

	CustomerGroupTranslation *CustomerGroupTranslationRepository

	MailHeaderFooter *MailHeaderFooterRepository

	FlowSequence *FlowSequenceRepository

	Locale *LocaleRepository

	OrderCustomer *OrderCustomerRepository

	OrderDeliveryPosition *OrderDeliveryPositionRepository

	ProductCrossSellingTranslation *ProductCrossSellingTranslationRepository

	Media *MediaRepository

	Order *OrderRepository

	CmsPage *CmsPageRepository

	FlowTemplate *FlowTemplateRepository

	MailHeaderFooterTranslation *MailHeaderFooterTranslationRepository

	MediaFolder *MediaFolderRepository

	MediaTag *MediaTagRepository

	User *UserRepository

	Country *CountryRepository

	NumberRangeType *NumberRangeTypeRepository

	SalesChannel *SalesChannelRepository

	CustomerGroup *CustomerGroupRepository

	ProductCrossSellingAssignedProducts *ProductCrossSellingAssignedProductsRepository

	OrderLineItem *OrderLineItemRepository

	OrderTransactionCaptureRefund *OrderTransactionCaptureRefundRepository

	ProductCrossSelling *ProductCrossSellingRepository

	ProductManufacturer *ProductManufacturerRepository

	TaxProvider *TaxProviderRepository

	Flow *FlowRepository

	ImportExportProfile *ImportExportProfileRepository

	ProductVisibility *ProductVisibilityRepository

	PromotionDiscount *PromotionDiscountRepository

	ScheduledTask *ScheduledTaskRepository

	DocumentBaseConfigSalesChannel *DocumentBaseConfigSalesChannelRepository

	NumberRangeSalesChannel *NumberRangeSalesChannelRepository

	ShippingMethodTranslation *ShippingMethodTranslationRepository

	AppCmsBlock *AppCmsBlockRepository

	AppTemplate *AppTemplateRepository

	CustomFieldSet *CustomFieldSetRepository

	LandingPageSalesChannel *LandingPageSalesChannelRepository

	ProductTag *ProductTagRepository

	Script *ScriptRepository

	Customer *CustomerRepository

	SeoUrl *SeoUrlRepository

	ProductProperty *ProductPropertyRepository

	Snippet *SnippetRepository

	CurrencyCountryRounding *CurrencyCountryRoundingRepository

	MailTemplateTypeTranslation *MailTemplateTypeTranslationRepository

	TaxRule *TaxRuleRepository

	CmsSection *CmsSectionRepository

	MediaDefaultFolder *MediaDefaultFolderRepository

	MediaFolderConfiguration *MediaFolderConfigurationRepository

	MediaTranslation *MediaTranslationRepository

	ProductReview *ProductReviewRepository

	SalesChannelShippingMethod *SalesChannelShippingMethodRepository

	DeliveryTimeTranslation *DeliveryTimeTranslationRepository

	ProductConfiguratorSetting *ProductConfiguratorSettingRepository

	ProductSortingTranslation *ProductSortingTranslationRepository

	RuleTag *RuleTagRepository

	SalesChannelCurrency *SalesChannelCurrencyRepository

	DocumentBaseConfig *DocumentBaseConfigRepository

	PropertyGroupOptionTranslation *PropertyGroupOptionTranslationRepository

	CurrencyTranslation *CurrencyTranslationRepository

	LocaleTranslation *LocaleTranslationRepository

	ProductStreamFilter *ProductStreamFilterRepository

	ProductSearchKeyword *ProductSearchKeywordRepository

	StateMachineStateTranslation *StateMachineStateTranslationRepository

	ThemeTranslation *ThemeTranslationRepository

	NumberRangeTypeTranslation *NumberRangeTypeTranslationRepository

	PaymentMethodTranslation *PaymentMethodTranslationRepository

	VersionCommit *VersionCommitRepository

	Webhook *WebhookRepository

	ProductSearchConfig *ProductSearchConfigRepository

	AppActionButtonTranslation *AppActionButtonTranslationRepository

	CustomerAddress *CustomerAddressRepository

	DocumentTypeTranslation *DocumentTypeTranslationRepository

	Product *ProductRepository

	ProductCustomFieldSet *ProductCustomFieldSetRepository

	PromotionDiscountPrices *PromotionDiscountPricesRepository

	OrderTransactionCapture *OrderTransactionCaptureRepository

	PropertyGroupTranslation *PropertyGroupTranslationRepository

	CustomFieldSetRelation *CustomFieldSetRelationRepository

	CustomerGroupRegistrationSalesChannels *CustomerGroupRegistrationSalesChannelsRepository

	ProductFeatureSetTranslation *ProductFeatureSetTranslationRepository

	ShippingMethodPrice *ShippingMethodPriceRepository

	StateMachineTransition *StateMachineTransitionRepository

	ProductFeatureSet *ProductFeatureSetRepository

	DeliveryTime *DeliveryTimeRepository

	ProductKeywordDictionary *ProductKeywordDictionaryRepository

	ProductManufacturerTranslation *ProductManufacturerTranslationRepository

	PromotionCartRule *PromotionCartRuleRepository

	SalesChannelType *SalesChannelTypeRepository

	Document *DocumentRepository

	UserRecovery *UserRecoveryRepository

	PromotionPersonaRule *PromotionPersonaRuleRepository

	VersionCommitData *VersionCommitDataRepository

	AppTranslation *AppTranslationRepository

	PromotionIndividualCode *PromotionIndividualCodeRepository

	Unit *UnitRepository

	TaxRuleType *TaxRuleTypeRepository

	AclRole *AclRoleRepository

	CustomerWishlistProduct *CustomerWishlistProductRepository

	MailTemplateType *MailTemplateTypeRepository

	NumberRange *NumberRangeRepository

	NumberRangeTranslation *NumberRangeTranslationRepository

	StateMachineState *StateMachineStateRepository

	CustomField *CustomFieldRepository

	LandingPageTag *LandingPageTagRepository

	SalesChannelTypeTranslation *SalesChannelTypeTranslationRepository

	SalutationTranslation *SalutationTranslationRepository

	ShippingMethodTag *ShippingMethodTagRepository

	AppFlowAction *AppFlowActionRepository

	Integration *IntegrationRepository

	LandingPage *LandingPageRepository

	Tag *TagRepository

	AppPaymentMethod *AppPaymentMethodRepository

	CountryState *CountryStateRepository

	NewsletterRecipient *NewsletterRecipientRepository

	ProductMedia *ProductMediaRepository

	ThemeSalesChannel *ThemeSalesChannelRepository

	CountryTranslation *CountryTranslationRepository

	Currency *CurrencyRepository

	DocumentType *DocumentTypeRepository

	MainCategory *MainCategoryRepository

	ProductCategory *ProductCategoryRepository

	SystemConfig *SystemConfigRepository

	CmsBlock *CmsBlockRepository

	AclUserRole *AclUserRoleRepository

	CmsSlotTranslation *CmsSlotTranslationRepository

	SeoUrlTemplate *SeoUrlTemplateRepository

	Category *CategoryRepository

	ImportExportLog *ImportExportLogRepository

	MediaFolderConfigurationMediaThumbnailSize *MediaFolderConfigurationMediaThumbnailSizeRepository

	PaymentMethod *PaymentMethodRepository

	ProductStreamMapping *ProductStreamMappingRepository

	PromotionSetgroupRule *PromotionSetgroupRuleRepository

	AppShippingMethod *AppShippingMethodRepository

	ImportExportFile *ImportExportFileRepository

	OrderTransactionCaptureRefundPosition *OrderTransactionCaptureRefundPositionRepository

	Promotion *PromotionRepository

	SalesChannelDomain *SalesChannelDomainRepository

	IntegrationRole *IntegrationRoleRepository

	MailTemplate *MailTemplateRepository

	OrderAddress *OrderAddressRepository

	ProductSorting *ProductSortingRepository

	CustomEntity *CustomEntityRepository

	MailTemplateMedia *MailTemplateMediaRepository

	PromotionTranslation *PromotionTranslationRepository
}

func NewRepository(client ClientService) Repository {
	repo := Repository{
		ClientService: client,
	}

	repo.CustomerWishlist = (*CustomerWishlistRepository)(&client)

	repo.SalesChannelAnalytics = (*SalesChannelAnalyticsRepository)(&client)

	repo.AppActionButton = (*AppActionButtonRepository)(&client)

	repo.AppFlowEvent = (*AppFlowEventRepository)(&client)

	repo.CustomerRecovery = (*CustomerRecoveryRepository)(&client)

	repo.ProductCategoryTree = (*ProductCategoryTreeRepository)(&client)

	repo.ProductStream = (*ProductStreamRepository)(&client)

	repo.App = (*AppRepository)(&client)

	repo.CategoryTranslation = (*CategoryTranslationRepository)(&client)

	repo.LandingPageTranslation = (*LandingPageTranslationRepository)(&client)

	repo.OrderTag = (*OrderTagRepository)(&client)

	repo.PromotionSetgroup = (*PromotionSetgroupRepository)(&client)

	repo.SalesChannelTranslation = (*SalesChannelTranslationRepository)(&client)

	repo.ThemeMedia = (*ThemeMediaRepository)(&client)

	repo.Theme = (*ThemeRepository)(&client)

	repo.ProductSearchConfigField = (*ProductSearchConfigFieldRepository)(&client)

	repo.StateMachineHistory = (*StateMachineHistoryRepository)(&client)

	repo.UserConfig = (*UserConfigRepository)(&client)

	repo.AppScriptConditionTranslation = (*AppScriptConditionTranslationRepository)(&client)

	repo.ProductExport = (*ProductExportRepository)(&client)

	repo.PromotionDiscountRule = (*PromotionDiscountRuleRepository)(&client)

	repo.RuleCondition = (*RuleConditionRepository)(&client)

	repo.TaxProviderTranslation = (*TaxProviderTranslationRepository)(&client)

	repo.WebhookEventLog = (*WebhookEventLogRepository)(&client)

	repo.CountryStateTranslation = (*CountryStateTranslationRepository)(&client)

	repo.MailTemplateTranslation = (*MailTemplateTranslationRepository)(&client)

	repo.NewsletterRecipientTag = (*NewsletterRecipientTagRepository)(&client)

	repo.SnippetSet = (*SnippetSetRepository)(&client)

	repo.UserAccessKey = (*UserAccessKeyRepository)(&client)

	repo.PropertyGroup = (*PropertyGroupRepository)(&client)

	repo.CustomerTag = (*CustomerTagRepository)(&client)

	repo.ImportExportProfileTranslation = (*ImportExportProfileTranslationRepository)(&client)

	repo.SalesChannelCountry = (*SalesChannelCountryRepository)(&client)

	repo.PromotionPersonaCustomer = (*PromotionPersonaCustomerRepository)(&client)

	repo.Version = (*VersionRepository)(&client)

	repo.Language = (*LanguageRepository)(&client)

	repo.PromotionSalesChannel = (*PromotionSalesChannelRepository)(&client)

	repo.Salutation = (*SalutationRepository)(&client)

	repo.StateMachine = (*StateMachineRepository)(&client)

	repo.StateMachineTranslation = (*StateMachineTranslationRepository)(&client)

	repo.UnitTranslation = (*UnitTranslationRepository)(&client)

	repo.ThemeChild = (*ThemeChildRepository)(&client)

	repo.AppScriptCondition = (*AppScriptConditionRepository)(&client)

	repo.CategoryTag = (*CategoryTagRepository)(&client)

	repo.LogEntry = (*LogEntryRepository)(&client)

	repo.MediaThumbnail = (*MediaThumbnailRepository)(&client)

	repo.OrderDelivery = (*OrderDeliveryRepository)(&client)

	repo.ShippingMethod = (*ShippingMethodRepository)(&client)

	repo.CmsPageTranslation = (*CmsPageTranslationRepository)(&client)

	repo.CmsSlot = (*CmsSlotRepository)(&client)

	repo.OrderTransaction = (*OrderTransactionRepository)(&client)

	repo.Plugin = (*PluginRepository)(&client)

	repo.OrderLineItemDownload = (*OrderLineItemDownloadRepository)(&client)

	repo.ProductOption = (*ProductOptionRepository)(&client)

	repo.ProductPrice = (*ProductPriceRepository)(&client)

	repo.ProductTranslation = (*ProductTranslationRepository)(&client)

	repo.TaxRuleTypeTranslation = (*TaxRuleTypeTranslationRepository)(&client)

	repo.ProductStreamTranslation = (*ProductStreamTranslationRepository)(&client)

	repo.Rule = (*RuleRepository)(&client)

	repo.SalesChannelLanguage = (*SalesChannelLanguageRepository)(&client)

	repo.PropertyGroupOption = (*PropertyGroupOptionRepository)(&client)

	repo.SalesChannelPaymentMethod = (*SalesChannelPaymentMethodRepository)(&client)

	repo.AppAdministrationSnippet = (*AppAdministrationSnippetRepository)(&client)

	repo.AppCmsBlockTranslation = (*AppCmsBlockTranslationRepository)(&client)

	repo.AppFlowActionTranslation = (*AppFlowActionTranslationRepository)(&client)

	repo.NumberRangeState = (*NumberRangeStateRepository)(&client)

	repo.PluginTranslation = (*PluginTranslationRepository)(&client)

	repo.PromotionOrderRule = (*PromotionOrderRuleRepository)(&client)

	repo.MediaThumbnailSize = (*MediaThumbnailSizeRepository)(&client)

	repo.ProductDownload = (*ProductDownloadRepository)(&client)

	repo.Tax = (*TaxRepository)(&client)

	repo.CustomerGroupTranslation = (*CustomerGroupTranslationRepository)(&client)

	repo.MailHeaderFooter = (*MailHeaderFooterRepository)(&client)

	repo.FlowSequence = (*FlowSequenceRepository)(&client)

	repo.Locale = (*LocaleRepository)(&client)

	repo.OrderCustomer = (*OrderCustomerRepository)(&client)

	repo.OrderDeliveryPosition = (*OrderDeliveryPositionRepository)(&client)

	repo.ProductCrossSellingTranslation = (*ProductCrossSellingTranslationRepository)(&client)

	repo.Media = (*MediaRepository)(&client)

	repo.Order = (*OrderRepository)(&client)

	repo.CmsPage = (*CmsPageRepository)(&client)

	repo.FlowTemplate = (*FlowTemplateRepository)(&client)

	repo.MailHeaderFooterTranslation = (*MailHeaderFooterTranslationRepository)(&client)

	repo.MediaFolder = (*MediaFolderRepository)(&client)

	repo.MediaTag = (*MediaTagRepository)(&client)

	repo.User = (*UserRepository)(&client)

	repo.Country = (*CountryRepository)(&client)

	repo.NumberRangeType = (*NumberRangeTypeRepository)(&client)

	repo.SalesChannel = (*SalesChannelRepository)(&client)

	repo.CustomerGroup = (*CustomerGroupRepository)(&client)

	repo.ProductCrossSellingAssignedProducts = (*ProductCrossSellingAssignedProductsRepository)(&client)

	repo.OrderLineItem = (*OrderLineItemRepository)(&client)

	repo.OrderTransactionCaptureRefund = (*OrderTransactionCaptureRefundRepository)(&client)

	repo.ProductCrossSelling = (*ProductCrossSellingRepository)(&client)

	repo.ProductManufacturer = (*ProductManufacturerRepository)(&client)

	repo.TaxProvider = (*TaxProviderRepository)(&client)

	repo.Flow = (*FlowRepository)(&client)

	repo.ImportExportProfile = (*ImportExportProfileRepository)(&client)

	repo.ProductVisibility = (*ProductVisibilityRepository)(&client)

	repo.PromotionDiscount = (*PromotionDiscountRepository)(&client)

	repo.ScheduledTask = (*ScheduledTaskRepository)(&client)

	repo.DocumentBaseConfigSalesChannel = (*DocumentBaseConfigSalesChannelRepository)(&client)

	repo.NumberRangeSalesChannel = (*NumberRangeSalesChannelRepository)(&client)

	repo.ShippingMethodTranslation = (*ShippingMethodTranslationRepository)(&client)

	repo.AppCmsBlock = (*AppCmsBlockRepository)(&client)

	repo.AppTemplate = (*AppTemplateRepository)(&client)

	repo.CustomFieldSet = (*CustomFieldSetRepository)(&client)

	repo.LandingPageSalesChannel = (*LandingPageSalesChannelRepository)(&client)

	repo.ProductTag = (*ProductTagRepository)(&client)

	repo.Script = (*ScriptRepository)(&client)

	repo.Customer = (*CustomerRepository)(&client)

	repo.SeoUrl = (*SeoUrlRepository)(&client)

	repo.ProductProperty = (*ProductPropertyRepository)(&client)

	repo.Snippet = (*SnippetRepository)(&client)

	repo.CurrencyCountryRounding = (*CurrencyCountryRoundingRepository)(&client)

	repo.MailTemplateTypeTranslation = (*MailTemplateTypeTranslationRepository)(&client)

	repo.TaxRule = (*TaxRuleRepository)(&client)

	repo.CmsSection = (*CmsSectionRepository)(&client)

	repo.MediaDefaultFolder = (*MediaDefaultFolderRepository)(&client)

	repo.MediaFolderConfiguration = (*MediaFolderConfigurationRepository)(&client)

	repo.MediaTranslation = (*MediaTranslationRepository)(&client)

	repo.ProductReview = (*ProductReviewRepository)(&client)

	repo.SalesChannelShippingMethod = (*SalesChannelShippingMethodRepository)(&client)

	repo.DeliveryTimeTranslation = (*DeliveryTimeTranslationRepository)(&client)

	repo.ProductConfiguratorSetting = (*ProductConfiguratorSettingRepository)(&client)

	repo.ProductSortingTranslation = (*ProductSortingTranslationRepository)(&client)

	repo.RuleTag = (*RuleTagRepository)(&client)

	repo.SalesChannelCurrency = (*SalesChannelCurrencyRepository)(&client)

	repo.DocumentBaseConfig = (*DocumentBaseConfigRepository)(&client)

	repo.PropertyGroupOptionTranslation = (*PropertyGroupOptionTranslationRepository)(&client)

	repo.CurrencyTranslation = (*CurrencyTranslationRepository)(&client)

	repo.LocaleTranslation = (*LocaleTranslationRepository)(&client)

	repo.ProductStreamFilter = (*ProductStreamFilterRepository)(&client)

	repo.ProductSearchKeyword = (*ProductSearchKeywordRepository)(&client)

	repo.StateMachineStateTranslation = (*StateMachineStateTranslationRepository)(&client)

	repo.ThemeTranslation = (*ThemeTranslationRepository)(&client)

	repo.NumberRangeTypeTranslation = (*NumberRangeTypeTranslationRepository)(&client)

	repo.PaymentMethodTranslation = (*PaymentMethodTranslationRepository)(&client)

	repo.VersionCommit = (*VersionCommitRepository)(&client)

	repo.Webhook = (*WebhookRepository)(&client)

	repo.ProductSearchConfig = (*ProductSearchConfigRepository)(&client)

	repo.AppActionButtonTranslation = (*AppActionButtonTranslationRepository)(&client)

	repo.CustomerAddress = (*CustomerAddressRepository)(&client)

	repo.DocumentTypeTranslation = (*DocumentTypeTranslationRepository)(&client)

	repo.Product = (*ProductRepository)(&client)

	repo.ProductCustomFieldSet = (*ProductCustomFieldSetRepository)(&client)

	repo.PromotionDiscountPrices = (*PromotionDiscountPricesRepository)(&client)

	repo.OrderTransactionCapture = (*OrderTransactionCaptureRepository)(&client)

	repo.PropertyGroupTranslation = (*PropertyGroupTranslationRepository)(&client)

	repo.CustomFieldSetRelation = (*CustomFieldSetRelationRepository)(&client)

	repo.CustomerGroupRegistrationSalesChannels = (*CustomerGroupRegistrationSalesChannelsRepository)(&client)

	repo.ProductFeatureSetTranslation = (*ProductFeatureSetTranslationRepository)(&client)

	repo.ShippingMethodPrice = (*ShippingMethodPriceRepository)(&client)

	repo.StateMachineTransition = (*StateMachineTransitionRepository)(&client)

	repo.ProductFeatureSet = (*ProductFeatureSetRepository)(&client)

	repo.DeliveryTime = (*DeliveryTimeRepository)(&client)

	repo.ProductKeywordDictionary = (*ProductKeywordDictionaryRepository)(&client)

	repo.ProductManufacturerTranslation = (*ProductManufacturerTranslationRepository)(&client)

	repo.PromotionCartRule = (*PromotionCartRuleRepository)(&client)

	repo.SalesChannelType = (*SalesChannelTypeRepository)(&client)

	repo.Document = (*DocumentRepository)(&client)

	repo.UserRecovery = (*UserRecoveryRepository)(&client)

	repo.PromotionPersonaRule = (*PromotionPersonaRuleRepository)(&client)

	repo.VersionCommitData = (*VersionCommitDataRepository)(&client)

	repo.AppTranslation = (*AppTranslationRepository)(&client)

	repo.PromotionIndividualCode = (*PromotionIndividualCodeRepository)(&client)

	repo.Unit = (*UnitRepository)(&client)

	repo.TaxRuleType = (*TaxRuleTypeRepository)(&client)

	repo.AclRole = (*AclRoleRepository)(&client)

	repo.CustomerWishlistProduct = (*CustomerWishlistProductRepository)(&client)

	repo.MailTemplateType = (*MailTemplateTypeRepository)(&client)

	repo.NumberRange = (*NumberRangeRepository)(&client)

	repo.NumberRangeTranslation = (*NumberRangeTranslationRepository)(&client)

	repo.StateMachineState = (*StateMachineStateRepository)(&client)

	repo.CustomField = (*CustomFieldRepository)(&client)

	repo.LandingPageTag = (*LandingPageTagRepository)(&client)

	repo.SalesChannelTypeTranslation = (*SalesChannelTypeTranslationRepository)(&client)

	repo.SalutationTranslation = (*SalutationTranslationRepository)(&client)

	repo.ShippingMethodTag = (*ShippingMethodTagRepository)(&client)

	repo.AppFlowAction = (*AppFlowActionRepository)(&client)

	repo.Integration = (*IntegrationRepository)(&client)

	repo.LandingPage = (*LandingPageRepository)(&client)

	repo.Tag = (*TagRepository)(&client)

	repo.AppPaymentMethod = (*AppPaymentMethodRepository)(&client)

	repo.CountryState = (*CountryStateRepository)(&client)

	repo.NewsletterRecipient = (*NewsletterRecipientRepository)(&client)

	repo.ProductMedia = (*ProductMediaRepository)(&client)

	repo.ThemeSalesChannel = (*ThemeSalesChannelRepository)(&client)

	repo.CountryTranslation = (*CountryTranslationRepository)(&client)

	repo.Currency = (*CurrencyRepository)(&client)

	repo.DocumentType = (*DocumentTypeRepository)(&client)

	repo.MainCategory = (*MainCategoryRepository)(&client)

	repo.ProductCategory = (*ProductCategoryRepository)(&client)

	repo.SystemConfig = (*SystemConfigRepository)(&client)

	repo.CmsBlock = (*CmsBlockRepository)(&client)

	repo.AclUserRole = (*AclUserRoleRepository)(&client)

	repo.CmsSlotTranslation = (*CmsSlotTranslationRepository)(&client)

	repo.SeoUrlTemplate = (*SeoUrlTemplateRepository)(&client)

	repo.Category = (*CategoryRepository)(&client)

	repo.ImportExportLog = (*ImportExportLogRepository)(&client)

	repo.MediaFolderConfigurationMediaThumbnailSize = (*MediaFolderConfigurationMediaThumbnailSizeRepository)(&client)

	repo.PaymentMethod = (*PaymentMethodRepository)(&client)

	repo.ProductStreamMapping = (*ProductStreamMappingRepository)(&client)

	repo.PromotionSetgroupRule = (*PromotionSetgroupRuleRepository)(&client)

	repo.AppShippingMethod = (*AppShippingMethodRepository)(&client)

	repo.ImportExportFile = (*ImportExportFileRepository)(&client)

	repo.OrderTransactionCaptureRefundPosition = (*OrderTransactionCaptureRefundPositionRepository)(&client)

	repo.Promotion = (*PromotionRepository)(&client)

	repo.SalesChannelDomain = (*SalesChannelDomainRepository)(&client)

	repo.IntegrationRole = (*IntegrationRoleRepository)(&client)

	repo.MailTemplate = (*MailTemplateRepository)(&client)

	repo.OrderAddress = (*OrderAddressRepository)(&client)

	repo.ProductSorting = (*ProductSortingRepository)(&client)

	repo.CustomEntity = (*CustomEntityRepository)(&client)

	repo.MailTemplateMedia = (*MailTemplateMediaRepository)(&client)

	repo.PromotionTranslation = (*PromotionTranslationRepository)(&client)

	return repo
}
