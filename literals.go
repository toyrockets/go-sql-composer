package sqlcomposer

import (
	"fmt"
	"strconv"
)

// Literals

// type SQLLiteral interface {
//     GenerateLiteralSQL() (SQL string, values []interface{})
//     GenerateSQL() (SQL string, values []interface{})
//     GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{})
// }

type SQLStringLiteral struct {
	Value string
}

func (self *SQLStringLiteral) GeneralLiteralSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *SQLStringLiteral) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLStringLiteral) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = self.Value
	values = []interface{}{}
	return
}

type SQLBooleanLiteral struct {
	Value bool
}

func (self *SQLBooleanLiteral) GeneralLiteralSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *SQLBooleanLiteral) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLBooleanLiteral) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatBool(self.Value)
	values = []interface{}{}
	return
}

type SQLNullLiteral struct {
}

func (self *SQLNullLiteral) GeneralLiteralSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *SQLNullLiteral) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLNullLiteral) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = "null"
	values = []interface{}{}
	return
}

type SQLUIntLiteral struct {
	Value uint
}

func (self *SQLUIntLiteral) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLUIntLiteral) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatUint(uint64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLUInt8Literal struct {
	Value uint8
}

func (self *SQLUInt8Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLUInt8Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatUint(uint64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLUInt16Literal struct {
	Value uint16
}

func (self *SQLUInt16Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLUInt16Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatUint(uint64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLUInt32Literal struct {
	Value uint32
}

func (self *SQLUInt32Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLUInt32Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatUint(uint64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLUInt64Literal struct {
	Value uint64
}

func (self *SQLUInt64Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLUInt64Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatUint(self.Value, 10)
	values = []interface{}{}
	return
}

type SQLIntLiteral struct {
	Value int
}

func (self *SQLIntLiteral) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLIntLiteral) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatInt(int64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLInt8Literal struct {
	Value int8
}

func (self *SQLInt8Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLInt8Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatInt(int64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLInt16Literal struct {
	Value int16
}

func (self *SQLInt16Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLInt16Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatInt(int64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLInt32Literal struct {
	Value int32
}

func (self *SQLInt32Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLInt32Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatInt(int64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLInt64Literal struct {
	Value int64
}

func (self *SQLInt64Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLInt64Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatInt(int64(self.Value), 10)
	values = []interface{}{}
	return
}

type SQLFloat32Literal struct {
	Value float32
}

func (self *SQLFloat32Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLFloat32Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatFloat(float64(self.Value), 'G', 6, 32)
	values = []interface{}{}
	return
}

type SQLFloat64Literal struct {
	Value float64
}

func (self *SQLFloat64Literal) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLFloat64Literal) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = strconv.FormatFloat(self.Value, 'G', 6, 64)
	values = []interface{}{}
	return
}

func SQLLiteral(value interface{}) SQLExpression {
	if value == nil {
		return &SQLNullLiteral{}
	}

	switch val := value.(type) {
	case string:
		return &SQLStringLiteral{Value: val}
	case bool:
		return &SQLBooleanLiteral{Value: val}
	case uint:
		return &SQLUIntLiteral{Value: val}
	case uint8:
		return &SQLUInt8Literal{Value: val}
	case uint16:
		return &SQLUInt16Literal{Value: val}
	case uint32:
		return &SQLUInt32Literal{Value: val}
	case uint64:
		return &SQLUInt64Literal{Value: val}
	case int:
		return &SQLIntLiteral{Value: val}
	case int8:
		return &SQLInt8Literal{Value: val}
	case int16:
		return &SQLInt16Literal{Value: val}
	case int32:
		return &SQLInt32Literal{Value: val}
	case int64:
		return &SQLInt64Literal{Value: val}
	case float32:
		return &SQLFloat32Literal{Value: val}
	case float64:
		return &SQLFloat64Literal{Value: val}
	default:
		panic(fmt.Errorf("%T cannot be used as a SQL literal", val))
	}
}
