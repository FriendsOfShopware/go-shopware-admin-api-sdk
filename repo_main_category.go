package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MainCategoryRepository struct {
	*GenericRepository[MainCategory]
}

func NewMainCategoryRepository(client *Client) *MainCategoryRepository {
	return &MainCategoryRepository{
		GenericRepository: NewGenericRepository[MainCategory](client),
	}
}

func (t *MainCategoryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MainCategory], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "main-category")
}

func (t *MainCategoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MainCategory], *http.Response, error) {
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

func (t *MainCategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "main-category")
}

func (t *MainCategoryRepository) Upsert(ctx ApiContext, entity []MainCategory) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "main_category")
}

func (t *MainCategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "main_category")
}

type MainCategory struct {

	Category      *Category  `json:"category,omitempty"`

	CategoryId      string  `json:"categoryId,omitempty"`

	CategoryVersionId      string  `json:"categoryVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
