package sqlcomposer

import (
	"fmt"
)

type JoinType string

const (
	InnerJoin      JoinType = ""
	LeftOuterJoin  JoinType = "left outer"
	RightOuterJoin JoinType = "right outer"
	FullOuterJoin  JoinType = "full outer"
	CrossJoin      JoinType = "cross"
)

type Join struct {
	joinType       JoinType
	tableReference *TableReference
	condition      Predicate
}

func (self *Join) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *Join) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = fmt.Sprintf(" %s join ", self.joinType)

	tableSQL, tableValues := self.tableReference.GenerateSQLWithContext(context)
	SQL += tableSQL
	values = append(values, tableValues...)

	if self.condition != nil {
		predicateSQL, predicateValues := self.condition.GenerateSQLWithContext(context)
		SQL += fmt.Sprintf(" on %s", predicateSQL)
		values = append(values, predicateValues...)
	}

	return
}
