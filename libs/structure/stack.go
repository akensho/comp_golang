package structure

type Stack struct {
	slice []rune
}

func (s Stack) Init() {
	s.slice = make([]rune, 0)
}

func (s *Stack) Push(r rune) {

}

func (s *Stack) Pop() (r rune) {
	return r
}
