package sqlcomposer

import (
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

	stringValue, ok := value.(string)

	if ok {
		return &SQLStringLiteral{Value: stringValue}
	}

	boolValue, ok := value.(bool)

	if ok {
		return &SQLBooleanLiteral{Value: boolValue}
	}

	uintValue, ok := value.(int)

	if ok {
		return &SQLIntLiteral{Value: uintValue}
	}

	uint8Value, ok := value.(uint8)

	if ok {
		return &SQLUInt8Literal{Value: uint8Value}
	}

	uint16Value, ok := value.(uint16)

	if ok {
		return &SQLUInt16Literal{Value: uint16Value}
	}

	uint32Value, ok := value.(uint32)

	if ok {
		return &SQLUInt32Literal{Value: uint32Value}
	}

	uint64Value, ok := value.(uint64)

	if ok {
		return &SQLUInt64Literal{Value: uint64Value}
	}

	intValue, ok := value.(int)

	if ok {
		return &SQLIntLiteral{Value: intValue}
	}

	int8Value, ok := value.(int8)

	if ok {
		return &SQLInt8Literal{Value: int8Value}
	}

	int16Value, ok := value.(int16)

	if ok {
		return &SQLInt16Literal{Value: int16Value}
	}

	int32Value, ok := value.(int32)

	if ok {
		return &SQLInt32Literal{Value: int32Value}
	}

	int64Value, ok := value.(int64)

	if ok {
		return &SQLInt64Literal{Value: int64Value}
	}

	float32Value, ok := value.(float32)

	if ok {
		return &SQLFloat32Literal{Value: float32Value}
	}

	float64Value, ok := value.(float64)

	if ok {
		return &SQLFloat64Literal{Value: float64Value}
	}

	return nil

}
