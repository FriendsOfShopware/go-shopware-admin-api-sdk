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

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	RatingAverage      float64  `json:"ratingAverage,omitempty"`

	CustomFieldSets      []CustomFieldSet  `json:"customFieldSets,omitempty"`

	ChildCount      float64  `json:"childCount,omitempty"`

	Unit      *Unit  `json:"unit,omitempty"`

	Streams      []ProductStream  `json:"streams,omitempty"`

	Cover      *ProductMedia  `json:"cover,omitempty"`

	FeatureSet      *ProductFeatureSet  `json:"featureSet,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	FeatureSetId      string  `json:"featureSetId,omitempty"`

	ProductNumber      string  `json:"productNumber,omitempty"`

	Active      bool  `json:"active,omitempty"`

	MinPurchase      float64  `json:"minPurchase,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	Translations      []ProductTranslation  `json:"translations,omitempty"`

	TaxId      string  `json:"taxId,omitempty"`

	DisplayGroup      string  `json:"displayGroup,omitempty"`

	CustomFieldSetSelectionActive      bool  `json:"customFieldSetSelectionActive,omitempty"`

	Sales      float64  `json:"sales,omitempty"`

	IsCloseout      bool  `json:"isCloseout,omitempty"`

	MarkAsTopseller      bool  `json:"markAsTopseller,omitempty"`

	PurchasePrices      interface{}  `json:"purchasePrices,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ReferenceUnit      float64  `json:"referenceUnit,omitempty"`

	ManufacturerId      string  `json:"manufacturerId,omitempty"`

	RestockTime      float64  `json:"restockTime,omitempty"`

	ProductManufacturerVersionId      string  `json:"productManufacturerVersionId,omitempty"`

	PropertyIds      interface{}  `json:"propertyIds,omitempty"`

	StreamIds      interface{}  `json:"streamIds,omitempty"`

	Properties      []PropertyGroupOption  `json:"properties,omitempty"`

	VariantRestrictions      interface{}  `json:"variantRestrictions,omitempty"`

	Height      float64  `json:"height,omitempty"`

	Options      []PropertyGroupOption  `json:"options,omitempty"`

	ManufacturerNumber      string  `json:"manufacturerNumber,omitempty"`

	Ean      string  `json:"ean,omitempty"`

	Weight      float64  `json:"weight,omitempty"`

	OptionIds      interface{}  `json:"optionIds,omitempty"`

	Parent      *Product  `json:"parent,omitempty"`

	CrossSellings      []ProductCrossSelling  `json:"crossSellings,omitempty"`

	CategoriesRo      []Category  `json:"categoriesRo,omitempty"`

	ParentVersionId      string  `json:"parentVersionId,omitempty"`

	ProductMediaVersionId      string  `json:"productMediaVersionId,omitempty"`

	ConfiguratorSettings      []ProductConfiguratorSetting  `json:"configuratorSettings,omitempty"`

	SearchKeywords      []ProductSearchKeyword  `json:"searchKeywords,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	DeliveryTimeId      string  `json:"deliveryTimeId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	States      interface{}  `json:"states,omitempty"`

	CustomSearchKeywords      interface{}  `json:"customSearchKeywords,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	UnitId      string  `json:"unitId,omitempty"`

	Manufacturer      *ProductManufacturer  `json:"manufacturer,omitempty"`

	Prices      []ProductPrice  `json:"prices,omitempty"`

	Downloads      []ProductDownload  `json:"downloads,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	AvailableStock      float64  `json:"availableStock,omitempty"`

	Stock      float64  `json:"stock,omitempty"`

	MaxPurchase      float64  `json:"maxPurchase,omitempty"`

	Description      string  `json:"description,omitempty"`

	Name      string  `json:"name,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CategoryIds      interface{}  `json:"categoryIds,omitempty"`

	CoverId      string  `json:"coverId,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	Media      []ProductMedia  `json:"media,omitempty"`

	Available      bool  `json:"available,omitempty"`

	CrossSellingAssignedProducts      []ProductCrossSellingAssignedProducts  `json:"crossSellingAssignedProducts,omitempty"`

	VariantListingConfig      interface{}  `json:"variantListingConfig,omitempty"`

	Tax      *Tax  `json:"tax,omitempty"`

	PurchaseUnit      float64  `json:"purchaseUnit,omitempty"`

	ShippingFree      bool  `json:"shippingFree,omitempty"`

	DeliveryTime      *DeliveryTime  `json:"deliveryTime,omitempty"`

	Visibilities      []ProductVisibility  `json:"visibilities,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	OrderLineItems      []OrderLineItem  `json:"orderLineItems,omitempty"`

	CanonicalProductId      string  `json:"canonicalProductId,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	PurchaseSteps      float64  `json:"purchaseSteps,omitempty"`

	ReleaseDate      time.Time  `json:"releaseDate,omitempty"`

	PackUnitPlural      string  `json:"packUnitPlural,omitempty"`

	PackUnit      string  `json:"packUnit,omitempty"`

	CanonicalProduct      *Product  `json:"canonicalProduct,omitempty"`

	TagIds      interface{}  `json:"tagIds,omitempty"`

	Children      []Product  `json:"children,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Id      string  `json:"id,omitempty"`

	Width      float64  `json:"width,omitempty"`

	Wishlists      []CustomerWishlistProduct  `json:"wishlists,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	Variation      interface{}  `json:"variation,omitempty"`

	Length      float64  `json:"length,omitempty"`

	CategoryTree      interface{}  `json:"categoryTree,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

}

type ProductCollection struct {
	EntityCollection

	Data []Product `json:"data"`
}
