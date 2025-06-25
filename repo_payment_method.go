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

	PluginId      string  `json:"pluginId,omitempty"`

	AfterOrderEnabled      bool  `json:"afterOrderEnabled,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Plugin      *Plugin  `json:"plugin,omitempty"`

	HandlerIdentifier      string  `json:"handlerIdentifier,omitempty"`

	Asynchronous      bool  `json:"asynchronous,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	AppPaymentMethod      *AppPaymentMethod  `json:"appPaymentMethod,omitempty"`

	Id      string  `json:"id,omitempty"`

	Prepared      bool  `json:"prepared,omitempty"`

	Recurring      bool  `json:"recurring,omitempty"`

	AvailabilityRule      *Rule  `json:"availabilityRule,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Description      string  `json:"description,omitempty"`

	Synchronous      bool  `json:"synchronous,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	ShortName      string  `json:"shortName,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Active      bool  `json:"active,omitempty"`

	FormattedHandlerIdentifier      string  `json:"formattedHandlerIdentifier,omitempty"`

	Refundable      bool  `json:"refundable,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	DistinguishableName      string  `json:"distinguishableName,omitempty"`

	AvailabilityRuleId      string  `json:"availabilityRuleId,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	OrderTransactions      []OrderTransaction  `json:"orderTransactions,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Translations      []PaymentMethodTranslation  `json:"translations,omitempty"`

}
