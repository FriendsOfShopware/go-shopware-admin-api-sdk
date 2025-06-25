package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppTemplateRepository struct {
	*GenericRepository[AppTemplate]
}

func NewAppTemplateRepository(client *Client) *AppTemplateRepository {
	return &AppTemplateRepository{
		GenericRepository: NewGenericRepository[AppTemplate](client),
	}
}

func (t *AppTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppTemplate], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-template")
}

func (t *AppTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppTemplate], *http.Response, error) {
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

func (t *AppTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-template")
}

func (t *AppTemplateRepository) Upsert(ctx ApiContext, entity []AppTemplate) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_template")
}

func (t *AppTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_template")
}

type AppTemplate struct {

	Template      string  `json:"template,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	App      *App  `json:"app,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Path      string  `json:"path,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Hash      string  `json:"hash,omitempty"`

}
