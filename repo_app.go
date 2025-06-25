package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppRepository struct {
	*GenericRepository[App]
}

func NewAppRepository(client *Client) *AppRepository {
	return &AppRepository{
		GenericRepository: NewGenericRepository[App](client),
	}
}

func (t *AppRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[App], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app")
}

func (t *AppRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[App], *http.Response, error) {
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

func (t *AppRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app")
}

func (t *AppRepository) Upsert(ctx ApiContext, entity []App) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app")
}

func (t *AppRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app")
}

type App struct {

	AclRole      *AclRole  `json:"aclRole,omitempty"`

	AclRoleId      string  `json:"aclRoleId,omitempty"`

	ActionButtons      []AppActionButton  `json:"actionButtons,omitempty"`

	Active      bool  `json:"active,omitempty"`

	AllowDisable      bool  `json:"allowDisable,omitempty"`

	AllowedHosts      interface{}  `json:"allowedHosts,omitempty"`

	AppSecret      string  `json:"appSecret,omitempty"`

	AppShippingMethods      []AppShippingMethod  `json:"appShippingMethods,omitempty"`

	Author      string  `json:"author,omitempty"`

	BaseAppUrl      string  `json:"baseAppUrl,omitempty"`

	CheckoutGatewayUrl      string  `json:"checkoutGatewayUrl,omitempty"`

	CmsBlocks      []AppCmsBlock  `json:"cmsBlocks,omitempty"`

	Configurable      bool  `json:"configurable,omitempty"`

	Cookies      interface{}  `json:"cookies,omitempty"`

	Copyright      string  `json:"copyright,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFieldSets      []CustomFieldSet  `json:"customFieldSets,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	FlowActions      []AppFlowAction  `json:"flowActions,omitempty"`

	FlowEvents      []AppFlowEvent  `json:"flowEvents,omitempty"`

	Icon      string  `json:"icon,omitempty"`

	IconRaw      interface{}  `json:"iconRaw,omitempty"`

	Id      string  `json:"id,omitempty"`

	InAppPurchasesGatewayUrl      string  `json:"inAppPurchasesGatewayUrl,omitempty"`

	Integration      *Integration  `json:"integration,omitempty"`

	IntegrationId      string  `json:"integrationId,omitempty"`

	Label      string  `json:"label,omitempty"`

	License      string  `json:"license,omitempty"`

	MainModule      interface{}  `json:"mainModule,omitempty"`

	Modules      interface{}  `json:"modules,omitempty"`

	Name      string  `json:"name,omitempty"`

	Path      string  `json:"path,omitempty"`

	PaymentMethods      []AppPaymentMethod  `json:"paymentMethods,omitempty"`

	Privacy      string  `json:"privacy,omitempty"`

	PrivacyPolicyExtensions      string  `json:"privacyPolicyExtensions,omitempty"`

	ScriptConditions      []AppScriptCondition  `json:"scriptConditions,omitempty"`

	Scripts      []Script  `json:"scripts,omitempty"`

	SelfManaged      bool  `json:"selfManaged,omitempty"`

	SourceConfig      interface{}  `json:"sourceConfig,omitempty"`

	SourceType      string  `json:"sourceType,omitempty"`

	TaxProviders      []TaxProvider  `json:"taxProviders,omitempty"`

	TemplateLoadPriority      float64  `json:"templateLoadPriority,omitempty"`

	Templates      []AppTemplate  `json:"templates,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []AppTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Version      string  `json:"version,omitempty"`

	Webhooks      []Webhook  `json:"webhooks,omitempty"`

}
