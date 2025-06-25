package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomerGroupTranslationRepository struct {
	*GenericRepository[CustomerGroupTranslation]
}

func NewCustomerGroupTranslationRepository(client *Client) *CustomerGroupTranslationRepository {
	return &CustomerGroupTranslationRepository{
		GenericRepository: NewGenericRepository[CustomerGroupTranslation](client),
	}
}

func (t *CustomerGroupTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerGroupTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-group-translation")
}

func (t *CustomerGroupTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerGroupTranslation], *http.Response, error) {
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

func (t *CustomerGroupTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-group-translation")
}

func (t *CustomerGroupTranslationRepository) Upsert(ctx ApiContext, entity []CustomerGroupTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_group_translation")
}

func (t *CustomerGroupTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_group_translation")
}

type CustomerGroupTranslation struct {

	Name      string  `json:"name,omitempty"`

	RegistrationTitle      string  `json:"registrationTitle,omitempty"`

	RegistrationOnlyCompanyRegistration      bool  `json:"registrationOnlyCompanyRegistration,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	RegistrationIntroduction      string  `json:"registrationIntroduction,omitempty"`

	RegistrationSeoMetaDescription      string  `json:"registrationSeoMetaDescription,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CustomerGroupId      string  `json:"customerGroupId,omitempty"`

	CustomerGroup      *CustomerGroup  `json:"customerGroup,omitempty"`

}
