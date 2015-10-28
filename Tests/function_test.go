package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestFunction(t *testing.T) {
	sqlFunction := sql.Func("crypt", "fooble", sql.Func("gen_salt", "bf"))
	SQL, values := sqlFunction.GenerateSQL()

	expectedValues := []interface{}{"fooble", "bf"}

	expectedSQL := "crypt($1, gen_salt($2))"
	if SQL != expectedSQL {
		t.Error("Expected ", expectedSQL, " got ", SQL)
	}

	if len(values) != len(expectedValues) {
		t.Error("Expected ", expectedValues, " got ", values)
	} else {
		for index, value := range values {
			if value != expectedValues[index] {
				t.Error("Expected ", expectedValues, " got ", values)
				break
			}
		}
	}
}
