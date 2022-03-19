package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppRepository ClientService

func (t AppRepository) Search(ctx ApiContext, criteria Criteria) (*AppCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppCollection, *http.Response, error) {
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

func (t AppRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app", criteria)

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

func (t AppRepository) Upsert(ctx ApiContext, entity []App) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app": {
		Entity:  "app",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app": {
		Entity:  "app",
		Action:  "delete",
		Payload: payload,
	}})
}

type App struct {
	Copyright string `json:"copyright,omitempty"`

	IconRaw interface{} `json:"iconRaw,omitempty"`

	Integration *Integration `json:"integration,omitempty"`

	Templates []AppTemplate `json:"templates,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	Scripts []Script `json:"scripts,omitempty"`

	Active bool `json:"active,omitempty"`

	Configurable bool `json:"configurable,omitempty"`

	Version string `json:"version,omitempty"`

	Translations []AppTranslation `json:"translations,omitempty"`

	Description string `json:"description,omitempty"`

	AclRole *AclRole `json:"aclRole,omitempty"`

	CmsBlocks []AppCmsBlock `json:"cmsBlocks,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	Icon string `json:"icon,omitempty"`

	MainModule interface{} `json:"mainModule,omitempty"`

	Cookies interface{} `json:"cookies,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	IntegrationId string `json:"integrationId,omitempty"`

	PaymentMethods []AppPaymentMethod `json:"paymentMethods,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Privacy string `json:"privacy,omitempty"`

	CustomFieldSets []CustomFieldSet `json:"customFieldSets,omitempty"`

	Path string `json:"path,omitempty"`

	AppSecret string `json:"appSecret,omitempty"`

	Label string `json:"label,omitempty"`

	ActionButtons []AppActionButton `json:"actionButtons,omitempty"`

	License string `json:"license,omitempty"`

	Modules interface{} `json:"modules,omitempty"`

	PrivacyPolicyExtensions string `json:"privacyPolicyExtensions,omitempty"`

	AclRoleId string `json:"aclRoleId,omitempty"`

	Author string `json:"author,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type AppCollection struct {
	EntityCollection

	Data []App `json:"data"`
}
