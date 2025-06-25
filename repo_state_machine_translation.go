package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type StateMachineTranslationRepository struct {
	*GenericRepository[StateMachineTranslation]
}

func NewStateMachineTranslationRepository(client *Client) *StateMachineTranslationRepository {
	return &StateMachineTranslationRepository{
		GenericRepository: NewGenericRepository[StateMachineTranslation](client),
	}
}

func (t *StateMachineTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "state-machine-translation")
}

func (t *StateMachineTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineTranslation], *http.Response, error) {
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

func (t *StateMachineTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "state-machine-translation")
}

func (t *StateMachineTranslationRepository) Upsert(ctx ApiContext, entity []StateMachineTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "state_machine_translation")
}

func (t *StateMachineTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "state_machine_translation")
}

type StateMachineTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	StateMachine      *StateMachine  `json:"stateMachine,omitempty"`

	StateMachineId      string  `json:"stateMachineId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
