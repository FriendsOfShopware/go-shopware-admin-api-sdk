package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SnippetSetRepository struct {
	*GenericRepository[SnippetSet]
}

func NewSnippetSetRepository(client *Client) *SnippetSetRepository {
	return &SnippetSetRepository{
		GenericRepository: NewGenericRepository[SnippetSet](client),
	}
}

func (t *SnippetSetRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SnippetSet], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "snippet-set")
}

func (t *SnippetSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SnippetSet], *http.Response, error) {
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

func (t *SnippetSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "snippet-set")
}

func (t *SnippetSetRepository) Upsert(ctx ApiContext, entity []SnippetSet) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "snippet_set")
}

func (t *SnippetSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "snippet_set")
}

type SnippetSet struct {

	BaseFile      string  `json:"baseFile,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Iso      string  `json:"iso,omitempty"`

	Name      string  `json:"name,omitempty"`

	SalesChannelDomains      []SalesChannelDomain  `json:"salesChannelDomains,omitempty"`

	Snippets      []Snippet  `json:"snippets,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
