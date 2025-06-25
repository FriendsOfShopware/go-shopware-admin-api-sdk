package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type LocaleRepository struct {
	*GenericRepository[Locale]
}

func NewLocaleRepository(client *Client) *LocaleRepository {
	return &LocaleRepository{
		GenericRepository: NewGenericRepository[Locale](client),
	}
}

func (t *LocaleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Locale], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "locale")
}

func (t *LocaleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Locale], *http.Response, error) {
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

func (t *LocaleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "locale")
}

func (t *LocaleRepository) Upsert(ctx ApiContext, entity []Locale) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "locale")
}

func (t *LocaleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "locale")
}

type Locale struct {

	Code      string  `json:"code,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Languages      []Language  `json:"languages,omitempty"`

	Name      string  `json:"name,omitempty"`

	Territory      string  `json:"territory,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []LocaleTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Users      []User  `json:"users,omitempty"`

}
