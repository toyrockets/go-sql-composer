package sqlcomposer

import (
	"time"
    "fmt"
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
    case string, bool, uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, time.Time:
        return &SQLValue{value: val}
    default:
        panic(fmt.Errorf("%T cannot be a SQL value", value))
        return nil
    }
}
