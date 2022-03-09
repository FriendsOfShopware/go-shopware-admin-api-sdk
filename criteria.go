package go_shopware_admin_sdk

const (
	TotalCountModeDefault  = 0
	TotalCountModeExact    = 1
	TotalCountModeNextPage = 2

	SearchFilterTypeEquals    = "equals"
	SearchFilterTypeEqualsAny = "equalsAny"

	SearchSortDirectionAscending  = "ASC"
	SearchSortDirectionDescending = "DESC"
)

type Criteria struct {
	Includes       map[string][]string `json:"includes,omitempty"`
	Page           int64               `json:"page,omitempty"`
	Limit          int64               `json:"limit,omitempty"`
	IDs            []string            `json:"ids,omitempty"`
	Filter         []CriteriaFilter    `json:"filter,omitempty"`
	PostFilter     []CriteriaFilter    `json:"postFilter,omitempty"`
	Sort           []CriteriaSort      `json:"sort,omitempty"`
	Associations   map[string]Criteria `json:"associations,omitempty"`
	Term           string              `json:"term,omitempty"`
	TotalCountMode int                 `json:"totalCountMode,omitempty"`
}

type CriteriaFilter struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type CriteriaSort struct {
	Direction      string `json:"order"`
	Field          string `json:"field"`
	NaturalSorting bool   `json:"naturalSorting"`
}
