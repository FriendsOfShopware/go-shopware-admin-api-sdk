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

	Active      bool  `json:"active,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	Available      bool  `json:"available,omitempty"`

	AvailableStock      float64  `json:"availableStock,omitempty"`

	CanonicalProduct      *Product  `json:"canonicalProduct,omitempty"`

	CanonicalProductId      string  `json:"canonicalProductId,omitempty"`

	CanonicalProductVersionId      string  `json:"canonicalProductVersionId,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	CategoriesRo      []Category  `json:"categoriesRo,omitempty"`

	CategoryIds      interface{}  `json:"categoryIds,omitempty"`

	CategoryTree      interface{}  `json:"categoryTree,omitempty"`

	ChildCount      float64  `json:"childCount,omitempty"`

	Children      []Product  `json:"children,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	ConfiguratorSettings      []ProductConfiguratorSetting  `json:"configuratorSettings,omitempty"`

	Cover      *ProductMedia  `json:"cover,omitempty"`

	CoverId      string  `json:"coverId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CrossSellingAssignedProducts      []ProductCrossSellingAssignedProducts  `json:"crossSellingAssignedProducts,omitempty"`

	CrossSellings      []ProductCrossSelling  `json:"crossSellings,omitempty"`

	CustomFieldSetSelectionActive      bool  `json:"customFieldSetSelectionActive,omitempty"`

	CustomFieldSets      []CustomFieldSet  `json:"customFieldSets,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomSearchKeywords      interface{}  `json:"customSearchKeywords,omitempty"`

	DeliveryTime      *DeliveryTime  `json:"deliveryTime,omitempty"`

	DeliveryTimeId      string  `json:"deliveryTimeId,omitempty"`

	Description      string  `json:"description,omitempty"`

	DisplayGroup      string  `json:"displayGroup,omitempty"`

	Downloads      []ProductDownload  `json:"downloads,omitempty"`

	Ean      string  `json:"ean,omitempty"`

	FeatureSet      *ProductFeatureSet  `json:"featureSet,omitempty"`

	FeatureSetId      string  `json:"featureSetId,omitempty"`

	Height      float64  `json:"height,omitempty"`

	Id      string  `json:"id,omitempty"`

	IsCloseout      bool  `json:"isCloseout,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	Length      float64  `json:"length,omitempty"`

	MainCategories      []MainCategory  `json:"mainCategories,omitempty"`

	Manufacturer      *ProductManufacturer  `json:"manufacturer,omitempty"`

	ManufacturerId      string  `json:"manufacturerId,omitempty"`

	ManufacturerNumber      string  `json:"manufacturerNumber,omitempty"`

	MarkAsTopseller      bool  `json:"markAsTopseller,omitempty"`

	MaxPurchase      float64  `json:"maxPurchase,omitempty"`

	Media      []ProductMedia  `json:"media,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	MinPurchase      float64  `json:"minPurchase,omitempty"`

	Name      string  `json:"name,omitempty"`

	OptionIds      interface{}  `json:"optionIds,omitempty"`

	Options      []PropertyGroupOption  `json:"options,omitempty"`

	OrderLineItems      []OrderLineItem  `json:"orderLineItems,omitempty"`

	PackUnit      string  `json:"packUnit,omitempty"`

	PackUnitPlural      string  `json:"packUnitPlural,omitempty"`

	Parent      *Product  `json:"parent,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	ParentVersionId      string  `json:"parentVersionId,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	Prices      []ProductPrice  `json:"prices,omitempty"`

	ProductManufacturerVersionId      string  `json:"productManufacturerVersionId,omitempty"`

	ProductMediaVersionId      string  `json:"productMediaVersionId,omitempty"`

	ProductNumber      string  `json:"productNumber,omitempty"`

	ProductReviews      []ProductReview  `json:"productReviews,omitempty"`

	Properties      []PropertyGroupOption  `json:"properties,omitempty"`

	PropertyIds      interface{}  `json:"propertyIds,omitempty"`

	PurchasePrices      interface{}  `json:"purchasePrices,omitempty"`

	PurchaseSteps      float64  `json:"purchaseSteps,omitempty"`

	PurchaseUnit      float64  `json:"purchaseUnit,omitempty"`

	RatingAverage      float64  `json:"ratingAverage,omitempty"`

	ReferenceUnit      float64  `json:"referenceUnit,omitempty"`

	ReleaseDate      time.Time  `json:"releaseDate,omitempty"`

	RestockTime      float64  `json:"restockTime,omitempty"`

	Sales      float64  `json:"sales,omitempty"`

	SearchKeywords      []ProductSearchKeyword  `json:"searchKeywords,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	ShippingFree      bool  `json:"shippingFree,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	States      interface{}  `json:"states,omitempty"`

	Stock      float64  `json:"stock,omitempty"`

	StreamIds      interface{}  `json:"streamIds,omitempty"`

	Streams      []ProductStream  `json:"streams,omitempty"`

	TagIds      interface{}  `json:"tagIds,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Tax      *Tax  `json:"tax,omitempty"`

	TaxId      string  `json:"taxId,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []ProductTranslation  `json:"translations,omitempty"`

	Unit      *Unit  `json:"unit,omitempty"`

	UnitId      string  `json:"unitId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VariantListingConfig      interface{}  `json:"variantListingConfig,omitempty"`

	VariantRestrictions      interface{}  `json:"variantRestrictions,omitempty"`

	Variation      interface{}  `json:"variation,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Visibilities      []ProductVisibility  `json:"visibilities,omitempty"`

	Weight      float64  `json:"weight,omitempty"`

	Width      float64  `json:"width,omitempty"`

	Wishlists      []CustomerWishlistProduct  `json:"wishlists,omitempty"`

}
