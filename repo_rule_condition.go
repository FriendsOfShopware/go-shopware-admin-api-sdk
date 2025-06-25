package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type RuleConditionRepository struct {
	*GenericRepository[RuleCondition]
}

func NewRuleConditionRepository(client *Client) *RuleConditionRepository {
	return &RuleConditionRepository{
		GenericRepository: NewGenericRepository[RuleCondition](client),
	}
}

func (t *RuleConditionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[RuleCondition], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "rule-condition")
}

func (t *RuleConditionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[RuleCondition], *http.Response, error) {
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

func (t *RuleConditionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "rule-condition")
}

func (t *RuleConditionRepository) Upsert(ctx ApiContext, entity []RuleCondition) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "rule_condition")
}

func (t *RuleConditionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "rule_condition")
}

type RuleCondition struct {

	Id      string  `json:"id,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	ScriptId      string  `json:"scriptId,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	Value      interface{}  `json:"value,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

	Children      []RuleCondition  `json:"children,omitempty"`

	Type      string  `json:"type,omitempty"`

	AppScriptCondition      *AppScriptCondition  `json:"appScriptCondition,omitempty"`

	Parent      *RuleCondition  `json:"parent,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
