package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type RuleTagRepository ClientService

func (t RuleTagRepository) Search(ctx ApiContext, criteria Criteria) (*RuleTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/rule-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(RuleTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t RuleTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*RuleTagCollection, *http.Response, error) {
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

func (t RuleTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/rule-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t RuleTagRepository) Upsert(ctx ApiContext, entity []RuleTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"rule_tag": {
		Entity:  "rule_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t RuleTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"rule_tag": {
		Entity:  "rule_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type RuleTag struct {

	RuleId      string  `json:"ruleId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

}

type RuleTagCollection struct {
	EntityCollection

	Data []RuleTag `json:"data"`
}
