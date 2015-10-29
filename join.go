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
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *Join) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	if len(self.joinType) > 0 {
		SQL = fmt.Sprintf("%s join ", self.joinType)
	} else {
		SQL = "join "
	}

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
