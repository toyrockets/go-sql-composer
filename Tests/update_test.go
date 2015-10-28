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
	SQL, values := statement.GenerateSQL()

	result := "update \"user\" set foo = $1, blarg = $2, wongle = null"
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

func TestUpdateWithReturningClause(t *testing.T) {
	statement := sql.Update("user").Set(map[string]interface{}{
		"foo":   "bar",
		"blarg": 10,
	}).Returning("*")
	SQL, values := statement.GenerateSQL()

	result := "update \"user\" set foo = $1, blarg = $2 returning *"
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
