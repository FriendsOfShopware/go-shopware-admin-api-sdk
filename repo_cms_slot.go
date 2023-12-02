package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CmsSlotRepository ClientService

func (t CmsSlotRepository) Search(ctx ApiContext, criteria Criteria) (*CmsSlotCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-slot", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsSlotCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsSlotRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CmsSlotCollection, *http.Response, error) {
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

func (t CmsSlotRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-slot", criteria)

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

func (t CmsSlotRepository) Upsert(ctx ApiContext, entity []CmsSlot) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_slot": {
		Entity:  "cms_slot",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsSlotRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_slot": {
		Entity:  "cms_slot",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsSlot struct {

	Id      string  `json:"id,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Type      string  `json:"type,omitempty"`

	Slot      string  `json:"slot,omitempty"`

	Block      *CmsBlock  `json:"block,omitempty"`

	Translations      []CmsSlotTranslation  `json:"translations,omitempty"`

	CmsBlockVersionId      string  `json:"cmsBlockVersionId,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Data      interface{}  `json:"data,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	BlockId      string  `json:"blockId,omitempty"`

	FieldConfig      interface{}  `json:"fieldConfig,omitempty"`

}

type CmsSlotCollection struct {
	EntityCollection

	Data []CmsSlot `json:"data"`
}
