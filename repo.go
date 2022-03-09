package go_shopware_admin_sdk

type Repository struct {
	ClientService

	Currency *CurrencyRepository

	OrderCustomer *OrderCustomerRepository

	PromotionSetgroupRule *PromotionSetgroupRuleRepository

	Webhook *WebhookRepository

	AppCmsBlockTranslation *AppCmsBlockTranslationRepository

	CategoryTranslation *CategoryTranslationRepository

	ProductManufacturerTranslation *ProductManufacturerTranslationRepository

	Promotion *PromotionRepository

	App *AppRepository

	Locale *LocaleRepository

	LandingPageTranslation *LandingPageTranslationRepository

	Language *LanguageRepository

	PluginTranslation *PluginTranslationRepository

	SalesChannelCurrency *SalesChannelCurrencyRepository

	SalesChannelTranslation *SalesChannelTranslationRepository

	SalesChannelTypeTranslation *SalesChannelTypeTranslationRepository

	AppActionButton *AppActionButtonRepository

	CustomerGroupTranslation *CustomerGroupTranslationRepository

	ShippingMethodTranslation *ShippingMethodTranslationRepository

	StateMachine *StateMachineRepository

	ProductProperty *ProductPropertyRepository

	ProductSearchConfig *ProductSearchConfigRepository

	PropertyGroupTranslation *PropertyGroupTranslationRepository

	RuleCondition *RuleConditionRepository

	StateMachineState *StateMachineStateRepository

	CmsSection *CmsSectionRepository

	PaymentMethod *PaymentMethodRepository

	PropertyGroupOptionTranslation *PropertyGroupOptionTranslationRepository

	ProductReview *ProductReviewRepository

	ProductStreamMapping *ProductStreamMappingRepository

	WebhookEventLog *WebhookEventLogRepository

	MediaTranslation *MediaTranslationRepository

	OrderAddress *OrderAddressRepository

	MediaThumbnailSize *MediaThumbnailSizeRepository

	AppTemplate *AppTemplateRepository

	Flow *FlowRepository

	SalesChannelType *SalesChannelTypeRepository

	VersionCommit *VersionCommitRepository

	NumberRangeSalesChannel *NumberRangeSalesChannelRepository

	Product *ProductRepository

	EventActionSalesChannel *EventActionSalesChannelRepository

	ImportExportProfileTranslation *ImportExportProfileTranslationRepository

	MailHeaderFooterTranslation *MailHeaderFooterTranslationRepository

	ProductStreamFilter *ProductStreamFilterRepository

	PropertyGroupOption *PropertyGroupOptionRepository

	SalesChannelDomain *SalesChannelDomainRepository

	AppActionButtonTranslation *AppActionButtonTranslationRepository

	CustomFieldSetRelation *CustomFieldSetRelationRepository

	SalutationTranslation *SalutationTranslationRepository

	ShippingMethodTag *ShippingMethodTagRepository

	DocumentTypeTranslation *DocumentTypeTranslationRepository

	MediaThumbnail *MediaThumbnailRepository

	ProductCustomFieldSet *ProductCustomFieldSetRepository

	ProductKeywordDictionary *ProductKeywordDictionaryRepository

	Script *ScriptRepository

	CustomerRecovery *CustomerRecoveryRepository

	DocumentBaseConfig *DocumentBaseConfigRepository

	ProductCrossSellingAssignedProducts *ProductCrossSellingAssignedProductsRepository

	ProductPrice *ProductPriceRepository

	SalesChannelCountry *SalesChannelCountryRepository

	SystemConfig *SystemConfigRepository

	TaxRuleTypeTranslation *TaxRuleTypeTranslationRepository

	CmsPage *CmsPageRepository

	NumberRangeType *NumberRangeTypeRepository

	LogEntry *LogEntryRepository

	PromotionCartRule *PromotionCartRuleRepository

	SeoUrlTemplate *SeoUrlTemplateRepository

	ShippingMethodPrice *ShippingMethodPriceRepository

	Notification *NotificationRepository

	ProductCrossSellingTranslation *ProductCrossSellingTranslationRepository

	NumberRange *NumberRangeRepository

	OrderDelivery *OrderDeliveryRepository

	ProductCategory *ProductCategoryRepository

	ProductSortingTranslation *ProductSortingTranslationRepository

	DeliveryTimeTranslation *DeliveryTimeTranslationRepository

	NewsletterRecipientTag *NewsletterRecipientTagRepository

	MailHeaderFooter *MailHeaderFooterRepository

	ScheduledTask *ScheduledTaskRepository

	Theme *ThemeRepository

	AppPaymentMethod *AppPaymentMethodRepository

	CustomerAddress *CustomerAddressRepository

	Rule *RuleRepository

	Version *VersionRepository

	MessageQueueStats *MessageQueueStatsRepository

	OrderTransaction *OrderTransactionRepository

	OrderDeliveryPosition *OrderDeliveryPositionRepository

	ProductOption *ProductOptionRepository

	ShippingMethod *ShippingMethodRepository

	Unit *UnitRepository

	User *UserRepository

	EventAction *EventActionRepository

	LocaleTranslation *LocaleTranslationRepository

	Media *MediaRepository

	MediaFolderConfigurationMediaThumbnailSize *MediaFolderConfigurationMediaThumbnailSizeRepository

	ProductTag *ProductTagRepository

	ProductTranslation *ProductTranslationRepository

	PromotionSalesChannel *PromotionSalesChannelRepository

	DocumentBaseConfigSalesChannel *DocumentBaseConfigSalesChannelRepository

	EventActionRule *EventActionRuleRepository

	CustomFieldSet *CustomFieldSetRepository

	Integration *IntegrationRepository

	ProductMedia *ProductMediaRepository

	ProductStream *ProductStreamRepository

	ProductStreamTranslation *ProductStreamTranslationRepository

	PromotionPersonaRule *PromotionPersonaRuleRepository

	AclRole *AclRoleRepository

	CountryState *CountryStateRepository

	StateMachineStateTranslation *StateMachineStateTranslationRepository

	TaxRuleType *TaxRuleTypeRepository

	ThemeSalesChannel *ThemeSalesChannelRepository

	VersionCommitData *VersionCommitDataRepository

	SalesChannel *SalesChannelRepository

	SnippetSet *SnippetSetRepository

	ProductSorting *ProductSortingRepository

	SalesChannelShippingMethod *SalesChannelShippingMethodRepository

	AppCmsBlock *AppCmsBlockRepository

	Country *CountryRepository

	SalesChannelAnalytics *SalesChannelAnalyticsRepository

	TaxRule *TaxRuleRepository

	DeliveryTime *DeliveryTimeRepository

	MainCategory *MainCategoryRepository

	DocumentType *DocumentTypeRepository

	MailTemplateMedia *MailTemplateMediaRepository

	ProductFeatureSetTranslation *ProductFeatureSetTranslationRepository

	Tag *TagRepository

	UserRecovery *UserRecoveryRepository

	CmsPageTranslation *CmsPageTranslationRepository

	CustomField *CustomFieldRepository

	CustomerGroupRegistrationSalesChannels *CustomerGroupRegistrationSalesChannelsRepository

	MailTemplateTypeTranslation *MailTemplateTypeTranslationRepository

	NumberRangeState *NumberRangeStateRepository

	Plugin *PluginRepository

	ProductCrossSelling *ProductCrossSellingRepository

	ProductExport *ProductExportRepository

	Category *CategoryRepository

	CmsSlotTranslation *CmsSlotTranslationRepository

	ProductSearchKeyword *ProductSearchKeywordRepository

	PropertyGroup *PropertyGroupRepository

	NumberRangeTranslation *NumberRangeTranslationRepository

	ProductFeatureSet *ProductFeatureSetRepository

	PromotionIndividualCode *PromotionIndividualCodeRepository

	PromotionTranslation *PromotionTranslationRepository

	StateMachineTranslation *StateMachineTranslationRepository

	ThemeMedia *ThemeMediaRepository

	CountryTranslation *CountryTranslationRepository

	LandingPage *LandingPageRepository

	UserAccessKey *UserAccessKeyRepository

	Document *DocumentRepository

	IntegrationRole *IntegrationRoleRepository

	MailTemplateType *MailTemplateTypeRepository

	Order *OrderRepository

	PaymentMethodTranslation *PaymentMethodTranslationRepository

	PromotionDiscount *PromotionDiscountRepository

	CountryStateTranslation *CountryStateTranslationRepository

	CustomerGroup *CustomerGroupRepository

	PromotionSetgroup *PromotionSetgroupRepository

	UnitTranslation *UnitTranslationRepository

	MailTemplateTranslation *MailTemplateTranslationRepository

	ProductCategoryTree *ProductCategoryTreeRepository

	ProductSearchConfigField *ProductSearchConfigFieldRepository

	SeoUrl *SeoUrlRepository

	AclUserRole *AclUserRoleRepository

	ImportExportProfile *ImportExportProfileRepository

	OrderLineItem *OrderLineItemRepository

	ProductManufacturer *ProductManufacturerRepository

	CurrencyTranslation *CurrencyTranslationRepository

	CustomerWishlist *CustomerWishlistRepository

	MediaTag *MediaTagRepository

	PromotionOrderRule *PromotionOrderRuleRepository

	SalesChannelLanguage *SalesChannelLanguageRepository

	Salutation *SalutationRepository

	Snippet *SnippetRepository

	Tax *TaxRepository

	CustomerWishlistProduct *CustomerWishlistProductRepository

	ImportExportLog *ImportExportLogRepository

	DeadMessage *DeadMessageRepository

	LandingPageTag *LandingPageTagRepository

	MediaDefaultFolder *MediaDefaultFolderRepository

	NumberRangeTypeTranslation *NumberRangeTypeTranslationRepository

	ProductVisibility *ProductVisibilityRepository

	StateMachineTransition *StateMachineTransitionRepository

	AppTranslation *AppTranslationRepository

	CategoryTag *CategoryTagRepository

	ImportExportFile *ImportExportFileRepository

	MailTemplate *MailTemplateRepository

	MediaFolderConfiguration *MediaFolderConfigurationRepository

	SalesChannelPaymentMethod *SalesChannelPaymentMethodRepository

	StateMachineHistory *StateMachineHistoryRepository

	ThemeTranslation *ThemeTranslationRepository

	CurrencyCountryRounding *CurrencyCountryRoundingRepository

	CustomerTag *CustomerTagRepository

	LandingPageSalesChannel *LandingPageSalesChannelRepository

	OrderTag *OrderTagRepository

	ProductConfiguratorSetting *ProductConfiguratorSettingRepository

	PromotionDiscountPrices *PromotionDiscountPricesRepository

	PromotionPersonaCustomer *PromotionPersonaCustomerRepository

	UserConfig *UserConfigRepository

	CmsSlot *CmsSlotRepository

	Customer *CustomerRepository

	MediaFolder *MediaFolderRepository

	NewsletterRecipient *NewsletterRecipientRepository

	PromotionDiscountRule *PromotionDiscountRuleRepository

	CmsBlock *CmsBlockRepository

	FlowSequence *FlowSequenceRepository
}

func NewRepository(client ClientService) Repository {
	repo := Repository{
		ClientService: client,
	}

	repo.Currency = (*CurrencyRepository)(&client)

	repo.OrderCustomer = (*OrderCustomerRepository)(&client)

	repo.PromotionSetgroupRule = (*PromotionSetgroupRuleRepository)(&client)

	repo.Webhook = (*WebhookRepository)(&client)

	repo.AppCmsBlockTranslation = (*AppCmsBlockTranslationRepository)(&client)

	repo.CategoryTranslation = (*CategoryTranslationRepository)(&client)

	repo.ProductManufacturerTranslation = (*ProductManufacturerTranslationRepository)(&client)

	repo.Promotion = (*PromotionRepository)(&client)

	repo.App = (*AppRepository)(&client)

	repo.Locale = (*LocaleRepository)(&client)

	repo.LandingPageTranslation = (*LandingPageTranslationRepository)(&client)

	repo.Language = (*LanguageRepository)(&client)

	repo.PluginTranslation = (*PluginTranslationRepository)(&client)

	repo.SalesChannelCurrency = (*SalesChannelCurrencyRepository)(&client)

	repo.SalesChannelTranslation = (*SalesChannelTranslationRepository)(&client)

	repo.SalesChannelTypeTranslation = (*SalesChannelTypeTranslationRepository)(&client)

	repo.AppActionButton = (*AppActionButtonRepository)(&client)

	repo.CustomerGroupTranslation = (*CustomerGroupTranslationRepository)(&client)

	repo.ShippingMethodTranslation = (*ShippingMethodTranslationRepository)(&client)

	repo.StateMachine = (*StateMachineRepository)(&client)

	repo.ProductProperty = (*ProductPropertyRepository)(&client)

	repo.ProductSearchConfig = (*ProductSearchConfigRepository)(&client)

	repo.PropertyGroupTranslation = (*PropertyGroupTranslationRepository)(&client)

	repo.RuleCondition = (*RuleConditionRepository)(&client)

	repo.StateMachineState = (*StateMachineStateRepository)(&client)

	repo.CmsSection = (*CmsSectionRepository)(&client)

	repo.PaymentMethod = (*PaymentMethodRepository)(&client)

	repo.PropertyGroupOptionTranslation = (*PropertyGroupOptionTranslationRepository)(&client)

	repo.ProductReview = (*ProductReviewRepository)(&client)

	repo.ProductStreamMapping = (*ProductStreamMappingRepository)(&client)

	repo.WebhookEventLog = (*WebhookEventLogRepository)(&client)

	repo.MediaTranslation = (*MediaTranslationRepository)(&client)

	repo.OrderAddress = (*OrderAddressRepository)(&client)

	repo.MediaThumbnailSize = (*MediaThumbnailSizeRepository)(&client)

	repo.AppTemplate = (*AppTemplateRepository)(&client)

	repo.Flow = (*FlowRepository)(&client)

	repo.SalesChannelType = (*SalesChannelTypeRepository)(&client)

	repo.VersionCommit = (*VersionCommitRepository)(&client)

	repo.NumberRangeSalesChannel = (*NumberRangeSalesChannelRepository)(&client)

	repo.Product = (*ProductRepository)(&client)

	repo.EventActionSalesChannel = (*EventActionSalesChannelRepository)(&client)

	repo.ImportExportProfileTranslation = (*ImportExportProfileTranslationRepository)(&client)

	repo.MailHeaderFooterTranslation = (*MailHeaderFooterTranslationRepository)(&client)

	repo.ProductStreamFilter = (*ProductStreamFilterRepository)(&client)

	repo.PropertyGroupOption = (*PropertyGroupOptionRepository)(&client)

	repo.SalesChannelDomain = (*SalesChannelDomainRepository)(&client)

	repo.AppActionButtonTranslation = (*AppActionButtonTranslationRepository)(&client)

	repo.CustomFieldSetRelation = (*CustomFieldSetRelationRepository)(&client)

	repo.SalutationTranslation = (*SalutationTranslationRepository)(&client)

	repo.ShippingMethodTag = (*ShippingMethodTagRepository)(&client)

	repo.DocumentTypeTranslation = (*DocumentTypeTranslationRepository)(&client)

	repo.MediaThumbnail = (*MediaThumbnailRepository)(&client)

	repo.ProductCustomFieldSet = (*ProductCustomFieldSetRepository)(&client)

	repo.ProductKeywordDictionary = (*ProductKeywordDictionaryRepository)(&client)

	repo.Script = (*ScriptRepository)(&client)

	repo.CustomerRecovery = (*CustomerRecoveryRepository)(&client)

	repo.DocumentBaseConfig = (*DocumentBaseConfigRepository)(&client)

	repo.ProductCrossSellingAssignedProducts = (*ProductCrossSellingAssignedProductsRepository)(&client)

	repo.ProductPrice = (*ProductPriceRepository)(&client)

	repo.SalesChannelCountry = (*SalesChannelCountryRepository)(&client)

	repo.SystemConfig = (*SystemConfigRepository)(&client)

	repo.TaxRuleTypeTranslation = (*TaxRuleTypeTranslationRepository)(&client)

	repo.CmsPage = (*CmsPageRepository)(&client)

	repo.NumberRangeType = (*NumberRangeTypeRepository)(&client)

	repo.LogEntry = (*LogEntryRepository)(&client)

	repo.PromotionCartRule = (*PromotionCartRuleRepository)(&client)

	repo.SeoUrlTemplate = (*SeoUrlTemplateRepository)(&client)

	repo.ShippingMethodPrice = (*ShippingMethodPriceRepository)(&client)

	repo.Notification = (*NotificationRepository)(&client)

	repo.ProductCrossSellingTranslation = (*ProductCrossSellingTranslationRepository)(&client)

	repo.NumberRange = (*NumberRangeRepository)(&client)

	repo.OrderDelivery = (*OrderDeliveryRepository)(&client)

	repo.ProductCategory = (*ProductCategoryRepository)(&client)

	repo.ProductSortingTranslation = (*ProductSortingTranslationRepository)(&client)

	repo.DeliveryTimeTranslation = (*DeliveryTimeTranslationRepository)(&client)

	repo.NewsletterRecipientTag = (*NewsletterRecipientTagRepository)(&client)

	repo.MailHeaderFooter = (*MailHeaderFooterRepository)(&client)

	repo.ScheduledTask = (*ScheduledTaskRepository)(&client)

	repo.Theme = (*ThemeRepository)(&client)

	repo.AppPaymentMethod = (*AppPaymentMethodRepository)(&client)

	repo.CustomerAddress = (*CustomerAddressRepository)(&client)

	repo.Rule = (*RuleRepository)(&client)

	repo.Version = (*VersionRepository)(&client)

	repo.MessageQueueStats = (*MessageQueueStatsRepository)(&client)

	repo.OrderTransaction = (*OrderTransactionRepository)(&client)

	repo.OrderDeliveryPosition = (*OrderDeliveryPositionRepository)(&client)

	repo.ProductOption = (*ProductOptionRepository)(&client)

	repo.ShippingMethod = (*ShippingMethodRepository)(&client)

	repo.Unit = (*UnitRepository)(&client)

	repo.User = (*UserRepository)(&client)

	repo.EventAction = (*EventActionRepository)(&client)

	repo.LocaleTranslation = (*LocaleTranslationRepository)(&client)

	repo.Media = (*MediaRepository)(&client)

	repo.MediaFolderConfigurationMediaThumbnailSize = (*MediaFolderConfigurationMediaThumbnailSizeRepository)(&client)

	repo.ProductTag = (*ProductTagRepository)(&client)

	repo.ProductTranslation = (*ProductTranslationRepository)(&client)

	repo.PromotionSalesChannel = (*PromotionSalesChannelRepository)(&client)

	repo.DocumentBaseConfigSalesChannel = (*DocumentBaseConfigSalesChannelRepository)(&client)

	repo.EventActionRule = (*EventActionRuleRepository)(&client)

	repo.CustomFieldSet = (*CustomFieldSetRepository)(&client)

	repo.Integration = (*IntegrationRepository)(&client)

	repo.ProductMedia = (*ProductMediaRepository)(&client)

	repo.ProductStream = (*ProductStreamRepository)(&client)

	repo.ProductStreamTranslation = (*ProductStreamTranslationRepository)(&client)

	repo.PromotionPersonaRule = (*PromotionPersonaRuleRepository)(&client)

	repo.AclRole = (*AclRoleRepository)(&client)

	repo.CountryState = (*CountryStateRepository)(&client)

	repo.StateMachineStateTranslation = (*StateMachineStateTranslationRepository)(&client)

	repo.TaxRuleType = (*TaxRuleTypeRepository)(&client)

	repo.ThemeSalesChannel = (*ThemeSalesChannelRepository)(&client)

	repo.VersionCommitData = (*VersionCommitDataRepository)(&client)

	repo.SalesChannel = (*SalesChannelRepository)(&client)

	repo.SnippetSet = (*SnippetSetRepository)(&client)

	repo.ProductSorting = (*ProductSortingRepository)(&client)

	repo.SalesChannelShippingMethod = (*SalesChannelShippingMethodRepository)(&client)

	repo.AppCmsBlock = (*AppCmsBlockRepository)(&client)

	repo.Country = (*CountryRepository)(&client)

	repo.SalesChannelAnalytics = (*SalesChannelAnalyticsRepository)(&client)

	repo.TaxRule = (*TaxRuleRepository)(&client)

	repo.DeliveryTime = (*DeliveryTimeRepository)(&client)

	repo.MainCategory = (*MainCategoryRepository)(&client)

	repo.DocumentType = (*DocumentTypeRepository)(&client)

	repo.MailTemplateMedia = (*MailTemplateMediaRepository)(&client)

	repo.ProductFeatureSetTranslation = (*ProductFeatureSetTranslationRepository)(&client)

	repo.Tag = (*TagRepository)(&client)

	repo.UserRecovery = (*UserRecoveryRepository)(&client)

	repo.CmsPageTranslation = (*CmsPageTranslationRepository)(&client)

	repo.CustomField = (*CustomFieldRepository)(&client)

	repo.CustomerGroupRegistrationSalesChannels = (*CustomerGroupRegistrationSalesChannelsRepository)(&client)

	repo.MailTemplateTypeTranslation = (*MailTemplateTypeTranslationRepository)(&client)

	repo.NumberRangeState = (*NumberRangeStateRepository)(&client)

	repo.Plugin = (*PluginRepository)(&client)

	repo.ProductCrossSelling = (*ProductCrossSellingRepository)(&client)

	repo.ProductExport = (*ProductExportRepository)(&client)

	repo.Category = (*CategoryRepository)(&client)

	repo.CmsSlotTranslation = (*CmsSlotTranslationRepository)(&client)

	repo.ProductSearchKeyword = (*ProductSearchKeywordRepository)(&client)

	repo.PropertyGroup = (*PropertyGroupRepository)(&client)

	repo.NumberRangeTranslation = (*NumberRangeTranslationRepository)(&client)

	repo.ProductFeatureSet = (*ProductFeatureSetRepository)(&client)

	repo.PromotionIndividualCode = (*PromotionIndividualCodeRepository)(&client)

	repo.PromotionTranslation = (*PromotionTranslationRepository)(&client)

	repo.StateMachineTranslation = (*StateMachineTranslationRepository)(&client)

	repo.ThemeMedia = (*ThemeMediaRepository)(&client)

	repo.CountryTranslation = (*CountryTranslationRepository)(&client)

	repo.LandingPage = (*LandingPageRepository)(&client)

	repo.UserAccessKey = (*UserAccessKeyRepository)(&client)

	repo.Document = (*DocumentRepository)(&client)

	repo.IntegrationRole = (*IntegrationRoleRepository)(&client)

	repo.MailTemplateType = (*MailTemplateTypeRepository)(&client)

	repo.Order = (*OrderRepository)(&client)

	repo.PaymentMethodTranslation = (*PaymentMethodTranslationRepository)(&client)

	repo.PromotionDiscount = (*PromotionDiscountRepository)(&client)

	repo.CountryStateTranslation = (*CountryStateTranslationRepository)(&client)

	repo.CustomerGroup = (*CustomerGroupRepository)(&client)

	repo.PromotionSetgroup = (*PromotionSetgroupRepository)(&client)

	repo.UnitTranslation = (*UnitTranslationRepository)(&client)

	repo.MailTemplateTranslation = (*MailTemplateTranslationRepository)(&client)

	repo.ProductCategoryTree = (*ProductCategoryTreeRepository)(&client)

	repo.ProductSearchConfigField = (*ProductSearchConfigFieldRepository)(&client)

	repo.SeoUrl = (*SeoUrlRepository)(&client)

	repo.AclUserRole = (*AclUserRoleRepository)(&client)

	repo.ImportExportProfile = (*ImportExportProfileRepository)(&client)

	repo.OrderLineItem = (*OrderLineItemRepository)(&client)

	repo.ProductManufacturer = (*ProductManufacturerRepository)(&client)

	repo.CurrencyTranslation = (*CurrencyTranslationRepository)(&client)

	repo.CustomerWishlist = (*CustomerWishlistRepository)(&client)

	repo.MediaTag = (*MediaTagRepository)(&client)

	repo.PromotionOrderRule = (*PromotionOrderRuleRepository)(&client)

	repo.SalesChannelLanguage = (*SalesChannelLanguageRepository)(&client)

	repo.Salutation = (*SalutationRepository)(&client)

	repo.Snippet = (*SnippetRepository)(&client)

	repo.Tax = (*TaxRepository)(&client)

	repo.CustomerWishlistProduct = (*CustomerWishlistProductRepository)(&client)

	repo.ImportExportLog = (*ImportExportLogRepository)(&client)

	repo.DeadMessage = (*DeadMessageRepository)(&client)

	repo.LandingPageTag = (*LandingPageTagRepository)(&client)

	repo.MediaDefaultFolder = (*MediaDefaultFolderRepository)(&client)

	repo.NumberRangeTypeTranslation = (*NumberRangeTypeTranslationRepository)(&client)

	repo.ProductVisibility = (*ProductVisibilityRepository)(&client)

	repo.StateMachineTransition = (*StateMachineTransitionRepository)(&client)

	repo.AppTranslation = (*AppTranslationRepository)(&client)

	repo.CategoryTag = (*CategoryTagRepository)(&client)

	repo.ImportExportFile = (*ImportExportFileRepository)(&client)

	repo.MailTemplate = (*MailTemplateRepository)(&client)

	repo.MediaFolderConfiguration = (*MediaFolderConfigurationRepository)(&client)

	repo.SalesChannelPaymentMethod = (*SalesChannelPaymentMethodRepository)(&client)

	repo.StateMachineHistory = (*StateMachineHistoryRepository)(&client)

	repo.ThemeTranslation = (*ThemeTranslationRepository)(&client)

	repo.CurrencyCountryRounding = (*CurrencyCountryRoundingRepository)(&client)

	repo.CustomerTag = (*CustomerTagRepository)(&client)

	repo.LandingPageSalesChannel = (*LandingPageSalesChannelRepository)(&client)

	repo.OrderTag = (*OrderTagRepository)(&client)

	repo.ProductConfiguratorSetting = (*ProductConfiguratorSettingRepository)(&client)

	repo.PromotionDiscountPrices = (*PromotionDiscountPricesRepository)(&client)

	repo.PromotionPersonaCustomer = (*PromotionPersonaCustomerRepository)(&client)

	repo.UserConfig = (*UserConfigRepository)(&client)

	repo.CmsSlot = (*CmsSlotRepository)(&client)

	repo.Customer = (*CustomerRepository)(&client)

	repo.MediaFolder = (*MediaFolderRepository)(&client)

	repo.NewsletterRecipient = (*NewsletterRecipientRepository)(&client)

	repo.PromotionDiscountRule = (*PromotionDiscountRuleRepository)(&client)

	repo.CmsBlock = (*CmsBlockRepository)(&client)

	repo.FlowSequence = (*FlowSequenceRepository)(&client)

	return repo
}
