package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type ProductExportRepository ClientService

func (t ProductExportRepository) Search(ctx ApiContext, criteria Criteria) (*ProductExportCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-export", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductExportCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductExportRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductExportCollection, *http.Response, error) {
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

func (t ProductExportRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-export", criteria)

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

func (t ProductExportRepository) Upsert(ctx ApiContext, entity []ProductExport) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_export": {
		Entity:  "product_export",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductExportRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_export": {
		Entity:  "product_export",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductExport struct {
	Id string `json:"id,omitempty"`

	StorefrontSalesChannelId string `json:"storefrontSalesChannelId,omitempty"`

	FooterTemplate string `json:"footerTemplate,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	SalesChannelDomainId string `json:"salesChannelDomainId,omitempty"`

	Encoding string `json:"encoding,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	GeneratedAt time.Time `json:"generatedAt,omitempty"`

	IsRunning bool `json:"isRunning,omitempty"`

	BodyTemplate string `json:"bodyTemplate,omitempty"`

	PausedSchedule bool `json:"pausedSchedule,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	FileFormat string `json:"fileFormat,omitempty"`

	SalesChannelDomain *SalesChannelDomain `json:"salesChannelDomain,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	FileName string `json:"fileName,omitempty"`

	IncludeVariants bool `json:"includeVariants,omitempty"`

	GenerateByCronjob bool `json:"generateByCronjob,omitempty"`

	HeaderTemplate string `json:"headerTemplate,omitempty"`

	StorefrontSalesChannel *SalesChannel `json:"storefrontSalesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Interval float64 `json:"interval,omitempty"`
}

type ProductExportCollection struct {
	EntityCollection

	Data []ProductExport `json:"data"`
}
