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
    sortDescriptors []*SortDescriptor
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
	
	if len(self.sortDescriptors) > 0 {
		sqlFragments = []string{}

		for _, descriptor := range self.sortDescriptors {
			descriptorSQL, descriptorValues := descriptor.GenerateSQLWithContext(context)
			sqlFragments = append(sqlFragments, descriptorSQL)
			values = append(values, descriptorValues...)
		}

		SQL += fmt.Sprintf(" order by %s", strings.Join(sqlFragments, ", "))
	}
    
	return
}

func (self *SelectStatement) From(tables ...interface{}) *SelectStatement {
	if len(tables) > 0 {	
		for _, val := range tables {		
		    switch reflect.ValueOf(val).Kind() {
			    case reflect.String:
					tableReference := TableReference{tableExpression: &Table{Name: val.(string)}}
					self.tableReferences = append(self.tableReferences, tableReference)
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

func (self *SelectStatement) OrderBy(descriptors ...interface{}) *SelectStatement {
    if self.sortDescriptors == nil {
        self.sortDescriptors = []*SortDescriptor{}
    }
    
    sortDescriptors := ParseSortDescriptors(descriptors...)
    self.sortDescriptors = append(self.sortDescriptors, sortDescriptors...)
    return self
}

func Select (selectList ...interface{}) *SelectStatement {
	fooble := []SQLExpression{}
	
	for _, val := range selectList {
	    switch reflect.ValueOf(val).Kind() {
		    case reflect.String:
				fooble = append(fooble, &Column{Name: val.(string)})
			default:
		        fmt.Println("No clue what this is: ", val)
		}
	}
	
	return &SelectStatement{selectList: fooble, tableReferences: []TableReference{}}
}
