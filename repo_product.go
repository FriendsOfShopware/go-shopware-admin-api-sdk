package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductRepository ClientService

func (t ProductRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product", criteria)

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

func (t ProductRepository) Upsert(ctx ApiContext, entity []Product) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product": {
		Entity:  "product",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product": {
		Entity:  "product",
		Action:  "delete",
		Payload: payload,
	}})
}

type Product struct {
	ManufacturerNumber string `json:"manufacturerNumber,omitempty"`

	MinPurchase float64 `json:"minPurchase,omitempty"`

	OptionIds interface{} `json:"optionIds,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`

	Id string `json:"id,omitempty"`

	ProductMediaVersionId string `json:"productMediaVersionId,omitempty"`

	IsCloseout bool `json:"isCloseout,omitempty"`

	MainVariantId string `json:"mainVariantId,omitempty"`

	PackUnitPlural string `json:"packUnitPlural,omitempty"`

	Streams []ProductStream `json:"streams,omitempty"`

	CategoriesRo []Category `json:"categoriesRo,omitempty"`

	CheapestPrice interface{} `json:"cheapestPrice,omitempty"`

	ProductNumber string `json:"productNumber,omitempty"`

	Width float64 `json:"width,omitempty"`

	CrossSellingAssignedProducts []ProductCrossSellingAssignedProducts `json:"crossSellingAssignedProducts,omitempty"`

	Properties []PropertyGroupOption `json:"properties,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	VariantRestrictions interface{} `json:"variantRestrictions,omitempty"`

	CustomSearchKeywords interface{} `json:"customSearchKeywords,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	Media []ProductMedia `json:"media,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	DeliveryTimeId string `json:"deliveryTimeId,omitempty"`

	Stock float64 `json:"stock,omitempty"`

	Variation interface{} `json:"variation,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	CustomFieldSets []CustomFieldSet `json:"customFieldSets,omitempty"`

	StreamIds interface{} `json:"streamIds,omitempty"`

	Prices []ProductPrice `json:"prices,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ShippingFree bool `json:"shippingFree,omitempty"`

	ReleaseDate time.Time `json:"releaseDate,omitempty"`

	CustomFieldSetSelectionActive bool `json:"customFieldSetSelectionActive,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Ean string `json:"ean,omitempty"`

	MarkAsTopseller bool `json:"markAsTopseller,omitempty"`

	RatingAverage float64 `json:"ratingAverage,omitempty"`

	Sales float64 `json:"sales,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ProductManufacturerVersionId string `json:"productManufacturerVersionId,omitempty"`

	CoverId string `json:"coverId,omitempty"`

	AvailableStock float64 `json:"availableStock,omitempty"`

	PackUnit string `json:"packUnit,omitempty"`

	Tax *Tax `json:"tax,omitempty"`

	SearchKeywords []ProductSearchKeyword `json:"searchKeywords,omitempty"`

	Options []PropertyGroupOption `json:"options,omitempty"`

	ConfiguratorGroupConfig interface{} `json:"configuratorGroupConfig,omitempty"`

	Manufacturer *ProductManufacturer `json:"manufacturer,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	DeliveryTime *DeliveryTime `json:"deliveryTime,omitempty"`

	CrossSellings []ProductCrossSelling `json:"crossSellings,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	ManufacturerId string `json:"manufacturerId,omitempty"`

	UnitId string `json:"unitId,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	Height float64 `json:"height,omitempty"`

	RestockTime float64 `json:"restockTime,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	ConfiguratorSettings []ProductConfiguratorSetting `json:"configuratorSettings,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	PurchasePrices interface{} `json:"purchasePrices,omitempty"`

	Length float64 `json:"length,omitempty"`

	PropertyIds interface{} `json:"propertyIds,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	TaxId string `json:"taxId,omitempty"`

	CanonicalProductId string `json:"canonicalProductId,omitempty"`

	FeatureSet *ProductFeatureSet `json:"featureSet,omitempty"`

	CanonicalProduct *Product `json:"canonicalProduct,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	OrderLineItems []OrderLineItem `json:"orderLineItems,omitempty"`

	PurchaseSteps float64 `json:"purchaseSteps,omitempty"`

	CategoryIds interface{} `json:"categoryIds,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Wishlists []CustomerWishlistProduct `json:"wishlists,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	CategoryTree interface{} `json:"categoryTree,omitempty"`

	Children []Product `json:"children,omitempty"`

	Translations []ProductTranslation `json:"translations,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	MaxPurchase float64 `json:"maxPurchase,omitempty"`

	ReferenceUnit float64 `json:"referenceUnit,omitempty"`

	Weight float64 `json:"weight,omitempty"`

	Price interface{} `json:"price,omitempty"`

	PurchaseUnit float64 `json:"purchaseUnit,omitempty"`

	Visibilities []ProductVisibility `json:"visibilities,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	FeatureSetId string `json:"featureSetId,omitempty"`

	Active bool `json:"active,omitempty"`

	DisplayGroup string `json:"displayGroup,omitempty"`

	Parent *Product `json:"parent,omitempty"`

	Available bool `json:"available,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Unit *Unit `json:"unit,omitempty"`

	Cover *ProductMedia `json:"cover,omitempty"`
}

type ProductCollection struct {
	EntityCollection

	Data []Product `json:"data"`
}
