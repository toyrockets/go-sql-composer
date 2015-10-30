package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestFunction(t *testing.T) {
	sqlFunction := sql.Func("crypt", "fooble", sql.Func("gen_salt", "bf"))
	actualSQL, actualValues := sqlFunction.GenerateSQL()

	expectedSQL := `crypt($1, gen_salt($2))`
	expectedValues := []interface{}{"fooble", "bf"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}
