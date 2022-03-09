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
	CmsBlockVersionId string `json:"cmsBlockVersionId,omitempty"`

	Id string `json:"id,omitempty"`

	Block *CmsBlock `json:"block,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Translations []CmsSlotTranslation `json:"translations,omitempty"`

	BlockId string `json:"blockId,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Slot string `json:"slot,omitempty"`

	Data interface{} `json:"data,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Type string `json:"type,omitempty"`

	Locked bool `json:"locked,omitempty"`
}

type CmsSlotCollection struct {
	EntityCollection

	Data []CmsSlot `json:"data"`
}
