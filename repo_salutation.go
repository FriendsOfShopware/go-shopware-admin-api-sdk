package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalutationRepository struct {
	*GenericRepository[Salutation]
}

func NewSalutationRepository(client *Client) *SalutationRepository {
	return &SalutationRepository{
		GenericRepository: NewGenericRepository[Salutation](client),
	}
}

func (t *SalutationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Salutation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "salutation")
}

func (t *SalutationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Salutation], *http.Response, error) {
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

func (t *SalutationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "salutation")
}

func (t *SalutationRepository) Upsert(ctx ApiContext, entity []Salutation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "salutation")
}

func (t *SalutationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "salutation")
}

type Salutation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomerAddresses      []CustomerAddress  `json:"customerAddresses,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	DisplayName      string  `json:"displayName,omitempty"`

	Id      string  `json:"id,omitempty"`

	LetterName      string  `json:"letterName,omitempty"`

	NewsletterRecipients      []NewsletterRecipient  `json:"newsletterRecipients,omitempty"`

	OrderAddresses      []OrderAddress  `json:"orderAddresses,omitempty"`

	OrderCustomers      []OrderCustomer  `json:"orderCustomers,omitempty"`

	SalutationKey      string  `json:"salutationKey,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []SalutationTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
