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

	PrivacyPolicyExtensions      string  `json:"privacyPolicyExtensions,omitempty"`

	Webhooks      []Webhook  `json:"webhooks,omitempty"`

	AppShippingMethods      []AppShippingMethod  `json:"appShippingMethods,omitempty"`

	Description      string  `json:"description,omitempty"`

	Id      string  `json:"id,omitempty"`

	Configurable      bool  `json:"configurable,omitempty"`

	AppSecret      string  `json:"appSecret,omitempty"`

	TemplateLoadPriority      float64  `json:"templateLoadPriority,omitempty"`

	SourceType      string  `json:"sourceType,omitempty"`

	Scripts      []Script  `json:"scripts,omitempty"`

	TaxProviders      []TaxProvider  `json:"taxProviders,omitempty"`

	Path      string  `json:"path,omitempty"`

	Author      string  `json:"author,omitempty"`

	Version      string  `json:"version,omitempty"`

	IconRaw      interface{}  `json:"iconRaw,omitempty"`

	BaseAppUrl      string  `json:"baseAppUrl,omitempty"`

	Integration      *Integration  `json:"integration,omitempty"`

	CustomFieldSets      []CustomFieldSet  `json:"customFieldSets,omitempty"`

	AllowedHosts      interface{}  `json:"allowedHosts,omitempty"`

	InAppPurchasesGatewayUrl      string  `json:"inAppPurchasesGatewayUrl,omitempty"`

	IntegrationId      string  `json:"integrationId,omitempty"`

	ScriptConditions      []AppScriptCondition  `json:"scriptConditions,omitempty"`

	FlowActions      []AppFlowAction  `json:"flowActions,omitempty"`

	License      string  `json:"license,omitempty"`

	AllowDisable      bool  `json:"allowDisable,omitempty"`

	ActionButtons      []AppActionButton  `json:"actionButtons,omitempty"`

	Translations      []AppTranslation  `json:"translations,omitempty"`

	Privacy      string  `json:"privacy,omitempty"`

	SourceConfig      interface{}  `json:"sourceConfig,omitempty"`

	SelfManaged      bool  `json:"selfManaged,omitempty"`

	PaymentMethods      []AppPaymentMethod  `json:"paymentMethods,omitempty"`

	CmsBlocks      []AppCmsBlock  `json:"cmsBlocks,omitempty"`

	FlowEvents      []AppFlowEvent  `json:"flowEvents,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Icon      string  `json:"icon,omitempty"`

	MainModule      interface{}  `json:"mainModule,omitempty"`

	CheckoutGatewayUrl      string  `json:"checkoutGatewayUrl,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	AclRoleId      string  `json:"aclRoleId,omitempty"`

	Templates      []AppTemplate  `json:"templates,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Copyright      string  `json:"copyright,omitempty"`

	Label      string  `json:"label,omitempty"`

	AclRole      *AclRole  `json:"aclRole,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Modules      interface{}  `json:"modules,omitempty"`

	Cookies      interface{}  `json:"cookies,omitempty"`

}
