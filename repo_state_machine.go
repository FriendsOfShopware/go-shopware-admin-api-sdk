package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type StateMachineRepository struct {
	*GenericRepository[StateMachine]
}

func NewStateMachineRepository(client *Client) *StateMachineRepository {
	return &StateMachineRepository{
		GenericRepository: NewGenericRepository[StateMachine](client),
	}
}

func (t *StateMachineRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachine], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "state-machine")
}

func (t *StateMachineRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[StateMachine], *http.Response, error) {
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

func (t *StateMachineRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "state-machine")
}

func (t *StateMachineRepository) Upsert(ctx ApiContext, entity []StateMachine) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "state_machine")
}

func (t *StateMachineRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "state_machine")
}

type StateMachine struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	HistoryEntries      []StateMachineHistory  `json:"historyEntries,omitempty"`

	Id      string  `json:"id,omitempty"`

	InitialStateId      string  `json:"initialStateId,omitempty"`

	Name      string  `json:"name,omitempty"`

	States      []StateMachineState  `json:"states,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	Transitions      []StateMachineTransition  `json:"transitions,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []StateMachineTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}
