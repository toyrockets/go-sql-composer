package sqlcomposer

import (
	"fmt"
	"strings"
)

type ComparisonOperator string

const (
	EqualOperator              ComparisonOperator = "="
	NotEqualOperator           ComparisonOperator = "!="
	LessThanOperator           ComparisonOperator = "<"
	LessThanOrEqualOperator    ComparisonOperator = "<="
	GreaterThanOperator        ComparisonOperator = ">"
	GreaterThanOrEqualOperator ComparisonOperator = ">"
	InOperator                 ComparisonOperator = "in"
)

type Predicate interface {
	GeneratePredicateSQL() (SQL string, values []interface{})
	GenerateSQL() (SQL string, values []interface{})
	GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{})
}

type ComparisonPredicate struct {
	leftValue  *SQLExpression
	rightValue *SQLExpression
	operator   ComparisonOperator
}

func (self *ComparisonPredicate) GeneratePredicateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *ComparisonPredicate) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *ComparisonPredicate) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	var left SQLExpression = *(self.leftValue)
	var right SQLExpression = *(self.rightValue)
	leftSQL, leftValues := left.GenerateSQLWithContext(context)
	rightSQL, rightValues := right.GenerateSQLWithContext(context)

	values = []interface{}{}
	values = append(values, leftValues...)
	values = append(values, rightValues...)

	SQL = strings.Join([]string{leftSQL, string(self.operator), rightSQL}, " ")
	return
}

type AndPredicate struct {
	predicates []Predicate
}

func (self *AndPredicate) GeneratePredicateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *AndPredicate) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *AndPredicate) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	var sqlFragments []string

	for _, predicate := range self.predicates {
		predicateSQL, predicateValues := predicate.GenerateSQLWithContext(context)

		_, ok := predicate.(*OrPredicate)

		if ok {
			predicateSQL = "(" + predicateSQL + ")"
		}

		sqlFragments = append(sqlFragments, predicateSQL)
		values = append(values, predicateValues...)
	}

	SQL = strings.Join(sqlFragments, " and ")
	return
}

type OrPredicate struct {
	predicates []Predicate
}

func (self *OrPredicate) GeneratePredicateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *OrPredicate) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *OrPredicate) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	var sqlFragments []string
	values = []interface{}{}

	for _, predicate := range self.predicates {
		predicateSQL, predicateValues := predicate.GenerateSQLWithContext(context)

		_, ok := predicate.(*AndPredicate)

		if ok {
			predicateSQL = "(" + predicateSQL + ")"
		}

		sqlFragments = append(sqlFragments, predicateSQL)
		values = append(values, predicateValues...)
	}

	SQL = strings.Join(sqlFragments, " or ")
	return
}

type NotPredicate struct {
	predicate Predicate
}

func (self *NotPredicate) GeneratePredicateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *NotPredicate) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *NotPredicate) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	SQL, values = self.predicate.GenerateSQLWithContext(context)
	SQL = fmt.Sprintf("not (%s)", SQL)
	return
}

type InPredicate struct {
	left   SQLExpression
	values []SQLExpression
}

func (self *InPredicate) GeneratePredicateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQL()
	return
}

func (self *InPredicate) GenerateSQL() (SQL string, values []interface{}) {
	SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
	return
}

func (self *InPredicate) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	var sqlFragments []string

	for _, value := range self.values {
		sqlFragment, _ := value.GenerateSQLWithContext(context)
		sqlFragments = append(sqlFragments, sqlFragment)
	}

	SQL = strings.Join(sqlFragments, ", ")
	return
}

type BetweenPredicate struct {
	left       SQLExpression
	lowerBound SQLExpression
	upperBound SQLExpression
}

func ParsePredicates(values ...interface{}) []Predicate {
	if len(values) == 0 {
		return []Predicate{}
	}

	predicates := []Predicate{}

	for _, value := range values {
		predicateMap, ok := value.(map[interface{}]interface{})

		if ok {
			subpredicates := ParsePredicateMap(predicateMap)
			predicates = append(predicates, subpredicates...)
			continue
		}

		predicateMap2, ok := value.(map[string]interface{})

		if ok {
			fooble := map[interface{}]interface{}{}

			for k, v := range predicateMap2 {
				fooble[k] = v
			}
			subpredicates := ParsePredicateMap(fooble)
			predicates = append(predicates, subpredicates...)
			continue
		}

		predicate, ok := value.(Predicate)

		if ok {
			predicates = append(predicates, predicate)
			continue
		}

		fmt.Printf("No clue what this is 2: %#v\n", value)
	}

	return predicates
}

func ParsePredicateMap(values map[interface{}]interface{}) []Predicate {
	var predicates []Predicate
	for key, value := range values {
		var leftValue SQLExpression

		stringValue, ok := key.(string)

		if ok {
			leftValue = &SQLIdentifier{Name: stringValue}
		} else {
			leftValue = SQLVariable(key)
		}

		rightValue := ParsePredicateRightValue(value)

		predicate, ok := value.(*ComparisonPredicate)

		if ok {
			predicate.leftValue = &leftValue
		} else {
			predicate = &ComparisonPredicate{leftValue: &leftValue, rightValue: &rightValue, operator: EqualOperator}
		}

		predicates = append(predicates, predicate)
	}

	return predicates
}

func ParsePredicateRightValue(value interface{}) SQLExpression {
	if predicateValue, ok := value.(SQLExpression); ok {
		return predicateValue
	} else {
		return SQLVariable(value)
	}
}

//
// Conjunctive
//

func And(values ...interface{}) *AndPredicate {
	predicates := ParsePredicates(values...)
	return &AndPredicate{predicates: predicates}
}

func Or(values ...interface{}) *OrPredicate {
	predicates := ParsePredicates(values...)
	return &OrPredicate{predicates: predicates}
}

func Not(values ...interface{}) *NotPredicate {
	andPredicate := And(values...)
	return &NotPredicate{predicate: andPredicate}
}

//
// Comparison
//

func Equal(value interface{}) *ComparisonPredicate {
	rightValue := ParsePredicateRightValue(value)
	return &ComparisonPredicate{leftValue: nil, rightValue: &rightValue, operator: EqualOperator}
}

func NotEqual(value interface{}) *ComparisonPredicate {
	rightValue := ParsePredicateRightValue(value)
	return &ComparisonPredicate{leftValue: nil, rightValue: &rightValue, operator: NotEqualOperator}
}

func LessThan(value interface{}) *ComparisonPredicate {
	rightValue := ParsePredicateRightValue(value)
	return &ComparisonPredicate{leftValue: nil, rightValue: &rightValue, operator: LessThanOperator}
}

func LessThanOrEqual(value interface{}) *ComparisonPredicate {
	rightValue := ParsePredicateRightValue(value)
	return &ComparisonPredicate{leftValue: nil, rightValue: &rightValue, operator: LessThanOrEqualOperator}
}

func GreaterThan(value interface{}) *ComparisonPredicate {
	rightValue := ParsePredicateRightValue(value)
	return &ComparisonPredicate{leftValue: nil, rightValue: &rightValue, operator: GreaterThanOperator}
}

func GreaterThanOrEqual(value interface{}) *ComparisonPredicate {
	rightValue := ParsePredicateRightValue(value)
	return &ComparisonPredicate{leftValue: nil, rightValue: &rightValue, operator: GreaterThanOrEqualOperator}
}
