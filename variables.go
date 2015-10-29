package sqlcomposer

import (
	"fmt"
	"time"
)

// Values

type SQLValue struct {
	value interface{}
}

func (self *SQLValue) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *SQLValue) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = context.GetNextParameterName()
	values = []interface{}{self.value}
	return
}

func SQLVariable(value interface{}) SQLExpression {
	if value == nil {
		return &SQLNullLiteral{}
	}

	switch val := value.(type) {
	case *string:
		value = *val
	case *bool:
		value = *val
	case *uint:
		value = *val
	case *uint8:
		value = *val
	case *uint16:
		value = *val
	case *uint32:
		value = *val
	case *uint64:
		value = *val
	case *int:
		value = *val
	case *int8:
		value = *val
	case *int16:
		value = *val
	case *int32:
		value = *val
	case *int64:
		value = *val
	case *float32:
		value = *val
	case *float64:
		value = *val
	case *time.Time:
		value = *val
	}

	switch val := value.(type) {
	case string, bool, uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, time.Time:
		return &SQLValue{value: val}
	default:
		panic(fmt.Errorf("%T cannot be a SQL value", value))
	}
}
