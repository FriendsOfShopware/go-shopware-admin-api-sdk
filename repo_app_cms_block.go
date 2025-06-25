package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppCmsBlockRepository struct {
	*GenericRepository[AppCmsBlock]
}

func NewAppCmsBlockRepository(client *Client) *AppCmsBlockRepository {
	return &AppCmsBlockRepository{
		GenericRepository: NewGenericRepository[AppCmsBlock](client),
	}
}

func (t *AppCmsBlockRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppCmsBlock], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-cms-block")
}

func (t *AppCmsBlockRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppCmsBlock], *http.Response, error) {
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

func (t *AppCmsBlockRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-cms-block")
}

func (t *AppCmsBlockRepository) Upsert(ctx ApiContext, entity []AppCmsBlock) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_cms_block")
}

func (t *AppCmsBlockRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_cms_block")
}

type AppCmsBlock struct {

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	Block      interface{}  `json:"block,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Label      string  `json:"label,omitempty"`

	Name      string  `json:"name,omitempty"`

	Styles      string  `json:"styles,omitempty"`

	Template      string  `json:"template,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []AppCmsBlockTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
