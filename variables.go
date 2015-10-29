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
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
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
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *bool:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *uint:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *uint8:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *uint16:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *uint32:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *uint64:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *int:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *int8:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *int16:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *int32:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *int64:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *float32:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *float64:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	case *time.Time:
		if val != nil {
			value = *val
		} else {
			return &SQLNullLiteral{}
		}
	}

	switch val := value.(type) {
	case string, bool, uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, time.Time:
		return &SQLValue{value: val}
	default:
		panic(fmt.Errorf("%T cannot be a SQL value", value))
	}
}
