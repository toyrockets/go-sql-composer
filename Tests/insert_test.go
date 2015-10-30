package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
	"time"
)

func TestSimpleInsert(t *testing.T) {
	time := time.Now()
	statement := sql.Insert("user").Values(map[interface{}]interface{}{
		"foo":        "bar",
		"blarg":      10,
		"created_at": time,
	})
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `insert into "user" ("blarg", "created_at", "foo") values ($1, $2, $3)`
	expectedValues := []interface{}{10, time, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestInsertWithSubSelect(t *testing.T) {
	selectStatement := sql.Select("id").From("table1").Where(map[interface{}]interface{}{
		"external_id": 10,
	})

	statement := sql.Insert("user").Values(map[interface{}]interface{}{
		"foo":   "bar",
		"blarg": 10,
		"id":    selectStatement,
	})
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `insert into "user" ("blarg", "foo", "id") values ($1, $2, (select "id" from "table1" where "external_id" = $3))`
	expectedValues := []interface{}{10, "bar", 10}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestInsertWithFunction(t *testing.T) {
	statement := sql.Insert("user").Values(map[interface{}]interface{}{
		"password": sql.Func("crypt", "fooble", sql.Func("gen_salt", "bf")),
	})
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `insert into "user" ("password") values (crypt($1, gen_salt($2)))`
	expectedValues := []interface{}{"fooble", "bf"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestInsertWithReturning(t *testing.T) {
	time := time.Now()
	statement := sql.Insert("user").Values(map[interface{}]interface{}{
		"foo":        "bar",
		"blarg":      10,
		"created_at": time,
	}).Returning("*")
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `insert into "user" ("blarg", "created_at", "foo") values ($1, $2, $3) returning *`
	expectedValues := []interface{}{10, time, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

