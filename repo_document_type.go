package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DocumentTypeRepository ClientService

func (t DocumentTypeRepository) Search(ctx ApiContext, criteria Criteria) (*DocumentTypeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/document-type", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DocumentTypeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DocumentTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*DocumentTypeCollection, *http.Response, error) {
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

func (t DocumentTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/document-type", criteria)

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

func (t DocumentTypeRepository) Upsert(ctx ApiContext, entity []DocumentType) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_type": {
		Entity:  "document_type",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DocumentTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_type": {
		Entity:  "document_type",
		Action:  "delete",
		Payload: payload,
	}})
}

type DocumentType struct {
	DocumentBaseConfigSalesChannels []DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`

	Name string `json:"name,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Translations []DocumentTypeTranslation `json:"translations,omitempty"`

	Documents []Document `json:"documents,omitempty"`

	Id string `json:"id,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DocumentBaseConfigs []DocumentBaseConfig `json:"documentBaseConfigs,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type DocumentTypeCollection struct {
	EntityCollection

	Data []DocumentType `json:"data"`
}
