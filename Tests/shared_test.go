package sqlcomposer_test

import (
	"testing"
)

func CompareTestResults(t *testing.T, expectedSQL string, actualSQL string, expectedValues []interface{}, actualValues []interface{}) {
	if expectedSQL != actualSQL {
		t.Errorf("Expected\n%s\ngot\n%s\n", expectedSQL, actualSQL)
	}

	if len(expectedValues) != len(actualValues) {
		t.Errorf("Expected\n%s\ngot\n%s\n", expectedValues, actualValues)
	} else {
		for index, value := range actualValues {
			if value != expectedValues[index] {
				t.Errorf("Expected\n%s\ngot\n%s\n", expectedValues, actualValues)
				break
			}
		}
	}
}
