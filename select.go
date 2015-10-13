package sqlcomposer

import (
	"fmt"
	"strings"
	"reflect"
)

type SelectStatement struct {
	selectList []SQLExpression
	tableReferences []TableReference
    predicates []Predicate
}

func (self *SelectStatement) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SelectStatement) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	var sqlFragments []string
	SQL = "select"
    values = []interface{}{}
	
	sqlFragments = []string{}

	for _, expression := range self.selectList {
		expressionSQL, selectListValues := expression.GenerateSQLWithContext(context)
		sqlFragments = append(sqlFragments, expressionSQL)
		values = append(values, selectListValues...)
	}

	SQL += " " + strings.Join(sqlFragments, ", ")
	
	if len(self.tableReferences) > 0 {
		sqlFragments = []string{}

		for _, reference := range self.tableReferences {
			expressionSQL, referenceValues := reference.GenerateSQLWithContext(context)
			sqlFragments = append(sqlFragments, expressionSQL)
			values = append(values, referenceValues...)
		}

		SQL += " from " + strings.Join(sqlFragments, ", ")
	}

	if len(self.predicates) > 0 {
        andPredicate := AndPredicate{predicates: self.predicates}
        predicateSQL, predicateValues := andPredicate.GenerateSQLWithContext(context)

		SQL += " where " + predicateSQL
		values = append(values, predicateValues...)
	}
	
	return
}

func (self *SelectStatement) From(tables ...interface{}) *SelectStatement {
	if len(tables) > 0 {	
		for _, val := range tables {
			// _, ok := val.(map[string]string)
			//
			// if ok {
			// 	fmt.Println("Found a map")
			// }
			//
			// 	    fmt.Printf("%v is a map? %v\n", val, reflect.ValueOf(val).Kind() == reflect.Map)
		
		    switch reflect.ValueOf(val).Kind() {
			    case reflect.String:
					tableReference := TableReference{tableExpression: &Table{Name: val.(string)}}
					self.tableReferences = append(self.tableReferences, tableReference)
			    // case int:
			    //     fmt.Println("Select lists don't support integers in the select list right now")
			    // case []interface{}:
			    //     fmt.Println("This is an array: ", val)
			    // case map[string]interface{}:
			    //     fmt.Println("This is a map: ", val)
				default:
			        fmt.Println("No clue what this is: ", val)
			}
		}
	}
	
	return self
}

func (self *SelectStatement) Where(predicates ...interface{}) *SelectStatement {
    if self.predicates == nil {
        self.predicates = []Predicate{}
    }
    
    subpredicates := ParsePredicates(predicates...)
    self.predicates = append(self.predicates, subpredicates...)
	
	return self
}

func Select (selectList ...interface{}) *SelectStatement {
	fooble := []SQLExpression{}
	
	for _, val := range selectList {
		// _, ok := val.(map[string]string)
		//
		// if ok {
		// 	fmt.Println("Found a map")
		// }
		//
		// 	    fmt.Printf("%v is a map? %v\n", val, reflect.ValueOf(val).Kind() == reflect.Map)
		
	    switch reflect.ValueOf(val).Kind() {
		    case reflect.String:
				fooble = append(fooble, &Column{Name: val.(string)})
		    // case int:
		    //     fmt.Println("Select lists don't support integers in the select list right now")
		    // case []interface{}:
		    //     fmt.Println("This is an array: ", val)
		    // case map[string]interface{}:
		    //     fmt.Println("This is a map: ", val)
			default:
		        fmt.Println("No clue what this is: ", val)
		}
	}
	
	return &SelectStatement{selectList: fooble, tableReferences: []TableReference{}}
}
