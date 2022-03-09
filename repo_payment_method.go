package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PaymentMethodRepository ClientService

func (t PaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*PaymentMethodCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/payment-method", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PaymentMethodCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/payment-method", criteria)

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

func (t PaymentMethodRepository) Upsert(ctx ApiContext, entity []PaymentMethod) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"payment_method": {
		Entity:  "payment_method",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"payment_method": {
		Entity:  "payment_method",
		Action:  "delete",
		Payload: payload,
	}})
}

type PaymentMethod struct {
	AfterOrderEnabled bool `json:"afterOrderEnabled,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	FormattedHandlerIdentifier string `json:"formattedHandlerIdentifier,omitempty"`

	Asynchronous bool `json:"asynchronous,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	Id string `json:"id,omitempty"`

	Position float64 `json:"position,omitempty"`

	Prepared bool `json:"prepared,omitempty"`

	Synchronous bool `json:"synchronous,omitempty"`

	Media *Media `json:"media,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	HandlerIdentifier string `json:"handlerIdentifier,omitempty"`

	AppPaymentMethod *AppPaymentMethod `json:"appPaymentMethod,omitempty"`

	Description string `json:"description,omitempty"`

	AvailabilityRuleId string `json:"availabilityRuleId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	OrderTransactions []OrderTransaction `json:"orderTransactions,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	PluginId string `json:"pluginId,omitempty"`

	DistinguishableName string `json:"distinguishableName,omitempty"`

	Translations []PaymentMethodTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Active bool `json:"active,omitempty"`

	AvailabilityRule *Rule `json:"availabilityRule,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	Plugin *Plugin `json:"plugin,omitempty"`
}

type PaymentMethodCollection struct {
	EntityCollection

	Data []PaymentMethod `json:"data"`
}
