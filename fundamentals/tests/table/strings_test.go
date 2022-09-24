package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - index: expected (%d) <> found (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{text: "I'm amazing", part: "amazing", expected: 4},
		{text: "", part: "", expected: 0},
		{text: "Ops, I spoke too much", part: "ops", expected: -1},
		{text: "Leonard", part: "o", expected: 2},
	}

	for _, test := range tests {
		t.Logf("Mass: %v", test)

		current := strings.Index(test.text, test.part)

		if current != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, current)
		}
	}
}
