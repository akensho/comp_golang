package math

//
// s:= []rune("14325")
// next_permutation(s)
// fmt.Println(string(s)) => 14352
//
func next_permutation(p []rune) {
	i := len(p) - 2
	for p[i] >= p[i+1] {
		i--
	}
	j := len(p) - 1
	for p[j] <= p[i] {
		j--
	}
	p[i], p[j] = p[j], p[i]
	i, j = i+1, len(p)-1
	for i < j {
		p[i], p[j] = p[j], p[i]
		i, j = i+1, j-1
	}
}
