package sqlcomposer

import (
    "strconv"
)

type BindVariableStyle int

const (
	NumericStyle BindVariableStyle = iota
	NamedStyle BindVariableStyle = iota
	QuestionMarkStyle BindVariableStyle = iota
	
	Postgres BindVariableStyle = NumericStyle
	Oracle BindVariableStyle = NamedStyle
	MySQL BindVariableStyle = QuestionMarkStyle
    DefaultStyle BindVariableStyle = NumericStyle
)

type SQLGenerationContext struct {
    Style BindVariableStyle
    parameterIndex int
    parameterName string
}

func (self *SQLGenerationContext) GetNextParameterName() string {
    switch (self.Style) {
        case NumericStyle:
            self.parameterIndex++
            return "$" + strconv.Itoa(self.parameterIndex)
        case NamedStyle:
            return ":" + self.parameterName
        case QuestionMarkStyle:
            fallthrough
        default:
            return "?"
    }
    
}

//func (self *SQLGenerationContext) GetBindVariableName

type SQLExpression interface {
	GenerateSQL() (SQL string, values []interface{})
	GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{})
}

// Identifier

type SQLIdentifier struct {
    Name string
}

func (self *SQLIdentifier) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLIdentifier) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = self.Name
    values = []interface{}{}
    return
}

type Column struct {
	Name string
	Alias string
}

func (self *Column) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *Column) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	if len(self.Alias) == 0 {
		SQL = self.Name
	} else {
		SQL = self.Name + " as " + self.Alias
	}

	values = []interface{}{}
	return
}

type TableExpression interface {
	GenerateSQL() (SQL string, values []interface{})
}

type TableReference struct {
	tableExpression TableExpression
	alias string
}

func (self *TableReference) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *TableReference) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	expressionSQL, _ := self.tableExpression.GenerateSQL()
	SQL = expressionSQL

	if len(self.alias) > 0 {
		SQL += " as " + self.alias
	}

	values = []interface{}{}
	return
}

type Table struct {
	Name string
}

func (self *Table) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *Table) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = self.Name
	values = []interface{}{}
	return
}
