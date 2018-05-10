package math

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	s := []rune("1234")
	next_permutation(s)
	actual := string(s)
	expected := "1243"
	if actual != expected {
		t.Errorf("bag")
	}
}
