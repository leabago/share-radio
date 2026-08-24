package sql

import (
	"fmt"
	"strings"
)

type SortDirection string

const (
	Asc  SortDirection = "ASC"
	Desc SortDirection = "DESC"
)

type OrderCondition struct {
	parts []string
}

func NewOrder() *OrderCondition {
	return &OrderCondition{}
}

func (o *OrderCondition) OrderBy(field string, direction SortDirection, nulls bool) *OrderCondition {
	var nullStr string
	if nulls {
		nullStr = " NULLS FIRST"
	} else {
		nullStr = " NULLS LAST"
	}

	o.parts = append(o.parts, fmt.Sprintf("%s %s%s", field, direction, nullStr))

	return o
}

func (o *OrderCondition) GetLenParts() int {
	return len(o.parts)
}

func (o *OrderCondition) SQL() string {
	if len(o.parts) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(o.parts, ",")
}
