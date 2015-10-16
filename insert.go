package sqlcomposer

import (
	"fmt"
	"strings"
	"sort"
)

type InsertStatement struct {
    table SQLIdentifier
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

    
	var columnList ColumnList
	for column := range self.values {
	    columnList = append(columnList, column)
	}
	sort.Sort(columnList)
	for _, column := range columnList {
		value := self.values[column]

        columnSQL, columnValues := column.GenerateSQLWithContext(context)
        columnFragments = append(columnFragments, columnSQL)
        values = append(values, columnValues...)
        
        valueSQL, stuff :=  value.GenerateSQLWithContext(context)
        valueFragments = append(valueFragments, valueSQL)
        values = append(values, stuff...)
    }

    SQL += fmt.Sprintf(" (%s) values (%s)", strings.Join(columnFragments, ", "), strings.Join(valueFragments, ", "))
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
            
			valueExpression, ok := value.(SQLExpression)
			
			if (ok) {
				self.values[column] = valueExpression;
			} else {
				expression := SQLVariable(value)
            
				if expression != nil {
					self.values[column] = expression
				}
			}
		}
	}
	
	return self
}

func Insert (table interface{}) *InsertStatement {
    var tableValue SQLIdentifier
    tableName, ok := table.(string)

    if ok {
        tableValue = SQLIdentifier{Name: tableName}
    } else {
        tableValue, ok = table.(SQLIdentifier)
        
        if !ok {
            return nil
        }
    }
    
    return &InsertStatement{table: tableValue}
}
