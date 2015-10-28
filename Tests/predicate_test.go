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
	SQL, values := andPredicate.GenerateSQL()

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

	expectedSQL := "\"foo\" = $1 and \"blarg\" = $2"
	if SQL != expectedSQL {
		t.Error("Expected ", expectedSQL, " got ", SQL)
	}

}
