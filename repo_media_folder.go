package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MediaFolderRepository ClientService

func (t MediaFolderRepository) Search(ctx ApiContext, criteria Criteria) (*MediaFolderCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-folder", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaFolderCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaFolderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-folder", criteria)

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

func (t MediaFolderRepository) Upsert(ctx ApiContext, entity []MediaFolder) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_folder": {
		Entity:  "media_folder",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaFolderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_folder": {
		Entity:  "media_folder",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaFolder struct {
	ChildCount float64 `json:"childCount,omitempty"`

	Media []Media `json:"media,omitempty"`

	DefaultFolder *MediaDefaultFolder `json:"defaultFolder,omitempty"`

	Configuration *MediaFolderConfiguration `json:"configuration,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Parent *MediaFolder `json:"parent,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Children []MediaFolder `json:"children,omitempty"`

	Id string `json:"id,omitempty"`

	ConfigurationId string `json:"configurationId,omitempty"`

	DefaultFolderId string `json:"defaultFolderId,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UseParentConfiguration bool `json:"useParentConfiguration,omitempty"`
}

type MediaFolderCollection struct {
	EntityCollection

	Data []MediaFolder `json:"data"`
}
