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
	SQL, values := statement.GenerateSQL()

	result := "delete from \"user\" where \"foo\" = $1 and \"blarg\" = $2 and \"wongle\" = null"
	if SQL != result {
		t.Error("Expected ", result, " got ", SQL)
	}

	expectedValues := []interface{}{10, "bar"}

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
