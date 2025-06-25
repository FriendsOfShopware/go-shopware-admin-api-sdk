package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MediaFolderRepository struct {
	*GenericRepository[MediaFolder]
}

func NewMediaFolderRepository(client *Client) *MediaFolderRepository {
	return &MediaFolderRepository{
		GenericRepository: NewGenericRepository[MediaFolder](client),
	}
}

func (t *MediaFolderRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaFolder], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-folder")
}

func (t *MediaFolderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaFolder], *http.Response, error) {
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

func (t *MediaFolderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-folder")
}

func (t *MediaFolderRepository) Upsert(ctx ApiContext, entity []MediaFolder) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_folder")
}

func (t *MediaFolderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_folder")
}

type MediaFolder struct {

	DefaultFolderId      string  `json:"defaultFolderId,omitempty"`

	Children      []MediaFolder  `json:"children,omitempty"`

	Configuration      *MediaFolderConfiguration  `json:"configuration,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UseParentConfiguration      bool  `json:"useParentConfiguration,omitempty"`

	ConfigurationId      string  `json:"configurationId,omitempty"`

	Path      string  `json:"path,omitempty"`

	Parent      *MediaFolder  `json:"parent,omitempty"`

	ChildCount      float64  `json:"childCount,omitempty"`

	Media      []Media  `json:"media,omitempty"`

	DefaultFolder      *MediaDefaultFolder  `json:"defaultFolder,omitempty"`

	Name      string  `json:"name,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}
