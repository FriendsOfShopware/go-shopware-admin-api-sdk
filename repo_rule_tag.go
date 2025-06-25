package go_shopware_admin_sdk

import (
	"net/http"

)

type RuleTagRepository struct {
	*GenericRepository[RuleTag]
}

func NewRuleTagRepository(client *Client) *RuleTagRepository {
	return &RuleTagRepository{
		GenericRepository: NewGenericRepository[RuleTag](client),
	}
}

func (t *RuleTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[RuleTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "rule-tag")
}

func (t *RuleTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[RuleTag], *http.Response, error) {
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

func (t *RuleTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "rule-tag")
}

func (t *RuleTagRepository) Upsert(ctx ApiContext, entity []RuleTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "rule_tag")
}

func (t *RuleTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "rule_tag")
}

type RuleTag struct {

	Tag      *Tag  `json:"tag,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

}
