package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PaymentMethodRepository struct {
	*GenericRepository[PaymentMethod]
}

func NewPaymentMethodRepository(client *Client) *PaymentMethodRepository {
	return &PaymentMethodRepository{
		GenericRepository: NewGenericRepository[PaymentMethod](client),
	}
}

func (t *PaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PaymentMethod], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "payment-method")
}

func (t *PaymentMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PaymentMethod], *http.Response, error) {
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

func (t *PaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "payment-method")
}

func (t *PaymentMethodRepository) Upsert(ctx ApiContext, entity []PaymentMethod) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "payment_method")
}

func (t *PaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "payment_method")
}

type PaymentMethod struct {

	Active      bool  `json:"active,omitempty"`

	AfterOrderEnabled      bool  `json:"afterOrderEnabled,omitempty"`

	AppPaymentMethod      *AppPaymentMethod  `json:"appPaymentMethod,omitempty"`

	Asynchronous      bool  `json:"asynchronous,omitempty"`

	AvailabilityRule      *Rule  `json:"availabilityRule,omitempty"`

	AvailabilityRuleId      string  `json:"availabilityRuleId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	Description      string  `json:"description,omitempty"`

	DistinguishableName      string  `json:"distinguishableName,omitempty"`

	FormattedHandlerIdentifier      string  `json:"formattedHandlerIdentifier,omitempty"`

	HandlerIdentifier      string  `json:"handlerIdentifier,omitempty"`

	Id      string  `json:"id,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Name      string  `json:"name,omitempty"`

	OrderTransactions      []OrderTransaction  `json:"orderTransactions,omitempty"`

	Plugin      *Plugin  `json:"plugin,omitempty"`

	PluginId      string  `json:"pluginId,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Prepared      bool  `json:"prepared,omitempty"`

	Recurring      bool  `json:"recurring,omitempty"`

	Refundable      bool  `json:"refundable,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	ShortName      string  `json:"shortName,omitempty"`

	Synchronous      bool  `json:"synchronous,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []PaymentMethodTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
