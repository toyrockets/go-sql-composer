package sqlcomposer

import (
	"fmt"
)

type ColumnReference struct {
	expression SQLExpression
	alias      *SQLIdentifier
}

func (self *ColumnReference) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *ColumnReference) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	expressionSQL, values := self.expression.GenerateSQLWithContext(context)

	if self.alias != nil {
		aliasSQL, aliasValues := self.alias.GenerateSQLWithContext(context)
		values = append(values, aliasValues...)
		SQL = fmt.Sprintf("%s as %s", expressionSQL, aliasSQL)
	} else {
		SQL = expressionSQL
	}

	return
}

type ColumnList []ColumnReference

func (columns ColumnList) Len() int {
	return len(columns)
}
func (columns ColumnList) Swap(i, j int) {
	columns[i], columns[j] = columns[j], columns[i]
}
func (columns ColumnList) Less(i, j int) bool {
	column_i := columns[i]
	column_j := columns[j]
	sql_i, _ := column_i.expression.GenerateSQL()
	sql_j, _ := column_j.expression.GenerateSQL()
	return sql_i < sql_j
}
