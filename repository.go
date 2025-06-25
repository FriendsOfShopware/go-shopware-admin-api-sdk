package go_shopware_admin_sdk

import (
	"fmt"
	"net/http"
)

type GenericRepository[T any] struct {
	Client *Client
}

func NewGenericRepository[T any](client *Client) *GenericRepository[T] {
	return &GenericRepository[T]{
		Client: client,
	}
}

func (r *GenericRepository[T]) Search(ctx ApiContext, criteria Criteria, entityName string) (*EntityCollection[T], *http.Response, error) {
	req, err := r.Client.NewRequest(ctx, "POST", fmt.Sprintf("/api/search/%s", entityName), criteria)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(EntityCollection[T])
	resp, err := r.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (r *GenericRepository[T]) SearchIds(ctx ApiContext, criteria Criteria, entityName string) (*SearchIdsResponse, *http.Response, error) {
	req, err := r.Client.NewRequest(ctx, "POST", fmt.Sprintf("/api/search-ids/%s", entityName), criteria)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := r.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (r *GenericRepository[T]) Upsert(ctx ApiContext, entity []T, entityName string) (*http.Response, error) {
	return r.Client.Bulk.Sync(ctx, map[string]SyncOperation{entityName: {
		Entity:  entityName,
		Action:  "upsert",
		Payload: entity,
	}})
}

func (r *GenericRepository[T]) Delete(ctx ApiContext, ids []string, entityName string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return r.Client.Bulk.Sync(ctx, map[string]SyncOperation{entityName: {
		Entity:  entityName,
		Action:  "delete",
		Payload: payload,
	}})
}
