package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestAndPredicate(t *testing.T) {
	andPredicate := sql.And(map[string]interface{}{
		"foo":   "bar",
		"blarg": 10,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"blarg" = $1 and "foo" = $2`
	expectedValues := []interface{}{10, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestOrPredicate(t *testing.T) {
	orPredicate := sql.Or(map[string]interface{}{
		"foo":   "bar",
		"blarg": 10,
	})
	actualSQL, actualValues := orPredicate.GenerateSQL()

	expectedSQL := `"blarg" = $1 or "foo" = $2`
	expectedValues := []interface{}{10, "bar"}
	
	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}
