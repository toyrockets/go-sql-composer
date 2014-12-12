package sqlcomposer

import (
	"strings"
)

type InsertStatement struct {
    table Table
    values map[Column]SQLExpression
}

func (self *InsertStatement) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: NumericStyle})
    return
}

func (self *InsertStatement) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = "insert into"
	
    tableSQL, _ := self.table.GenerateSQLWithContext(context)
    SQL += " " + tableSQL

    columnFragments := []string{}
	valueFragments := []string{}
    values = []interface{}{}

    
    for column, value := range self.values {
        columnSQL, columnValues := column.GenerateSQLWithContext(context)
        columnFragments = append(columnFragments, columnSQL)
        values = append(values, columnValues...)
        
        valueSQL, stuff :=  value.GenerateSQLWithContext(context)
        valueFragments = append(valueFragments, valueSQL)
        values = append(values, stuff...)
        
    }

    SQL += " (" + strings.Join(columnFragments, ", ") + ")"
    SQL += " values(" + strings.Join(valueFragments, ", ") + ")"
	return
}

func (self *InsertStatement) Values(values map[interface{}]interface{}) *InsertStatement {
    if self.values == nil {
        self.values = map[Column]SQLExpression{}
    }

	if len(values) > 0 {	
		for key, value := range values {
            var column Column
            columnName, ok := key.(string)

            if ok {
                column = Column{Name: columnName}
            } else {
                column, ok = key.(Column)
                
                if !ok {
                    continue
                }
            }
			//
			// 	    fmt.Printf("%v is a map? %v\n", val, reflect.ValueOf(val).Kind() == reflect.Map)
            
            expression := SQLVariable(value)
            
            if expression != nil {
                self.values[column] = expression
            }
		}
	}
	
	return self
}

func Insert (table interface{}) *InsertStatement {
    var tableValue Table
    tableName, ok := table.(string)

    if ok {
        tableValue = Table{Name: tableName}
    } else {
        tableValue, ok = table.(Table)
        
        if !ok {
            return nil
        }
    }
    
    return &InsertStatement{table: tableValue}
}
