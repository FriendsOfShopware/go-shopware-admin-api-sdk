package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductFeatureSetRepository ClientService

func (t ProductFeatureSetRepository) Search(ctx ApiContext, criteria Criteria) (*ProductFeatureSetCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-feature-set", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductFeatureSetCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductFeatureSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductFeatureSetCollection, *http.Response, error) {
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

func (t ProductFeatureSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-feature-set", criteria)

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

func (t ProductFeatureSetRepository) Upsert(ctx ApiContext, entity []ProductFeatureSet) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_feature_set": {
		Entity:  "product_feature_set",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductFeatureSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_feature_set": {
		Entity:  "product_feature_set",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductFeatureSet struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Features interface{} `json:"features,omitempty"`

	Products []Product `json:"products,omitempty"`

	Translations []ProductFeatureSetTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductFeatureSetCollection struct {
	EntityCollection

	Data []ProductFeatureSet `json:"data"`
}
