package mathmagic

import "testing"

const defaultError = "Expected value %v, got %v"

func TestAverage(t *testing.T) {
	expected := 7.28
	value := Average(7.2, 9.9, 5.5, 6.5)

	if value != expected {
		t.Errorf(defaultError, expected, value)
	}
}
