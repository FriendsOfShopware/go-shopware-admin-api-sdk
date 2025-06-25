package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppAdministrationSnippetRepository struct {
	*GenericRepository[AppAdministrationSnippet]
}

func NewAppAdministrationSnippetRepository(client *Client) *AppAdministrationSnippetRepository {
	return &AppAdministrationSnippetRepository{
		GenericRepository: NewGenericRepository[AppAdministrationSnippet](client),
	}
}

func (t *AppAdministrationSnippetRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppAdministrationSnippet], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-administration-snippet")
}

func (t *AppAdministrationSnippetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppAdministrationSnippet], *http.Response, error) {
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

func (t *AppAdministrationSnippetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-administration-snippet")
}

func (t *AppAdministrationSnippetRepository) Upsert(ctx ApiContext, entity []AppAdministrationSnippet) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_administration_snippet")
}

func (t *AppAdministrationSnippetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_administration_snippet")
}

type AppAdministrationSnippet struct {

	AppId      string  `json:"appId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	LocaleId      string  `json:"localeId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Value      string  `json:"value,omitempty"`

}
