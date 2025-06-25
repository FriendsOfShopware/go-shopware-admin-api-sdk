package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MediaDefaultFolderRepository struct {
	*GenericRepository[MediaDefaultFolder]
}

func NewMediaDefaultFolderRepository(client *Client) *MediaDefaultFolderRepository {
	return &MediaDefaultFolderRepository{
		GenericRepository: NewGenericRepository[MediaDefaultFolder](client),
	}
}

func (t *MediaDefaultFolderRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaDefaultFolder], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-default-folder")
}

func (t *MediaDefaultFolderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaDefaultFolder], *http.Response, error) {
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

func (t *MediaDefaultFolderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-default-folder")
}

func (t *MediaDefaultFolderRepository) Upsert(ctx ApiContext, entity []MediaDefaultFolder) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_default_folder")
}

func (t *MediaDefaultFolderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_default_folder")
}

type MediaDefaultFolder struct {

	Id      string  `json:"id,omitempty"`

	Entity      string  `json:"entity,omitempty"`

	Folder      *MediaFolder  `json:"folder,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
