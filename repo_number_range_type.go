package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type NumberRangeTypeRepository struct {
	*GenericRepository[NumberRangeType]
}

func NewNumberRangeTypeRepository(client *Client) *NumberRangeTypeRepository {
	return &NumberRangeTypeRepository{
		GenericRepository: NewGenericRepository[NumberRangeType](client),
	}
}

func (t *NumberRangeTypeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeType], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "number-range-type")
}

func (t *NumberRangeTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeType], *http.Response, error) {
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

func (t *NumberRangeTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "number-range-type")
}

func (t *NumberRangeTypeRepository) Upsert(ctx ApiContext, entity []NumberRangeType) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "number_range_type")
}

func (t *NumberRangeTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "number_range_type")
}

type NumberRangeType struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	TypeName      string  `json:"typeName,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	NumberRanges      []NumberRange  `json:"numberRanges,omitempty"`

	Translations      []NumberRangeTypeTranslation  `json:"translations,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Global      bool  `json:"global,omitempty"`

	NumberRangeSalesChannels      []NumberRangeSalesChannel  `json:"numberRangeSalesChannels,omitempty"`

}
