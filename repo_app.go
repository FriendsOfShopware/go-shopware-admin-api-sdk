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
	Name string `json:"name,omitempty"`

	Privacy string `json:"privacy,omitempty"`

	AclRole *AclRole `json:"aclRole,omitempty"`

	Translations []AppTranslation `json:"translations,omitempty"`

	AclRoleId string `json:"aclRoleId,omitempty"`

	Id string `json:"id,omitempty"`

	License string `json:"license,omitempty"`

	Modules interface{} `json:"modules,omitempty"`

	Icon string `json:"icon,omitempty"`

	MainModule interface{} `json:"mainModule,omitempty"`

	Description string `json:"description,omitempty"`

	Copyright string `json:"copyright,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	ActionButtons []AppActionButton `json:"actionButtons,omitempty"`

	Active bool `json:"active,omitempty"`

	Configurable bool `json:"configurable,omitempty"`

	Label string `json:"label,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Version string `json:"version,omitempty"`

	IntegrationId string `json:"integrationId,omitempty"`

	Templates []AppTemplate `json:"templates,omitempty"`

	Scripts []Script `json:"scripts,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	Path string `json:"path,omitempty"`

	PrivacyPolicyExtensions string `json:"privacyPolicyExtensions,omitempty"`

	CustomFieldSets []CustomFieldSet `json:"customFieldSets,omitempty"`

	Cookies interface{} `json:"cookies,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Integration *Integration `json:"integration,omitempty"`

	PaymentMethods []AppPaymentMethod `json:"paymentMethods,omitempty"`

	CmsBlocks []AppCmsBlock `json:"cmsBlocks,omitempty"`

	Author string `json:"author,omitempty"`

	IconRaw interface{} `json:"iconRaw,omitempty"`

	AppSecret string `json:"appSecret,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type AppCollection struct {
	EntityCollection

	Data []App `json:"data"`
}
