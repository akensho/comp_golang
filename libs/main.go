package main

import (
	"fmt"

	"github.com/akensho/comp_golang/libs/search"
)

func main() {
	a := []int{11, 22, 33, 44, 55}
	idx := search.Lower_bound(a, 43)
	fmt.Println(idx)
}
