package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type NumberRangeStateRepository struct {
	*GenericRepository[NumberRangeState]
}

func NewNumberRangeStateRepository(client *Client) *NumberRangeStateRepository {
	return &NumberRangeStateRepository{
		GenericRepository: NewGenericRepository[NumberRangeState](client),
	}
}

func (t *NumberRangeStateRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeState], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "number-range-state")
}

func (t *NumberRangeStateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeState], *http.Response, error) {
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

func (t *NumberRangeStateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "number-range-state")
}

func (t *NumberRangeStateRepository) Upsert(ctx ApiContext, entity []NumberRangeState) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "number_range_state")
}

func (t *NumberRangeStateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "number_range_state")
}

type NumberRangeState struct {

	Id      string  `json:"id,omitempty"`

	NumberRangeId      string  `json:"numberRangeId,omitempty"`

	LastValue      float64  `json:"lastValue,omitempty"`

	NumberRange      *NumberRange  `json:"numberRange,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
