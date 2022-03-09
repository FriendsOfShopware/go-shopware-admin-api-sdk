package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CmsSlotTranslationRepository ClientService

func (t CmsSlotTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CmsSlotTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-slot-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsSlotTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsSlotTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-slot-translation", criteria)

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

func (t CmsSlotTranslationRepository) Upsert(ctx ApiContext, entity []CmsSlotTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_slot_translation": {
		Entity:  "cms_slot_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsSlotTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_slot_translation": {
		Entity:  "cms_slot_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsSlotTranslation struct {
	Config interface{} `json:"config,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CmsSlotId string `json:"cmsSlotId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	CmsSlotVersionId string `json:"cmsSlotVersionId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CmsSlot *CmsSlot `json:"cmsSlot,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type CmsSlotTranslationCollection struct {
	EntityCollection

	Data []CmsSlotTranslation `json:"data"`
}
