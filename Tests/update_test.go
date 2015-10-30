package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestSimpleUpdate(t *testing.T) {
	statement := sql.Update("user").Set(map[string]interface{}{
		"foo":    "bar",
		"blarg":  10,
		"wongle": nil,
	})
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `update "user" set "blarg" = $1, "foo" = $2, "wongle" = null`
	expectedValues := []interface{}{10, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestUpdateWithReturningClause(t *testing.T) {
	statement := sql.Update("user").Set(map[string]interface{}{
		"foo":   "bar",
		"blarg": 10,
	}).Returning("*")
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `update "user" set "blarg" = $1, "foo" = $2 returning *`
	expectedValues := []interface{}{10, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}
