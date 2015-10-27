package sqlcomposer

import (
    "fmt"
)

type DeleteStatement struct {
    table Table
    predicates []Predicate
}

func (self *DeleteStatement) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: NumericStyle})
    return
}

func (self *DeleteStatement) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL = "delete from"
	
    tableSQL, _ := self.table.GenerateSQLWithContext(context)
    SQL += fmt.Sprintf(" %s", tableSQL)

    values = []interface{}{}
    
	if len(self.predicates) > 0 {
        andPredicate := AndPredicate{predicates: self.predicates}
        predicateSQL, predicateValues := andPredicate.GenerateSQLWithContext(context)
        values = append(values, predicateValues...)

		SQL += " where " + predicateSQL
	}
    
	return
}

func (self *DeleteStatement) Where(predicates ...interface{}) *DeleteStatement {
    if self.predicates == nil {
        self.predicates = []Predicate{}
    }
    
    subpredicates := ParsePredicates(predicates...)
    self.predicates = append(self.predicates, subpredicates...)
	
	return self
}

func Delete(table interface{}) *DeleteStatement {
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
    
    return &DeleteStatement{table: tableValue}
}
