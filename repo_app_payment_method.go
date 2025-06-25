package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppPaymentMethodRepository struct {
	*GenericRepository[AppPaymentMethod]
}

func NewAppPaymentMethodRepository(client *Client) *AppPaymentMethodRepository {
	return &AppPaymentMethodRepository{
		GenericRepository: NewGenericRepository[AppPaymentMethod](client),
	}
}

func (t *AppPaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppPaymentMethod], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-payment-method")
}

func (t *AppPaymentMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppPaymentMethod], *http.Response, error) {
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

func (t *AppPaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-payment-method")
}

func (t *AppPaymentMethodRepository) Upsert(ctx ApiContext, entity []AppPaymentMethod) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_payment_method")
}

func (t *AppPaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_payment_method")
}

type AppPaymentMethod struct {

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	AppName      string  `json:"appName,omitempty"`

	CaptureUrl      string  `json:"captureUrl,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	FinalizeUrl      string  `json:"finalizeUrl,omitempty"`

	Id      string  `json:"id,omitempty"`

	Identifier      string  `json:"identifier,omitempty"`

	OriginalMedia      *Media  `json:"originalMedia,omitempty"`

	OriginalMediaId      string  `json:"originalMediaId,omitempty"`

	PayUrl      string  `json:"payUrl,omitempty"`

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	RecurringUrl      string  `json:"recurringUrl,omitempty"`

	RefundUrl      string  `json:"refundUrl,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ValidateUrl      string  `json:"validateUrl,omitempty"`

}
