package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type UnitRepository struct {
	*GenericRepository[Unit]
}

func NewUnitRepository(client *Client) *UnitRepository {
	return &UnitRepository{
		GenericRepository: NewGenericRepository[Unit](client),
	}
}

func (t *UnitRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Unit], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "unit")
}

func (t *UnitRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Unit], *http.Response, error) {
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

func (t *UnitRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "unit")
}

func (t *UnitRepository) Upsert(ctx ApiContext, entity []Unit) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "unit")
}

func (t *UnitRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "unit")
}

type Unit struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	ShortCode      string  `json:"shortCode,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []UnitTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
