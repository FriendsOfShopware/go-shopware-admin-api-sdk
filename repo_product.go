package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductRepository struct {
	*GenericRepository[Product]
}

func NewProductRepository(client *Client) *ProductRepository {
	return &ProductRepository{
		GenericRepository: NewGenericRepository[Product](client),
	}
}

func (t *ProductRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Product], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product")
}

func (t *ProductRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Product], *http.Response, error) {
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

func (t *ProductRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product")
}

func (t *ProductRepository) Upsert(ctx ApiContext, entity []Product) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product")
}

func (t *ProductRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product")
}

type Product struct {

	CategoryIds      interface{}  `json:"categoryIds,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	Cover      *ProductMedia  `json:"cover,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	Translations      []ProductTranslation  `json:"translations,omitempty"`

	RestockTime      float64  `json:"restockTime,omitempty"`

	VariantListingConfig      interface{}  `json:"variantListingConfig,omitempty"`

	Name      string  `json:"name,omitempty"`

	CrossSellings      []ProductCrossSelling  `json:"crossSellings,omitempty"`

	Streams      []ProductStream  `json:"streams,omitempty"`

	ProductManufacturerVersionId      string  `json:"productManufacturerVersionId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	ReferenceUnit      float64  `json:"referenceUnit,omitempty"`

	StreamIds      interface{}  `json:"streamIds,omitempty"`

	Description      string  `json:"description,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	OrderLineItems      []OrderLineItem  `json:"orderLineItems,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	CanonicalProductVersionId      string  `json:"canonicalProductVersionId,omitempty"`

	ParentVersionId      string  `json:"parentVersionId,omitempty"`

	IsCloseout      bool  `json:"isCloseout,omitempty"`

	Variation      interface{}  `json:"variation,omitempty"`

	MinPurchase      float64  `json:"minPurchase,omitempty"`

	Prices      []ProductPrice  `json:"prices,omitempty"`

	SearchKeywords      []ProductSearchKeyword  `json:"searchKeywords,omitempty"`

	Options      []PropertyGroupOption  `json:"options,omitempty"`

	ManufacturerId      string  `json:"manufacturerId,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Ean      string  `json:"ean,omitempty"`

	CustomFieldSetSelectionActive      bool  `json:"customFieldSetSelectionActive,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	CustomFieldSets      []CustomFieldSet  `json:"customFieldSets,omitempty"`

	UnitId      string  `json:"unitId,omitempty"`

	PurchasePrices      interface{}  `json:"purchasePrices,omitempty"`

	States      interface{}  `json:"states,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	CustomSearchKeywords      interface{}  `json:"customSearchKeywords,omitempty"`

	ConfiguratorSettings      []ProductConfiguratorSetting  `json:"configuratorSettings,omitempty"`

	Visibilities      []ProductVisibility  `json:"visibilities,omitempty"`

	Properties      []PropertyGroupOption  `json:"properties,omitempty"`

	TaxId      string  `json:"taxId,omitempty"`

	ProductNumber      string  `json:"productNumber,omitempty"`

	VariantRestrictions      interface{}  `json:"variantRestrictions,omitempty"`

	RatingAverage      float64  `json:"ratingAverage,omitempty"`

	Parent      *Product  `json:"parent,omitempty"`

	DeliveryTime      *DeliveryTime  `json:"deliveryTime,omitempty"`

	Manufacturer      *ProductManufacturer  `json:"manufacturer,omitempty"`

	CategoriesRo      []Category  `json:"categoriesRo,omitempty"`

	PropertyIds      interface{}  `json:"propertyIds,omitempty"`

	OptionIds      interface{}  `json:"optionIds,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	ShippingFree      bool  `json:"shippingFree,omitempty"`

	PackUnit      string  `json:"packUnit,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	Unit      *Unit  `json:"unit,omitempty"`

	Wishlists      []CustomerWishlistProduct  `json:"wishlists,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	DeliveryTimeId      string  `json:"deliveryTimeId,omitempty"`

	ManufacturerNumber      string  `json:"manufacturerNumber,omitempty"`

	CategoryTree      interface{}  `json:"categoryTree,omitempty"`

	Children      []Product  `json:"children,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	FeatureSetId      string  `json:"featureSetId,omitempty"`

	MaxPurchase      float64  `json:"maxPurchase,omitempty"`

	Downloads      []ProductDownload  `json:"downloads,omitempty"`

	Id      string  `json:"id,omitempty"`

	CoverId      string  `json:"coverId,omitempty"`

	ProductMediaVersionId      string  `json:"productMediaVersionId,omitempty"`

	AvailableStock      float64  `json:"availableStock,omitempty"`

	Stock      float64  `json:"stock,omitempty"`

	Width      float64  `json:"width,omitempty"`

	Sales      float64  `json:"sales,omitempty"`

	PackUnitPlural      string  `json:"packUnitPlural,omitempty"`

	Tax      *Tax  `json:"tax,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	CanonicalProductId      string  `json:"canonicalProductId,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	PurchaseSteps      float64  `json:"purchaseSteps,omitempty"`

	TagIds      interface{}  `json:"tagIds,omitempty"`

	Media      []ProductMedia  `json:"media,omitempty"`

	MarkAsTopseller      bool  `json:"markAsTopseller,omitempty"`

	ReleaseDate      time.Time  `json:"releaseDate,omitempty"`

	CanonicalProduct      *Product  `json:"canonicalProduct,omitempty"`

	CrossSellingAssignedProducts      []ProductCrossSellingAssignedProducts  `json:"crossSellingAssignedProducts,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	Available      bool  `json:"available,omitempty"`

	PurchaseUnit      float64  `json:"purchaseUnit,omitempty"`

	Weight      float64  `json:"weight,omitempty"`

	Height      float64  `json:"height,omitempty"`

	Length      float64  `json:"length,omitempty"`

	ChildCount      float64  `json:"childCount,omitempty"`

	FeatureSet      *ProductFeatureSet  `json:"featureSet,omitempty"`

	DisplayGroup      string  `json:"displayGroup,omitempty"`

}
