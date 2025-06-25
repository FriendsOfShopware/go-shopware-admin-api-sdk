package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type LogEntryRepository struct {
	*GenericRepository[LogEntry]
}

func NewLogEntryRepository(client *Client) *LogEntryRepository {
	return &LogEntryRepository{
		GenericRepository: NewGenericRepository[LogEntry](client),
	}
}

func (t *LogEntryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[LogEntry], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "log-entry")
}

func (t *LogEntryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[LogEntry], *http.Response, error) {
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

func (t *LogEntryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "log-entry")
}

func (t *LogEntryRepository) Upsert(ctx ApiContext, entity []LogEntry) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "log_entry")
}

func (t *LogEntryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "log_entry")
}

type LogEntry struct {

	Context      interface{}  `json:"context,omitempty"`

	Extra      interface{}  `json:"extra,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Message      string  `json:"message,omitempty"`

	Level      float64  `json:"level,omitempty"`

	Channel      string  `json:"channel,omitempty"`

}
