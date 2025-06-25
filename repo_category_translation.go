package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CategoryTranslationRepository struct {
	*GenericRepository[CategoryTranslation]
}

func NewCategoryTranslationRepository(client *Client) *CategoryTranslationRepository {
	return &CategoryTranslationRepository{
		GenericRepository: NewGenericRepository[CategoryTranslation](client),
	}
}

func (t *CategoryTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CategoryTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "category-translation")
}

func (t *CategoryTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CategoryTranslation], *http.Response, error) {
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

func (t *CategoryTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "category-translation")
}

func (t *CategoryTranslationRepository) Upsert(ctx ApiContext, entity []CategoryTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "category_translation")
}

func (t *CategoryTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "category_translation")
}

type CategoryTranslation struct {

	Breadcrumb      interface{}  `json:"breadcrumb,omitempty"`

	Category      *Category  `json:"category,omitempty"`

	CategoryId      string  `json:"categoryId,omitempty"`

	CategoryVersionId      string  `json:"categoryVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	ExternalLink      string  `json:"externalLink,omitempty"`

	InternalLink      string  `json:"internalLink,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	LinkNewTab      bool  `json:"linkNewTab,omitempty"`

	LinkType      string  `json:"linkType,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	Name      string  `json:"name,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
