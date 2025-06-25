package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductExportRepository struct {
	*GenericRepository[ProductExport]
}

func NewProductExportRepository(client *Client) *ProductExportRepository {
	return &ProductExportRepository{
		GenericRepository: NewGenericRepository[ProductExport](client),
	}
}

func (t *ProductExportRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductExport], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-export")
}

func (t *ProductExportRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductExport], *http.Response, error) {
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

func (t *ProductExportRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-export")
}

func (t *ProductExportRepository) Upsert(ctx ApiContext, entity []ProductExport) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_export")
}

func (t *ProductExportRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_export")
}

type ProductExport struct {

	AccessKey      string  `json:"accessKey,omitempty"`

	BodyTemplate      string  `json:"bodyTemplate,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	Encoding      string  `json:"encoding,omitempty"`

	FileFormat      string  `json:"fileFormat,omitempty"`

	FileName      string  `json:"fileName,omitempty"`

	FooterTemplate      string  `json:"footerTemplate,omitempty"`

	GenerateByCronjob      bool  `json:"generateByCronjob,omitempty"`

	GeneratedAt      time.Time  `json:"generatedAt,omitempty"`

	HeaderTemplate      string  `json:"headerTemplate,omitempty"`

	Id      string  `json:"id,omitempty"`

	IncludeVariants      bool  `json:"includeVariants,omitempty"`

	Interval      float64  `json:"interval,omitempty"`

	IsRunning      bool  `json:"isRunning,omitempty"`

	PausedSchedule      bool  `json:"pausedSchedule,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelDomain      *SalesChannelDomain  `json:"salesChannelDomain,omitempty"`

	SalesChannelDomainId      string  `json:"salesChannelDomainId,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	StorefrontSalesChannel      *SalesChannel  `json:"storefrontSalesChannel,omitempty"`

	StorefrontSalesChannelId      string  `json:"storefrontSalesChannelId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
