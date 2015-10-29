package sqlcomposer

import (
	"fmt"
	"strconv"
	"strings"
)

var DefaultSQLGenerationContext *SQLGenerationContext

func init() {
	DefaultSQLGenerationContext = &SQLGenerationContext{Style: DefaultStyle, QuoteIdentifiers: true}
}

type BindVariableStyle int

const (
	NumericStyle      BindVariableStyle = iota
	NamedStyle        BindVariableStyle = iota
	QuestionMarkStyle BindVariableStyle = iota

	Postgres     BindVariableStyle = NumericStyle
	Oracle       BindVariableStyle = NamedStyle
	MySQL        BindVariableStyle = QuestionMarkStyle
	DefaultStyle BindVariableStyle = NumericStyle
)

type SQLGenerationContext struct {
	Style            BindVariableStyle
	QuoteIdentifiers bool
	parameterIndex   int
	parameterName    string
}

func (self *SQLGenerationContext) reset() {
	self.parameterIndex = 0
}

func (self *SQLGenerationContext) GetNextParameterName() string {
	switch self.Style {
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
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *SQLIdentifier) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	if context.QuoteIdentifiers {
		components := strings.Split(self.Name, ".")
		SQL = fmt.Sprintf("\"%s\"", strings.Join(components, "\".\""))
	} else {
		SQL = self.Name
	}

	values = []interface{}{}
	return
}

type TableExpression interface {
	GenerateSQL() (SQL string, values []interface{})
	GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{})
	GenerateTableExpressionSQLWithContext(contextr *SQLGenerationContext) (SQL string, values []interface{})
}

type TableReference struct {
	tableExpression TableExpression
	alias           *string
}

func (self *TableReference) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *TableReference) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	expressionSQL, _ := self.tableExpression.GenerateSQLWithContext(context)
	SQL = expressionSQL

	if self.alias != nil {
		SQL += fmt.Sprintf(" as %s", *self.alias)
	}

	values = []interface{}{}
	return
}

type Table struct {
	Name string
}

func (self *Table) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *Table) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	if context.QuoteIdentifiers {
		SQL = fmt.Sprintf("\"%s\"", self.Name)
	} else {
		SQL = self.Name
	}

	values = []interface{}{}
	return
}

func (self *Table) GenerateTableExpressionSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(context)
	return
}
