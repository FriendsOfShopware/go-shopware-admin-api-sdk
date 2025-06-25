package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductStreamTranslationRepository struct {
	*GenericRepository[ProductStreamTranslation]
}

func NewProductStreamTranslationRepository(client *Client) *ProductStreamTranslationRepository {
	return &ProductStreamTranslationRepository{
		GenericRepository: NewGenericRepository[ProductStreamTranslation](client),
	}
}

func (t *ProductStreamTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStreamTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-stream-translation")
}

func (t *ProductStreamTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStreamTranslation], *http.Response, error) {
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

func (t *ProductStreamTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-stream-translation")
}

func (t *ProductStreamTranslationRepository) Upsert(ctx ApiContext, entity []ProductStreamTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_stream_translation")
}

func (t *ProductStreamTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_stream_translation")
}

type ProductStreamTranslation struct {

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Language      *Language  `json:"language,omitempty"`

}
