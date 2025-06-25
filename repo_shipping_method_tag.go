package go_shopware_admin_sdk

import (
	"net/http"

)

type ShippingMethodTagRepository struct {
	*GenericRepository[ShippingMethodTag]
}

func NewShippingMethodTagRepository(client *Client) *ShippingMethodTagRepository {
	return &ShippingMethodTagRepository{
		GenericRepository: NewGenericRepository[ShippingMethodTag](client),
	}
}

func (t *ShippingMethodTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethodTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "shipping-method-tag")
}

func (t *ShippingMethodTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethodTag], *http.Response, error) {
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

func (t *ShippingMethodTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "shipping-method-tag")
}

func (t *ShippingMethodTagRepository) Upsert(ctx ApiContext, entity []ShippingMethodTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "shipping_method_tag")
}

func (t *ShippingMethodTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "shipping_method_tag")
}

type ShippingMethodTag struct {

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

}
