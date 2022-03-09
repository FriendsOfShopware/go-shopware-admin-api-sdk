package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductSearchKeywordRepository ClientService

func (t ProductSearchKeywordRepository) Search(ctx ApiContext, criteria Criteria) (*ProductSearchKeywordCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-search-keyword", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductSearchKeywordCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductSearchKeywordRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-search-keyword", criteria)

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

func (t ProductSearchKeywordRepository) Upsert(ctx ApiContext, entity []ProductSearchKeyword) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_keyword": {
		Entity:  "product_search_keyword",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductSearchKeywordRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_keyword": {
		Entity:  "product_search_keyword",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductSearchKeyword struct {
	Product *Product `json:"product,omitempty"`

	Language *Language `json:"language,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Keyword string `json:"keyword,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Ranking float64 `json:"ranking,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ProductSearchKeywordCollection struct {
	EntityCollection

	Data []ProductSearchKeyword `json:"data"`
}
