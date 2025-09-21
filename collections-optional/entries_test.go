package collections

import (
	"slices"
	"testing"
)

func TestEntries(t *testing.T) {
	input := map[int]string{2: "abc", 1: "def", 3: "ghi"}
	expected_output := []KeyValue{{1, "def"}, {2, "abc"}, {3, "ghi"}}

	actual_output := Entries(input)

	if !slices.Equal(actual_output, expected_output) {
		t.Errorf("Entries returned %v, expected %v", actual_output, expected_output)
	}
}
