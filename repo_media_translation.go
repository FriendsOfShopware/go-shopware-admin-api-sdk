package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MediaTranslationRepository struct {
	*GenericRepository[MediaTranslation]
}

func NewMediaTranslationRepository(client *Client) *MediaTranslationRepository {
	return &MediaTranslationRepository{
		GenericRepository: NewGenericRepository[MediaTranslation](client),
	}
}

func (t *MediaTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "media-translation")
}

func (t *MediaTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MediaTranslation], *http.Response, error) {
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

func (t *MediaTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "media-translation")
}

func (t *MediaTranslationRepository) Upsert(ctx ApiContext, entity []MediaTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "media_translation")
}

func (t *MediaTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "media_translation")
}

type MediaTranslation struct {

	Alt      string  `json:"alt,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Title      string  `json:"title,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
