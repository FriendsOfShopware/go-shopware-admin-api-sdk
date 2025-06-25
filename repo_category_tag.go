package go_shopware_admin_sdk

import (
	"net/http"

)

type CategoryTagRepository struct {
	*GenericRepository[CategoryTag]
}

func NewCategoryTagRepository(client *Client) *CategoryTagRepository {
	return &CategoryTagRepository{
		GenericRepository: NewGenericRepository[CategoryTag](client),
	}
}

func (t *CategoryTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CategoryTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "category-tag")
}

func (t *CategoryTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CategoryTag], *http.Response, error) {
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

func (t *CategoryTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "category-tag")
}

func (t *CategoryTagRepository) Upsert(ctx ApiContext, entity []CategoryTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "category_tag")
}

func (t *CategoryTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "category_tag")
}

type CategoryTag struct {

	Tag      *Tag  `json:"tag,omitempty"`

	CategoryId      string  `json:"categoryId,omitempty"`

	CategoryVersionId      string  `json:"categoryVersionId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	Category      *Category  `json:"category,omitempty"`

}
