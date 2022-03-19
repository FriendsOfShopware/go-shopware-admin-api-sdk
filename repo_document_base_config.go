package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DocumentBaseConfigRepository ClientService

func (t DocumentBaseConfigRepository) Search(ctx ApiContext, criteria Criteria) (*DocumentBaseConfigCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/document-base-config", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DocumentBaseConfigCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DocumentBaseConfigRepository) SearchAll(ctx ApiContext, criteria Criteria) (*DocumentBaseConfigCollection, *http.Response, error) {
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

func (t DocumentBaseConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/document-base-config", criteria)

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

func (t DocumentBaseConfigRepository) Upsert(ctx ApiContext, entity []DocumentBaseConfig) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_base_config": {
		Entity:  "document_base_config",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DocumentBaseConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_base_config": {
		Entity:  "document_base_config",
		Action:  "delete",
		Payload: payload,
	}})
}

type DocumentBaseConfig struct {
	Name string `json:"name,omitempty"`

	FilenamePrefix string `json:"filenamePrefix,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Global bool `json:"global,omitempty"`

	DocumentNumber string `json:"documentNumber,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalesChannels []DocumentBaseConfigSalesChannel `json:"salesChannels,omitempty"`

	Id string `json:"id,omitempty"`

	DocumentTypeId string `json:"documentTypeId,omitempty"`

	DocumentType *DocumentType `json:"documentType,omitempty"`

	LogoId string `json:"logoId,omitempty"`

	FilenameSuffix string `json:"filenameSuffix,omitempty"`

	Logo *Media `json:"logo,omitempty"`
}

type DocumentBaseConfigCollection struct {
	EntityCollection

	Data []DocumentBaseConfig `json:"data"`
}
