package go_shopware_admin_sdk

import (
	"net/http"
)

type MediaFolderConfigurationMediaThumbnailSizeRepository ClientService

func (t MediaFolderConfigurationMediaThumbnailSizeRepository) Search(ctx ApiContext, criteria Criteria) (*MediaFolderConfigurationMediaThumbnailSizeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-folder-configuration-media-thumbnail-size", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaFolderConfigurationMediaThumbnailSizeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaFolderConfigurationMediaThumbnailSizeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaFolderConfigurationMediaThumbnailSizeCollection, *http.Response, error) {
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

func (t MediaFolderConfigurationMediaThumbnailSizeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-folder-configuration-media-thumbnail-size", criteria)

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

func (t MediaFolderConfigurationMediaThumbnailSizeRepository) Upsert(ctx ApiContext, entity []MediaFolderConfigurationMediaThumbnailSize) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_folder_configuration_media_thumbnail_size": {
		Entity:  "media_folder_configuration_media_thumbnail_size",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaFolderConfigurationMediaThumbnailSizeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_folder_configuration_media_thumbnail_size": {
		Entity:  "media_folder_configuration_media_thumbnail_size",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaFolderConfigurationMediaThumbnailSize struct {
	MediaThumbnailSize *MediaThumbnailSize `json:"mediaThumbnailSize,omitempty"`

	MediaFolderConfigurationId string `json:"mediaFolderConfigurationId,omitempty"`

	MediaThumbnailSizeId string `json:"mediaThumbnailSizeId,omitempty"`

	MediaFolderConfiguration *MediaFolderConfiguration `json:"mediaFolderConfiguration,omitempty"`
}

type MediaFolderConfigurationMediaThumbnailSizeCollection struct {
	EntityCollection

	Data []MediaFolderConfigurationMediaThumbnailSize `json:"data"`
}
