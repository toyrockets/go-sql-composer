package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestPredicateNullVariable(t *testing.T) {
	andPredicate := sql.And(map[string]interface{}{
		"c1": nil,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = null`
	expectedValues := []interface{}{}

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func RunPredicateVariableTest(t *testing.T, value interface{}) {
	andPredicate := sql.And(map[string]interface{}{
		"c1": value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateStringVariable(t *testing.T) {
	RunPredicateVariableTest(t, "value")
}

func TestPredicateBoolVariable(t *testing.T) {
	RunPredicateVariableTest(t, true)
}

func TestPredicateUintVariable(t *testing.T) {
	RunPredicateVariableTest(t, uint(42))
}

func TestPredicateUint8Variable(t *testing.T) {
	RunPredicateVariableTest(t, uint8(42))
}

func TestPredicateUint16Variable(t *testing.T) {
	RunPredicateVariableTest(t, uint16(42))
}

func TestPredicateUint32Variable(t *testing.T) {
	RunPredicateVariableTest(t, uint32(42))
}

func TestPredicateUint64Variable(t *testing.T) {
	RunPredicateVariableTest(t, uint64(42))
}

func TestPredicateIntVariable(t *testing.T) {
	RunPredicateVariableTest(t, int(42))
}

func TestPredicateInt8Variable(t *testing.T) {
	RunPredicateVariableTest(t, int8(42))
}

func TestPredicateInt16Variable(t *testing.T) {
	RunPredicateVariableTest(t, int16(42))
}

func TestPredicateInt32Variable(t *testing.T) {
	RunPredicateVariableTest(t, int32(42))
}

func TestPredicateInt64Variable(t *testing.T) {
	RunPredicateVariableTest(t, int64(42))
}

func TestPredicateFloat32Variable(t *testing.T) {
	RunPredicateVariableTest(t, float32(42.5))
}

func TestPredicateFloat64Variable(t *testing.T) {
	RunPredicateVariableTest(t, float64(42.5))
}

func TestPredicateStringPointerVariable(t *testing.T) {
	value := "value"
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateBoolPointerVariable(t *testing.T) {
	value := true
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateUintPointerVariable(t *testing.T) {
	value := uint(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateUint8PointerVariable(t *testing.T) {
	value := uint8(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateUint16PointerVariable(t *testing.T) {
	value := uint16(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateUint32PointerVariable(t *testing.T) {
	value := uint32(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateUint64PointerVariable(t *testing.T) {
	value := uint64(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateIntPointerVariable(t *testing.T) {
	value := int(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateInt8PointerVariable(t *testing.T) {
	value := int8(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateInt16PointerVariable(t *testing.T) {
	value := int16(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateInt32PointerVariable(t *testing.T) {
	value := int32(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateInt64PointerVariable(t *testing.T) {
	value := int64(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateFloat32PointerVariable(t *testing.T) {
	value := float32(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

func TestPredicateFloat64PointerVariable(t *testing.T) {
	value := float64(42)
	andPredicate := sql.And(map[string]interface{}{
		"c1": &value,
	})
	actualSQL, actualValues := andPredicate.GenerateSQL()

	expectedSQL := `"c1" = $1`
	expectedValues := []interface{}{ value }

	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}

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

func TestNotPredicate(t *testing.T) {
	orPredicate := sql.Not(map[string]interface{}{
		"foo":   "bar",
		"blarg": 10,
	})
	actualSQL, actualValues := orPredicate.GenerateSQL()

	expectedSQL := `not ("blarg" = $1 and "foo" = $2)`
	expectedValues := []interface{}{10, "bar"}
	
	CompareTestResults(t, expectedSQL, actualSQL, expectedValues, actualValues)
}
