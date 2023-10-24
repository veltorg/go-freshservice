package freshservice

import (
	"fmt"
	"strings"
)

// RequesterGroups holds a list of Freshservice agents
type RequesterGroups struct {
	List []RequesterGroupDetails `json:"requester_groups"`
}

type RequesterGroupDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// Validate will confirm that an agent role is valid
func (rg *RequesterGroupDetails) Validate() error {
	validTypes := []string{
		"manual",
		"rule_based",
	}

	if !StringInSlice(rg.Type, validTypes) {
		return fmt.Errorf("Requester group type is invalid; choose from %s", strings.Join(validTypes, ","))
	}

	return nil
}

// RequesterGroupListFilter holds the filters available when listing Freservice agents
type RequesterGroupListFilter struct {
	PageQuery string
}

// QueryString allows the available filter items to meet the QueryFilter interface
func (af *RequesterGroupListFilter) QueryString() string {

	if af.PageQuery != "" {
		return af.PageQuery
	}

	return ""
}
