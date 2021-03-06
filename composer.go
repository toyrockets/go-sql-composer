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

// Alias

type SQLAlias struct {
	Name string
}

func (self *SQLAlias) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *SQLAlias) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	if context.QuoteIdentifiers {
		SQL = fmt.Sprintf("\"%s\"", self.Name)
	} else {
		SQL = self.Name
	}

	values = []interface{}{}
	return
}

type TableReference struct {
	tableExpression SQLExpression
	alias           *SQLAlias
}

func (self *TableReference) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *TableReference) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	expressionSQL, values := self.tableExpression.GenerateSQLWithContext(context)
	SQL = expressionSQL

	if self.alias != nil {
		aliasSQL, aliasValues := self.alias.GenerateSQLWithContext(context)
		values = append(values, aliasValues...)
		SQL = fmt.Sprintf("%s as %s", expressionSQL, aliasSQL)
	} else {
		SQL = expressionSQL
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
