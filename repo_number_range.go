package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type NumberRangeRepository struct {
	*GenericRepository[NumberRange]
}

func NewNumberRangeRepository(client *Client) *NumberRangeRepository {
	return &NumberRangeRepository{
		GenericRepository: NewGenericRepository[NumberRange](client),
	}
}

func (t *NumberRangeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRange], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "number-range")
}

func (t *NumberRangeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRange], *http.Response, error) {
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

func (t *NumberRangeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "number-range")
}

func (t *NumberRangeRepository) Upsert(ctx ApiContext, entity []NumberRange) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "number_range")
}

func (t *NumberRangeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "number_range")
}

type NumberRange struct {

	TypeId      string  `json:"typeId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Pattern      string  `json:"pattern,omitempty"`

	Start      float64  `json:"start,omitempty"`

	NumberRangeSalesChannels      []NumberRangeSalesChannel  `json:"numberRangeSalesChannels,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Description      string  `json:"description,omitempty"`

	State      *NumberRangeState  `json:"state,omitempty"`

	Translations      []NumberRangeTranslation  `json:"translations,omitempty"`

	Id      string  `json:"id,omitempty"`

	Global      bool  `json:"global,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Type      *NumberRangeType  `json:"type,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}
