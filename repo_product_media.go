package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductMediaRepository ClientService

func (t ProductMediaRepository) Search(ctx ApiContext, criteria Criteria) (*ProductMediaCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-media", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductMediaCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductMediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-media", criteria)

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

func (t ProductMediaRepository) Upsert(ctx ApiContext, entity []ProductMedia) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_media": {
		Entity:  "product_media",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductMediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_media": {
		Entity:  "product_media",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductMedia struct {
	ProductId string `json:"productId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Position float64 `json:"position,omitempty"`

	Media *Media `json:"media,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductMediaCollection struct {
	EntityCollection

	Data []ProductMedia `json:"data"`
}
