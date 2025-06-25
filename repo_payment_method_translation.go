package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PaymentMethodTranslationRepository struct {
	*GenericRepository[PaymentMethodTranslation]
}

func NewPaymentMethodTranslationRepository(client *Client) *PaymentMethodTranslationRepository {
	return &PaymentMethodTranslationRepository{
		GenericRepository: NewGenericRepository[PaymentMethodTranslation](client),
	}
}

func (t *PaymentMethodTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PaymentMethodTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "payment-method-translation")
}

func (t *PaymentMethodTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PaymentMethodTranslation], *http.Response, error) {
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

func (t *PaymentMethodTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "payment-method-translation")
}

func (t *PaymentMethodTranslationRepository) Upsert(ctx ApiContext, entity []PaymentMethodTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "payment_method_translation")
}

func (t *PaymentMethodTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "payment_method_translation")
}

type PaymentMethodTranslation struct {

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	PaymentMethod      *PaymentMethod  `json:"paymentMethod,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Description      string  `json:"description,omitempty"`

	PaymentMethodId      string  `json:"paymentMethodId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	DistinguishableName      string  `json:"distinguishableName,omitempty"`

}
