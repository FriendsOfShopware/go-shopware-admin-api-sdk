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

func (t MediaFolderConfigurationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaFolderConfigurationCollection, *http.Response, error) {
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
	CreateThumbnails bool `json:"createThumbnails,omitempty"`

	KeepAspectRatio bool `json:"keepAspectRatio,omitempty"`

	ThumbnailQuality float64 `json:"thumbnailQuality,omitempty"`

	NoAssociation bool `json:"noAssociation,omitempty"`

	MediaFolders []MediaFolder `json:"mediaFolders,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	MediaThumbnailSizes []MediaThumbnailSize `json:"mediaThumbnailSizes,omitempty"`

	MediaThumbnailSizesRo interface{} `json:"mediaThumbnailSizesRo,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Private bool `json:"private,omitempty"`
}

type MediaFolderConfigurationCollection struct {
	EntityCollection

	Data []MediaFolderConfiguration `json:"data"`
}
