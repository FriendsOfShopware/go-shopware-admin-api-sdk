package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MediaFolderConfigurationRepository ClientService

func (t MediaFolderConfigurationRepository) Search(ctx ApiContext, criteria Criteria) (*MediaFolderConfigurationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-folder-configuration", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaFolderConfigurationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaFolderConfigurationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-folder-configuration", criteria)

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

func (t MediaFolderConfigurationRepository) Upsert(ctx ApiContext, entity []MediaFolderConfiguration) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_folder_configuration": {
		Entity:  "media_folder_configuration",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaFolderConfigurationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_folder_configuration": {
		Entity:  "media_folder_configuration",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaFolderConfiguration struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	NoAssociation bool `json:"noAssociation,omitempty"`

	MediaThumbnailSizes []MediaThumbnailSize `json:"mediaThumbnailSizes,omitempty"`

	Private bool `json:"private,omitempty"`

	MediaFolders []MediaFolder `json:"mediaFolders,omitempty"`

	MediaThumbnailSizesRo interface{} `json:"mediaThumbnailSizesRo,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CreateThumbnails bool `json:"createThumbnails,omitempty"`

	KeepAspectRatio bool `json:"keepAspectRatio,omitempty"`

	ThumbnailQuality float64 `json:"thumbnailQuality,omitempty"`
}

type MediaFolderConfigurationCollection struct {
	EntityCollection

	Data []MediaFolderConfiguration `json:"data"`
}
