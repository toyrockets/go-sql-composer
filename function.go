package sqlcomposer

import (
	"fmt"
	"strings"
)

type SQLFunction struct {
	Name SQLStringLiteral
	Arguments []SQLExpression
}

func (self *SQLFunction) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLFunction) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	sqlFragments := []string{}
    values = []interface{}{}

	nameSQL, nameValues := self.Name.GenerateSQLWithContext(context)
	values = append(values, nameValues...)

	for _, expression := range self.Arguments {
		expressionSQL, expressionValues := expression.GenerateSQLWithContext(context)
		sqlFragments = append(sqlFragments, expressionSQL)
		values = append(values, expressionValues...)
	}

	SQL = fmt.Sprintf("%s(%s)", nameSQL, strings.Join(sqlFragments, ", "))

	return
}

func Func(name string, arguments ...interface{}) *SQLFunction {
	args := []SQLExpression{}
	
	for _, argument := range arguments {
		functionValue, ok := argument.(*SQLFunction)
		
		if ok {
			args = append(args, functionValue)
		} else {
			args = append(args, SQLVariable(argument))
		}
	}
	
	return &SQLFunction{Name: SQLStringLiteral{Value: name}, Arguments: args}
}
