package sqlcomposer

import (
    "reflect"
)

// Values

type SQLBooleanValue struct {
    value bool
}

func (self *SQLBooleanValue) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLBooleanValue) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}


type SQLStringValue struct {
    value string
}

func (self *SQLStringValue) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLStringValue) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}


type SQLUIntValue struct {
    value uint
}

func (self *SQLUIntValue) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLUIntValue) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLUInt8Value struct {
    value uint8
}

func (self *SQLUInt8Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLUInt8Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLUInt16Value struct {
    value uint16
}

func (self *SQLUInt16Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLUInt16Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLUInt32Value struct {
    value uint32
}

func (self *SQLUInt32Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLUInt32Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLUInt64Value struct {
    value uint64
}

func (self *SQLUInt64Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLUInt64Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLIntValue struct {
    value int
}

func (self *SQLIntValue) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLIntValue) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLInt8Value struct {
    value int8
}

func (self *SQLInt8Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLInt8Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLInt16Value struct {
    value int16
}

func (self *SQLInt16Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLInt16Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLInt32Value struct {
    value int32
}

func (self *SQLInt32Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLInt32Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLInt64Value struct {
    value int64
}

func (self *SQLInt64Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLInt64Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLFloat32Value struct {
    value float32
}

func (self *SQLFloat32Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLFloat32Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

type SQLFloat64Value struct {
    value float64
}

func (self *SQLFloat64Value) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SQLFloat64Value) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL = context.GetNextParameterName()
    values = []interface{}{self.value}
    return
}

func SQLVariable(value interface{}) SQLExpression {
    if reflect.TypeOf(value).Kind() == reflect.Ptr {
        if value == nil || reflect.ValueOf(value).IsNil() {
            return &SQLNullLiteral{}
        }
        
        stringValue, ok := value.(*string)
    
        if ok {
            return &SQLStringValue{value: *stringValue}
        }
    
        boolValue, ok := value.(*bool)
    
        if ok {
            return &SQLBooleanLiteral{Value: *boolValue}
        }
    
        uintValue, ok := value.(*int)
    
        if ok {
            return &SQLIntValue{value: *uintValue}
        }
    
        uint8Value, ok := value.(*uint8)
    
        if ok {
            return &SQLUInt8Value{value: *uint8Value}
        }
    
        uint16Value, ok := value.(*uint16)
    
        if ok {
            return &SQLUInt16Value{value: *uint16Value}
        }
    
        uint32Value, ok := value.(*uint32)
    
        if ok {
            return &SQLUInt32Value{value: *uint32Value}
        }
    
        uint64Value, ok := value.(*uint64)
    
        if ok {
            return &SQLUInt64Value{value: *uint64Value}
        }
    
        intValue, ok := value.(*int)
    
        if ok {
            return &SQLIntValue{value: *intValue}
        }
    
        int8Value, ok := value.(*int8)
    
        if ok {
            return &SQLInt8Value{value: *int8Value}
        }
    
        int16Value, ok := value.(*int16)
    
        if ok {
            return &SQLInt16Value{value: *int16Value}
        }
    
        int32Value, ok := value.(*int32)
    
        if ok {
            return &SQLInt32Value{value: *int32Value}
        }
    
        int64Value, ok := value.(*int64)
    
        if ok {
            return &SQLInt64Value{value: *int64Value}
        }
    
        float32Value, ok := value.(*float32)
    
        if ok {
            return &SQLFloat32Value{value: *float32Value}
        }
    
        float64Value, ok := value.(*float64)
    
        if ok {
            return &SQLFloat64Value{value: *float64Value}
        }
        
    } else {
    
        stringValue, ok := value.(string)
    
        if ok {
            return &SQLStringValue{value: stringValue}
        }
    
        boolValue, ok := value.(bool)
    
        if ok {
            return &SQLBooleanLiteral{Value: boolValue}
        }
    
        uintValue, ok := value.(int)
    
        if ok {
            return &SQLIntValue{value: uintValue}
        }
    
        uint8Value, ok := value.(uint8)
    
        if ok {
            return &SQLUInt8Value{value: uint8Value}
        }
    
        uint16Value, ok := value.(uint16)
    
        if ok {
            return &SQLUInt16Value{value: uint16Value}
        }
    
        uint32Value, ok := value.(uint32)
    
        if ok {
            return &SQLUInt32Value{value: uint32Value}
        }
    
        uint64Value, ok := value.(uint64)
    
        if ok {
            return &SQLUInt64Value{value: uint64Value}
        }
    
        intValue, ok := value.(int)
    
        if ok {
            return &SQLIntValue{value: intValue}
        }
    
        int8Value, ok := value.(int8)
    
        if ok {
            return &SQLInt8Value{value: int8Value}
        }
    
        int16Value, ok := value.(int16)
    
        if ok {
            return &SQLInt16Value{value: int16Value}
        }
    
        int32Value, ok := value.(int32)
    
        if ok {
            return &SQLInt32Value{value: int32Value}
        }
    
        int64Value, ok := value.(int64)
    
        if ok {
            return &SQLInt64Value{value: int64Value}
        }
    
        float32Value, ok := value.(float32)
    
        if ok {
            return &SQLFloat32Value{value: float32Value}
        }
    
        float64Value, ok := value.(float64)
    
        if ok {
            return &SQLFloat64Value{value: float64Value}
        }
    }
    
    return nil
}
