package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductSearchKeywordRepository struct {
	*GenericRepository[ProductSearchKeyword]
}

func NewProductSearchKeywordRepository(client *Client) *ProductSearchKeywordRepository {
	return &ProductSearchKeywordRepository{
		GenericRepository: NewGenericRepository[ProductSearchKeyword](client),
	}
}

func (t *ProductSearchKeywordRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSearchKeyword], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-search-keyword")
}

func (t *ProductSearchKeywordRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSearchKeyword], *http.Response, error) {
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

func (t *ProductSearchKeywordRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-search-keyword")
}

func (t *ProductSearchKeywordRepository) Upsert(ctx ApiContext, entity []ProductSearchKeyword) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_search_keyword")
}

func (t *ProductSearchKeywordRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_search_keyword")
}

type ProductSearchKeyword struct {

	LanguageId      string  `json:"languageId,omitempty"`

	Keyword      string  `json:"keyword,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Ranking      float64  `json:"ranking,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

}
