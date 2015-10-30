package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestSimpleDelete(t *testing.T) {
	statement := sql.Delete("user").Where(map[string]interface{}{
		"foo":    "bar",
		"blarg":  10,
		"wongle": nil,
	})
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `delete from "user" where "blarg" = $1 and "foo" = $2 and "wongle" = null`
	expectedValues := []interface{}{10, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}
