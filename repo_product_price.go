package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type ProductPriceRepository ClientService

func (t ProductPriceRepository) Search(ctx ApiContext, criteria Criteria) (*ProductPriceCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-price", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductPriceCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductPriceRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductPriceCollection, *http.Response, error) {
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

func (t ProductPriceRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-price", criteria)

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

func (t ProductPriceRepository) Upsert(ctx ApiContext, entity []ProductPrice) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_price": {
		Entity:  "product_price",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductPriceRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_price": {
		Entity:  "product_price",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductPrice struct {
	ProductVersionId string `json:"productVersionId,omitempty"`

	Price interface{} `json:"price,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	QuantityEnd float64 `json:"quantityEnd,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ProductId string `json:"productId,omitempty"`

	QuantityStart float64 `json:"quantityStart,omitempty"`

	Product *Product `json:"product,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductPriceCollection struct {
	EntityCollection

	Data []ProductPrice `json:"data"`
}
