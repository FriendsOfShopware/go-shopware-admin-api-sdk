package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SnippetRepository struct {
	*GenericRepository[Snippet]
}

func NewSnippetRepository(client *Client) *SnippetRepository {
	return &SnippetRepository{
		GenericRepository: NewGenericRepository[Snippet](client),
	}
}

func (t *SnippetRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Snippet], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "snippet")
}

func (t *SnippetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Snippet], *http.Response, error) {
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

func (t *SnippetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "snippet")
}

func (t *SnippetRepository) Upsert(ctx ApiContext, entity []Snippet) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "snippet")
}

func (t *SnippetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "snippet")
}

type Snippet struct {

	TranslationKey      string  `json:"translationKey,omitempty"`

	Set      *SnippetSet  `json:"set,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	SetId      string  `json:"setId,omitempty"`

	Value      string  `json:"value,omitempty"`

	Author      string  `json:"author,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}
