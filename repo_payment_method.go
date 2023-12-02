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

func (t PaymentMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PaymentMethodCollection, *http.Response, error) {
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

	ShortName      string  `json:"shortName,omitempty"`

	DistinguishableName      string  `json:"distinguishableName,omitempty"`

	Description      string  `json:"description,omitempty"`

	Position      float64  `json:"position,omitempty"`

	AvailabilityRuleId      string  `json:"availabilityRuleId,omitempty"`

	Recurring      bool  `json:"recurring,omitempty"`

	Translations      []PaymentMethodTranslation  `json:"translations,omitempty"`

	Name      string  `json:"name,omitempty"`

	Asynchronous      bool  `json:"asynchronous,omitempty"`

	Plugin      *Plugin  `json:"plugin,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	FormattedHandlerIdentifier      string  `json:"formattedHandlerIdentifier,omitempty"`

	Refundable      bool  `json:"refundable,omitempty"`

	OrderTransactions      []OrderTransaction  `json:"orderTransactions,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	PluginId      string  `json:"pluginId,omitempty"`

	AfterOrderEnabled      bool  `json:"afterOrderEnabled,omitempty"`

	Synchronous      bool  `json:"synchronous,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	AvailabilityRule      *Rule  `json:"availabilityRule,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	AppPaymentMethod      *AppPaymentMethod  `json:"appPaymentMethod,omitempty"`

	Active      bool  `json:"active,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	HandlerIdentifier      string  `json:"handlerIdentifier,omitempty"`

	Prepared      bool  `json:"prepared,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

type PaymentMethodCollection struct {
	EntityCollection

	Data []PaymentMethod `json:"data"`
}
