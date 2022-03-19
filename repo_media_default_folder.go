package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MediaDefaultFolderRepository ClientService

func (t MediaDefaultFolderRepository) Search(ctx ApiContext, criteria Criteria) (*MediaDefaultFolderCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-default-folder", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaDefaultFolderCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaDefaultFolderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaDefaultFolderCollection, *http.Response, error) {
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

func (t MediaDefaultFolderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-default-folder", criteria)

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

func (t MediaDefaultFolderRepository) Upsert(ctx ApiContext, entity []MediaDefaultFolder) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_default_folder": {
		Entity:  "media_default_folder",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaDefaultFolderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_default_folder": {
		Entity:  "media_default_folder",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaDefaultFolder struct {
	Id string `json:"id,omitempty"`

	AssociationFields interface{} `json:"associationFields,omitempty"`

	Entity string `json:"entity,omitempty"`

	Folder *MediaFolder `json:"folder,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type MediaDefaultFolderCollection struct {
	EntityCollection

	Data []MediaDefaultFolder `json:"data"`
}
