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
	PausedSchedule bool `json:"pausedSchedule,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	SalesChannelDomainId string `json:"salesChannelDomainId,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	GenerateByCronjob bool `json:"generateByCronjob,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	FileName string `json:"fileName,omitempty"`

	Interval float64 `json:"interval,omitempty"`

	BodyTemplate string `json:"bodyTemplate,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`

	StorefrontSalesChannel *SalesChannel `json:"storefrontSalesChannel,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Id string `json:"id,omitempty"`

	FileFormat string `json:"fileFormat,omitempty"`

	GeneratedAt time.Time `json:"generatedAt,omitempty"`

	SalesChannelDomain *SalesChannelDomain `json:"salesChannelDomain,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	IncludeVariants bool `json:"includeVariants,omitempty"`

	HeaderTemplate string `json:"headerTemplate,omitempty"`

	FooterTemplate string `json:"footerTemplate,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`

	StorefrontSalesChannelId string `json:"storefrontSalesChannelId,omitempty"`

	Encoding string `json:"encoding,omitempty"`
}

type ProductExportCollection struct {
	EntityCollection

	Data []ProductExport `json:"data"`
}
