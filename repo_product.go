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

func (t ProductRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCollection, *http.Response, error) {
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
	CrossSellingAssignedProducts []ProductCrossSellingAssignedProducts `json:"crossSellingAssignedProducts,omitempty"`

	UnitId string `json:"unitId,omitempty"`

	MaxPurchase float64 `json:"maxPurchase,omitempty"`

	OptionIds interface{} `json:"optionIds,omitempty"`

	Prices []ProductPrice `json:"prices,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	SearchKeywords []ProductSearchKeyword `json:"searchKeywords,omitempty"`

	OrderLineItems []OrderLineItem `json:"orderLineItems,omitempty"`

	Translations []ProductTranslation `json:"translations,omitempty"`

	Length float64 `json:"length,omitempty"`

	Sales float64 `json:"sales,omitempty"`

	Unit *Unit `json:"unit,omitempty"`

	FeatureSet *ProductFeatureSet `json:"featureSet,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Price interface{} `json:"price,omitempty"`

	Width float64 `json:"width,omitempty"`

	Name string `json:"name,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	CategoriesRo []Category `json:"categoriesRo,omitempty"`

	TaxId string `json:"taxId,omitempty"`

	ConfiguratorGroupConfig interface{} `json:"configuratorGroupConfig,omitempty"`

	ReleaseDate time.Time `json:"releaseDate,omitempty"`

	Parent *Product `json:"parent,omitempty"`

	ReferenceUnit float64 `json:"referenceUnit,omitempty"`

	Weight float64 `json:"weight,omitempty"`

	StreamIds interface{} `json:"streamIds,omitempty"`

	PackUnitPlural string `json:"packUnitPlural,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	CanonicalProductId string `json:"canonicalProductId,omitempty"`

	AvailableStock float64 `json:"availableStock,omitempty"`

	PurchaseSteps float64 `json:"purchaseSteps,omitempty"`

	Media []ProductMedia `json:"media,omitempty"`

	CrossSellings []ProductCrossSelling `json:"crossSellings,omitempty"`

	Options []PropertyGroupOption `json:"options,omitempty"`

	DeliveryTime *DeliveryTime `json:"deliveryTime,omitempty"`

	CanonicalProduct *Product `json:"canonicalProduct,omitempty"`

	ProductMediaVersionId string `json:"productMediaVersionId,omitempty"`

	Ean string `json:"ean,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	ShippingFree bool `json:"shippingFree,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ManufacturerId string `json:"manufacturerId,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	Stock float64 `json:"stock,omitempty"`

	ProductNumber string `json:"productNumber,omitempty"`

	Properties []PropertyGroupOption `json:"properties,omitempty"`

	ManufacturerNumber string `json:"manufacturerNumber,omitempty"`

	Streams []ProductStream `json:"streams,omitempty"`

	Id string `json:"id,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Active bool `json:"active,omitempty"`

	DisplayGroup string `json:"displayGroup,omitempty"`

	CustomFieldSets []CustomFieldSet `json:"customFieldSets,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	MainVariantId string `json:"mainVariantId,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	CategoryIds interface{} `json:"categoryIds,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	DeliveryTimeId string `json:"deliveryTimeId,omitempty"`

	Variation interface{} `json:"variation,omitempty"`

	MinPurchase float64 `json:"minPurchase,omitempty"`

	Height float64 `json:"height,omitempty"`

	CustomSearchKeywords interface{} `json:"customSearchKeywords,omitempty"`

	Children []Product `json:"children,omitempty"`

	Manufacturer *ProductManufacturer `json:"manufacturer,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	ProductManufacturerVersionId string `json:"productManufacturerVersionId,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	IsCloseout bool `json:"isCloseout,omitempty"`

	VariantRestrictions interface{} `json:"variantRestrictions,omitempty"`

	CustomFieldSetSelectionActive bool `json:"customFieldSetSelectionActive,omitempty"`

	Visibilities []ProductVisibility `json:"visibilities,omitempty"`

	CheapestPrice interface{} `json:"cheapestPrice,omitempty"`

	FeatureSetId string `json:"featureSetId,omitempty"`

	Available bool `json:"available,omitempty"`

	PurchasePrices interface{} `json:"purchasePrices,omitempty"`

	PropertyIds interface{} `json:"propertyIds,omitempty"`

	Description string `json:"description,omitempty"`

	PackUnit string `json:"packUnit,omitempty"`

	Tax *Tax `json:"tax,omitempty"`

	Cover *ProductMedia `json:"cover,omitempty"`

	CoverId string `json:"coverId,omitempty"`

	RestockTime float64 `json:"restockTime,omitempty"`

	MarkAsTopseller bool `json:"markAsTopseller,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	Wishlists []CustomerWishlistProduct `json:"wishlists,omitempty"`

	PurchaseUnit float64 `json:"purchaseUnit,omitempty"`

	RatingAverage float64 `json:"ratingAverage,omitempty"`

	CategoryTree interface{} `json:"categoryTree,omitempty"`

	ConfiguratorSettings []ProductConfiguratorSetting `json:"configuratorSettings,omitempty"`
}

type ProductCollection struct {
	EntityCollection

	Data []Product `json:"data"`
}
