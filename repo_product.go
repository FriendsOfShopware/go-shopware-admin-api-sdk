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
	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	Cover *ProductMedia `json:"cover,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	ManufacturerId string `json:"manufacturerId,omitempty"`

	TaxId string `json:"taxId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	CustomFieldSetSelectionActive bool `json:"customFieldSetSelectionActive,omitempty"`

	Manufacturer *ProductManufacturer `json:"manufacturer,omitempty"`

	PurchaseSteps float64 `json:"purchaseSteps,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	States interface{} `json:"states,omitempty"`

	Wishlists []CustomerWishlistProduct `json:"wishlists,omitempty"`

	Children []Product `json:"children,omitempty"`

	Prices []ProductPrice `json:"prices,omitempty"`

	Streams []ProductStream `json:"streams,omitempty"`

	Stock float64 `json:"stock,omitempty"`

	Available bool `json:"available,omitempty"`

	Sales float64 `json:"sales,omitempty"`

	VariantRestrictions interface{} `json:"variantRestrictions,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Description string `json:"description,omitempty"`

	Price interface{} `json:"price,omitempty"`

	ProductNumber string `json:"productNumber,omitempty"`

	StreamIds interface{} `json:"streamIds,omitempty"`

	Unit *Unit `json:"unit,omitempty"`

	OptionIds interface{} `json:"optionIds,omitempty"`

	DisplayGroup string `json:"displayGroup,omitempty"`

	DeliveryTime *DeliveryTime `json:"deliveryTime,omitempty"`

	ProductManufacturerVersionId string `json:"productManufacturerVersionId,omitempty"`

	Height float64 `json:"height,omitempty"`

	RatingAverage float64 `json:"ratingAverage,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	ShippingFree bool `json:"shippingFree,omitempty"`

	CustomSearchKeywords interface{} `json:"customSearchKeywords,omitempty"`

	Ean string `json:"ean,omitempty"`

	Parent *Product `json:"parent,omitempty"`

	ConfiguratorSettings []ProductConfiguratorSetting `json:"configuratorSettings,omitempty"`

	Translations []ProductTranslation `json:"translations,omitempty"`

	AvailableStock float64 `json:"availableStock,omitempty"`

	Width float64 `json:"width,omitempty"`

	PropertyIds interface{} `json:"propertyIds,omitempty"`

	PackUnit string `json:"packUnit,omitempty"`

	MainCategories []MainCategory `json:"mainCategories,omitempty"`

	VariantListingConfig interface{} `json:"variantListingConfig,omitempty"`

	OrderLineItems []OrderLineItem `json:"orderLineItems,omitempty"`

	Id string `json:"id,omitempty"`

	CategoryIds interface{} `json:"categoryIds,omitempty"`

	ReferenceUnit float64 `json:"referenceUnit,omitempty"`

	Weight float64 `json:"weight,omitempty"`

	CoverId string `json:"coverId,omitempty"`

	RestockTime float64 `json:"restockTime,omitempty"`

	CanonicalProduct *Product `json:"canonicalProduct,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	PackUnitPlural string `json:"packUnitPlural,omitempty"`

	FeatureSet *ProductFeatureSet `json:"featureSet,omitempty"`

	Media []ProductMedia `json:"media,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CanonicalProductId string `json:"canonicalProductId,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	ProductReviews []ProductReview `json:"productReviews,omitempty"`

	Options []PropertyGroupOption `json:"options,omitempty"`

	CategoriesRo []Category `json:"categoriesRo,omitempty"`

	MarkAsTopseller bool `json:"markAsTopseller,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Downloads []ProductDownload `json:"downloads,omitempty"`

	Tax *Tax `json:"tax,omitempty"`

	CategoryTree interface{} `json:"categoryTree,omitempty"`

	CustomFieldSets []CustomFieldSet `json:"customFieldSets,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	ManufacturerNumber string `json:"manufacturerNumber,omitempty"`

	DeliveryTimeId string `json:"deliveryTimeId,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	IsCloseout bool `json:"isCloseout,omitempty"`

	Variation interface{} `json:"variation,omitempty"`

	Length float64 `json:"length,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	ProductMediaVersionId string `json:"productMediaVersionId,omitempty"`

	Properties []PropertyGroupOption `json:"properties,omitempty"`

	Visibilities []ProductVisibility `json:"visibilities,omitempty"`

	SearchKeywords []ProductSearchKeyword `json:"searchKeywords,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`

	CrossSellings []ProductCrossSelling `json:"crossSellings,omitempty"`

	UnitId string `json:"unitId,omitempty"`

	PurchaseUnit float64 `json:"purchaseUnit,omitempty"`

	Name string `json:"name,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	FeatureSetId string `json:"featureSetId,omitempty"`

	ReleaseDate time.Time `json:"releaseDate,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	MinPurchase float64 `json:"minPurchase,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	MaxPurchase float64 `json:"maxPurchase,omitempty"`

	PurchasePrices interface{} `json:"purchasePrices,omitempty"`

	CrossSellingAssignedProducts []ProductCrossSellingAssignedProducts `json:"crossSellingAssignedProducts,omitempty"`

	Active bool `json:"active,omitempty"`
}

type ProductCollection struct {
	EntityCollection

	Data []Product `json:"data"`
}
