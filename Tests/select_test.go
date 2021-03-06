package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestSimpleSelect(t *testing.T) {
	statement := sql.Select("a", "b", "c").From("t1", "t2", "t3").Where(map[string]interface{}{
		"foo":   "bar",
		"blarg": sql.GreaterThan(10),
	}).OrderBy("a")
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `select "a", "b", "c" from "t1", "t2", "t3" where "blarg" > $1 and "foo" = $2 order by "a" asc`
	expectedValues := []interface{}{10, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestSelectWithAliases(t *testing.T) {
	statement := sql.Select(map[string]interface{} {
		"a" : "c1",
		"b": "c.2",
		"c": nil,
	}).From(map[string]interface{}{
		"t1": "t.1",
		"t2": nil,
	});
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `select "a" as "c1", "b" as "c.2", "c" from "t1" as "t.1", "t2"`
	expectedValues := []interface{}{}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestSelectWithJoin(t *testing.T) {
	statement := sql.Select("a", "b", "c").From("t1", "t2", "t3").Where(map[string]interface{}{
		"foo":   "bar",
		"blarg": sql.GreaterThan(10),
	}).Join("t4", map[string]interface{}{
		"t4.id": &sql.SQLIdentifier{Name: "t3.parent_id"},
	}).OrderBy("a")
	actualSQL, actualValues := statement.GenerateSQL()

	expectedSQL := `select "a", "b", "c" from "t1", "t2", "t3" join "t4" on "t4"."id" = "t3"."parent_id" where "blarg" > $1 and "foo" = $2 order by "a" asc`
	expectedValues := []interface{}{10, "bar"}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

