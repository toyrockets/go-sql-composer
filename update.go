package sqlcomposer

import (
	"fmt"
	"sort"
	"strings"
)

type UpdateStatement struct {
	table                Table
	values               map[ColumnReference]SQLExpression
	predicates           []Predicate
	returningExpressions []SQLExpression
}

func (self *UpdateStatement) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *UpdateStatement) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = "update"

	tableSQL, _ := self.table.GenerateSQLWithContext(context)
	SQL += " " + tableSQL

	sqlFragments := []string{}
	values = []interface{}{}

	var columnList ColumnList
	for column := range self.values {
		columnList = append(columnList, column)
	}
	sort.Stable(columnList)
	for _, column := range columnList {
		value := self.values[column]

		columnSQL, columnValues := column.GenerateSQLWithContext(context)
		values = append(values, columnValues...)
		valueSQL, stuff := value.GenerateSQLWithContext(context)
		values = append(values, stuff...)

		sqlFragment := columnSQL + " = " + valueSQL
		sqlFragments = append(sqlFragments, sqlFragment)
	}

	SQL += fmt.Sprintf(" set %s", strings.Join(sqlFragments, ", "))

	if len(self.predicates) > 0 {
		andPredicate := AndPredicate{predicates: self.predicates}
		predicateSQL, predicateValues := andPredicate.GenerateSQLWithContext(context)
		values = append(values, predicateValues...)

		SQL += " where " + predicateSQL
	}

	if len(self.returningExpressions) > 0 {
		sqlFragments = []string{}

		for _, expression := range self.returningExpressions {
			expressionSQL, expressionValues := expression.GenerateSQLWithContext(context)
			sqlFragments = append(sqlFragments, expressionSQL)
			values = append(values, expressionValues...)
		}

		SQL += fmt.Sprintf(" returning %s", strings.Join(sqlFragments, ", "))
	}

	return
}

func (self *UpdateStatement) Set(values map[string]interface{}) *UpdateStatement {
	if self.values == nil {
		self.values = map[ColumnReference]SQLExpression{}
	}

	if len(values) > 0 {
		for key, value := range values {
			column := ColumnReference{expression: &SQLIdentifier{Name: key}}
			expression := SQLVariable(value)

			if expression != nil {
				self.values[column] = expression
			}
		}
	}

	return self
}

func (self *UpdateStatement) Where(predicates ...interface{}) *UpdateStatement {
	if self.predicates == nil {
		self.predicates = []Predicate{}
	}

	subpredicates := ParsePredicates(predicates...)
	self.predicates = append(self.predicates, subpredicates...)

	return self
}

func (self *UpdateStatement) Returning(values ...interface{}) *UpdateStatement {
	if self.returningExpressions == nil {
		self.returningExpressions = []SQLExpression{}
	}

	for _, value := range values {
		expression := SQLLiteral(value)
		self.returningExpressions = append(self.returningExpressions, expression)
	}

	return self
}

func Update(table interface{}) *UpdateStatement {
	var tableValue Table
	tableName, ok := table.(string)

	if ok {
		tableValue = Table{Name: tableName}
	} else {
		tableValue, ok = table.(Table)

		if !ok {
			return nil
		}
	}

	return &UpdateStatement{table: tableValue}
}
