package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type TagRepository struct {
	*GenericRepository[Tag]
}

func NewTagRepository(client *Client) *TagRepository {
	return &TagRepository{
		GenericRepository: NewGenericRepository[Tag](client),
	}
}

func (t *TagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Tag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "tag")
}

func (t *TagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Tag], *http.Response, error) {
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

func (t *TagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "tag")
}

func (t *TagRepository) Upsert(ctx ApiContext, entity []Tag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "tag")
}

func (t *TagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "tag")
}

type Tag struct {

	Orders      []Order  `json:"orders,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	NewsletterRecipients      []NewsletterRecipient  `json:"newsletterRecipients,omitempty"`

	LandingPages      []LandingPage  `json:"landingPages,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	Rules      []Rule  `json:"rules,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Media      []Media  `json:"media,omitempty"`

}
