package search

import (
	"testing"
)

func TestLowerBound(t *testing.T) {
	a := []int{1, 11, 22, 33, 44, 55, 66}
	actual, err := lower_bound(a, 36)
	expected := 4
	if err != nil {
		t.Errorf("invalid")
	}
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
