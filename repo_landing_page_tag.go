package go_shopware_admin_sdk

import (
	"net/http"

)

type LandingPageTagRepository struct {
	*GenericRepository[LandingPageTag]
}

func NewLandingPageTagRepository(client *Client) *LandingPageTagRepository {
	return &LandingPageTagRepository{
		GenericRepository: NewGenericRepository[LandingPageTag](client),
	}
}

func (t *LandingPageTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "landing-page-tag")
}

func (t *LandingPageTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageTag], *http.Response, error) {
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

func (t *LandingPageTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "landing-page-tag")
}

func (t *LandingPageTagRepository) Upsert(ctx ApiContext, entity []LandingPageTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "landing_page_tag")
}

func (t *LandingPageTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "landing_page_tag")
}

type LandingPageTag struct {

	LandingPageId      string  `json:"landingPageId,omitempty"`

	LandingPageVersionId      string  `json:"landingPageVersionId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	LandingPage      *LandingPage  `json:"landingPage,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

}
