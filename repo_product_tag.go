package go_shopware_admin_sdk

import (
	"net/http"
)

type ProductTagRepository ClientService

func (t ProductTagRepository) Search(ctx ApiContext, criteria Criteria) (*ProductTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-tag", criteria)

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

func (t ProductTagRepository) Upsert(ctx ApiContext, entity []ProductTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_tag": {
		Entity:  "product_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_tag": {
		Entity:  "product_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductTag struct {
	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	TagId string `json:"tagId,omitempty"`

	Product *Product `json:"product,omitempty"`

	Tag *Tag `json:"tag,omitempty"`
}

type ProductTagCollection struct {
	EntityCollection

	Data []ProductTag `json:"data"`
}
