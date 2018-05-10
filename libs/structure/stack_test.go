package structure

import (
	"container/list"
	"testing"
)

func TestStackPush(t *testing.T) {
	cnt := 10
	s := list.New()
	for i := 0; i < cnt; i++ {
		s.PushFront('a')
	}
	for i := cnt - 1; i >= 0; i-- {
		actual := 'a'
		expected := s.Remove(s.Front()).(rune)
		if actual != expected {
			t.Errorf("got %v : want %v", actual, expected)
		}
	}
}
