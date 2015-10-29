package sqlcomposer

import (
	"fmt"
	"sort"
	"strings"
)

type InsertStatement struct {
	table                SQLIdentifier
	values               map[ColumnReference]SQLExpression
	returningExpressions []SQLExpression
}

func (self *InsertStatement) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *InsertStatement) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = "insert into"

	tableSQL, _ := self.table.GenerateSQLWithContext(context)
	SQL += " " + tableSQL

	columnFragments := []string{}
	valueFragments := []string{}
	values = []interface{}{}

	var columnList ColumnList
	for column := range self.values {
		columnList = append(columnList, column)
	}
	sort.Stable(columnList)
	for _, column := range columnList {
		value := self.values[column]

		columnSQL, columnValues := column.GenerateSQLWithContext(context)
		columnFragments = append(columnFragments, columnSQL)
		values = append(values, columnValues...)

		valueSQL, stuff := value.GenerateSQLWithContext(context)
		_, ok := value.(*SelectStatement)
		if ok {
			valueSQL = fmt.Sprintf("(%s)", valueSQL)
		}

		valueFragments = append(valueFragments, valueSQL)
		values = append(values, stuff...)
	}

	SQL += fmt.Sprintf(" (%s) values (%s)", strings.Join(columnFragments, ", "), strings.Join(valueFragments, ", "))

	if len(self.returningExpressions) > 0 {
		sqlFragments := []string{}

		for _, expression := range self.returningExpressions {
			expressionSQL, expressionValues := expression.GenerateSQLWithContext(context)
			sqlFragments = append(sqlFragments, expressionSQL)
			values = append(values, expressionValues...)
		}

		SQL += fmt.Sprintf(" returning %s", strings.Join(sqlFragments, ", "))
	}

	return
}

func (self *InsertStatement) Values(values map[interface{}]interface{}) *InsertStatement {
	if self.values == nil {
		self.values = map[ColumnReference]SQLExpression{}
	}

	if len(values) > 0 {
		for key, value := range values {
			var column ColumnReference
			columnName, ok := key.(string)

			if ok {
				column = ColumnReference{expression: &SQLIdentifier{Name: columnName}}
			} else {
				if column, ok = key.(ColumnReference); ok {
					continue
				}
			}

			if valueExpression, ok := value.(SQLExpression); ok {
				self.values[column] = valueExpression
			} else {
				expression := SQLVariable(value)

				if expression != nil {
					self.values[column] = expression
				}
			}
		}
	}

	return self
}

func (self *InsertStatement) Returning(values ...interface{}) *InsertStatement {
	if self.returningExpressions == nil {
		self.returningExpressions = []SQLExpression{}
	}

	for _, value := range values {
		expression := SQLLiteral(value)
		self.returningExpressions = append(self.returningExpressions, expression)
	}

	return self
}

func Insert(table interface{}) *InsertStatement {
	var tableValue SQLIdentifier
	tableName, ok := table.(string)

	if ok {
		tableValue = SQLIdentifier{Name: tableName}
	} else {
		tableValue, ok = table.(SQLIdentifier)

		if !ok {
			return nil
		}
	}

	return &InsertStatement{table: tableValue}
}
