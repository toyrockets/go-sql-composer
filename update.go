package sqlcomposer

import (
	"strings"
)

type UpdateStatement struct {
    table Table
    values map[Column]SQLExpression
    predicates []Predicate
}

func (self *UpdateStatement) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: NumericStyle})
    return
}

func (self *UpdateStatement) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = "update"
	
    tableSQL, _ := self.table.GenerateSQLWithContext(context)
    SQL += " " + tableSQL

    sqlFragments := []string{}
    values = []interface{}{}

    
    for column, value := range self.values {
        columnSQL, columnValues := column.GenerateSQLWithContext(context)
        values = append(values, columnValues...)
        valueSQL, stuff :=  value.GenerateSQLWithContext(context)
        values = append(values, stuff...)

        sqlFragment := columnSQL + " = " + valueSQL
        sqlFragments = append(sqlFragments, sqlFragment)
    }

    SQL += " " + strings.Join(sqlFragments, ", ")
    
	if len(self.predicates) > 0 {
        andPredicate := AndPredicate{predicates: self.predicates}
        predicateSQL, predicateValues := andPredicate.GenerateSQLWithContext(context)
        values = append(values, predicateValues...)

		SQL += " where " + predicateSQL
	}
	
	return
}

func (self *UpdateStatement) Set(values map[interface{}]interface{}) *UpdateStatement {
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

func (self *UpdateStatement) Where(predicates ...interface{}) *UpdateStatement {
    if self.predicates == nil {
        self.predicates = []Predicate{}
    }
    
    subpredicates := ParsePredicates(predicates...)
    self.predicates = append(self.predicates, subpredicates...)
	
	return self
}

func Update(table interface{}) *UpdateStatement {
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
    
    return &UpdateStatement{table: tableValue}
}
