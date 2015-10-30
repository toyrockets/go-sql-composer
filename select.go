package sqlcomposer

import (
	"fmt"
	"sort"
	"strings"
)

type SelectStatement struct {
	selectList      ColumnList
	tableReferences []*TableReference
	joins           []*Join
	predicates      []Predicate
	sortDescriptors []*SortDescriptor
}

func (self *SelectStatement) GenerateSQL() (SQL string, values []interface{}) {
	DefaultSQLGenerationContext.reset()
	SQL, values = self.GenerateSQLWithContext(DefaultSQLGenerationContext)
	return
}

func (self *SelectStatement) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	var sqlFragments []string
	SQL = "select"
	values = []interface{}{}

	sqlFragments = []string{}

	sort.Stable(self.selectList)
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

	if len(self.joins) > 0 {
		sqlFragments = []string{}

		for _, join := range self.joins {
			joinSQL, joinValues := join.GenerateSQLWithContext(context)
			sqlFragments = append(sqlFragments, joinSQL)
			values = append(values, joinValues...)
		}

		SQL += fmt.Sprintf(" %s", strings.Join(sqlFragments, " "))
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
	if self.tableReferences == nil {
		self.tableReferences = []*TableReference{}
	}

	if len(tables) > 0 {
		for _, val := range tables {
			if stringValue, ok := val.(string); ok {
				tableReference := &TableReference{tableExpression: &Table{Name: stringValue}}
				self.tableReferences = append(self.tableReferences, tableReference)
			} else if mapValue, ok := val.(map[string]interface{}); ok {
				tableReferences := ParseFromMap(mapValue)
				self.tableReferences = append(self.tableReferences, tableReferences...)
			} else {
				fmt.Println("No clue what this is: ", val)
			}
		}
	}

	return self
}

func (self *SelectStatement) join(joinType JoinType, table string, predicates ...interface{}) *SelectStatement {
	if self.joins == nil {
		self.joins = []*Join{}
	}

	tableReference := &TableReference{tableExpression: &Table{Name: table}}
	joinPredicates := And(predicates...)
	join := &Join{joinType: joinType, tableReference: tableReference, condition: joinPredicates}
	self.joins = append(self.joins, join)

	return self
}

func (self *SelectStatement) Join(table string, predicates ...interface{}) *SelectStatement {
	self.join(InnerJoin, table, predicates...)
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

func Select(selectList ...interface{}) *SelectStatement {
	references := []*ColumnReference{}

	for _, val := range selectList {
		if stringValue, ok := val.(string); ok {
			references = append(references, &ColumnReference{expression: &SQLIdentifier{Name: stringValue}})
		} else if mapValue, ok := val.(map[string]interface{}); ok {
			columnReferences := ParseSelectMap(mapValue)
			references = append(references, columnReferences...)
		} else {
			fmt.Println("No clue what this is: ", val)
		}
	}

	return &SelectStatement{selectList: references}
}

func ParseSelectMap(values map[string]interface{}) []*ColumnReference {
	references := []*ColumnReference{}

	for key, value := range values {
		var alias *SQLAlias = nil

		switch aliasName := value.(type) {
		case string:
			alias = &SQLAlias{Name: aliasName}
		default:
			alias = nil
		}

		reference := &ColumnReference{expression: &SQLIdentifier{Name: key}, alias: alias}
		references = append(references, reference)
	}

	return references
}

func ParseFromMap(values map[string]interface{}) []*TableReference {
	references := []*TableReference{}

	for key, value := range values {
		var alias *SQLAlias = nil

		switch aliasName := value.(type) {
		case string:
			alias = &SQLAlias{Name: aliasName}
		default:
			alias = nil
		}

		reference := &TableReference{tableExpression: &SQLIdentifier{Name: key}, alias: alias}
		references = append(references, reference)
	}

	return references
}
