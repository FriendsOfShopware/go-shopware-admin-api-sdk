package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductKeywordDictionaryRepository struct {
	*GenericRepository[ProductKeywordDictionary]
}

func NewProductKeywordDictionaryRepository(client *Client) *ProductKeywordDictionaryRepository {
	return &ProductKeywordDictionaryRepository{
		GenericRepository: NewGenericRepository[ProductKeywordDictionary](client),
	}
}

func (t *ProductKeywordDictionaryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductKeywordDictionary], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-keyword-dictionary")
}

func (t *ProductKeywordDictionaryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductKeywordDictionary], *http.Response, error) {
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

func (t *ProductKeywordDictionaryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-keyword-dictionary")
}

func (t *ProductKeywordDictionaryRepository) Upsert(ctx ApiContext, entity []ProductKeywordDictionary) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_keyword_dictionary")
}

func (t *ProductKeywordDictionaryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_keyword_dictionary")
}

type ProductKeywordDictionary struct {

	Language      *Language  `json:"language,omitempty"`

	Id      string  `json:"id,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Keyword      string  `json:"keyword,omitempty"`

	Reversed      string  `json:"reversed,omitempty"`

}
