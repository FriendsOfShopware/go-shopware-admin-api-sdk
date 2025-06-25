package go_shopware_admin_sdk

import (
	"fmt"
	"net/http"
)

type GenericRepository[T any, C any] struct {
	client *Client
}

func NewGenericRepository[T any, C any](client *Client) *GenericRepository[T, C] {
	return &GenericRepository[T, C]{
		client: client,
	}
}

func (r *GenericRepository[T, C]) Search(ctx ApiContext, criteria Criteria, entityName string) (*C, *http.Response, error) {
	req, err := r.client.NewRequest(ctx, "POST", fmt.Sprintf("/api/search/%s", entityName), criteria)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(C)
	resp, err := r.client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (r *GenericRepository[T, C]) SearchIds(ctx ApiContext, criteria Criteria, entityName string) (*SearchIdsResponse, *http.Response, error) {
	req, err := r.client.NewRequest(ctx, "POST", fmt.Sprintf("/api/search-ids/%s", entityName), criteria)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := r.client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (r *GenericRepository[T, C]) Upsert(ctx ApiContext, entity []T, entityName string) (*http.Response, error) {
	return r.client.Bulk.Sync(ctx, map[string]SyncOperation{entityName: {
		Entity:  entityName,
		Action:  "upsert",
		Payload: entity,
	}})
}

func (r *GenericRepository[T, C]) Delete(ctx ApiContext, ids []string, entityName string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return r.client.Bulk.Sync(ctx, map[string]SyncOperation{entityName: {
		Entity:  entityName,
		Action:  "delete",
		Payload: payload,
	}})
}
