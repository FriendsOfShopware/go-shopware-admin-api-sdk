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
	Active bool `json:"active,omitempty"`

	IconRaw interface{} `json:"iconRaw,omitempty"`

	Modules interface{} `json:"modules,omitempty"`

	AllowedHosts interface{} `json:"allowedHosts,omitempty"`

	Translations []AppTranslation `json:"translations,omitempty"`

	AppShippingMethods []AppShippingMethod `json:"appShippingMethods,omitempty"`

	Id string `json:"id,omitempty"`

	Author string `json:"author,omitempty"`

	Privacy string `json:"privacy,omitempty"`

	Templates []AppTemplate `json:"templates,omitempty"`

	PaymentMethods []AppPaymentMethod `json:"paymentMethods,omitempty"`

	AclRole *AclRole `json:"aclRole,omitempty"`

	Version string `json:"version,omitempty"`

	Icon string `json:"icon,omitempty"`

	MainModule interface{} `json:"mainModule,omitempty"`

	Cookies interface{} `json:"cookies,omitempty"`

	IntegrationId string `json:"integrationId,omitempty"`

	Integration *Integration `json:"integration,omitempty"`

	AclRoleId string `json:"aclRoleId,omitempty"`

	Name string `json:"name,omitempty"`

	PrivacyPolicyExtensions string `json:"privacyPolicyExtensions,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Path string `json:"path,omitempty"`

	License string `json:"license,omitempty"`

	AppSecret string `json:"appSecret,omitempty"`

	BaseAppUrl string `json:"baseAppUrl,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	ActionButtons []AppActionButton `json:"actionButtons,omitempty"`

	Scripts []Script `json:"scripts,omitempty"`

	TaxProviders []TaxProvider `json:"taxProviders,omitempty"`

	ScriptConditions []AppScriptCondition `json:"scriptConditions,omitempty"`

	CmsBlocks []AppCmsBlock `json:"cmsBlocks,omitempty"`

	Copyright string `json:"copyright,omitempty"`

	Configurable bool `json:"configurable,omitempty"`

	TemplateLoadPriority float64 `json:"templateLoadPriority,omitempty"`

	Label string `json:"label,omitempty"`

	Description string `json:"description,omitempty"`

	Webhooks []Webhook `json:"webhooks,omitempty"`

	FlowEvents []AppFlowEvent `json:"flowEvents,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AllowDisable bool `json:"allowDisable,omitempty"`

	CustomFieldSets []CustomFieldSet `json:"customFieldSets,omitempty"`

	FlowActions []AppFlowAction `json:"flowActions,omitempty"`
}

type AppCollection struct {
	EntityCollection

	Data []App `json:"data"`
}
