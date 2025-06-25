package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type StateMachineStateTranslationRepository struct {
	*GenericRepository[StateMachineStateTranslation]
}

func NewStateMachineStateTranslationRepository(client *Client) *StateMachineStateTranslationRepository {
	return &StateMachineStateTranslationRepository{
		GenericRepository: NewGenericRepository[StateMachineStateTranslation](client),
	}
}

func (t *StateMachineStateTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineStateTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "state-machine-state-translation")
}

func (t *StateMachineStateTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachineStateTranslation], *http.Response, error) {
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

func (t *StateMachineStateTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "state-machine-state-translation")
}

func (t *StateMachineStateTranslationRepository) Upsert(ctx ApiContext, entity []StateMachineStateTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "state_machine_state_translation")
}

func (t *StateMachineStateTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "state_machine_state_translation")
}

type StateMachineStateTranslation struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	StateMachineStateId      string  `json:"stateMachineStateId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}
