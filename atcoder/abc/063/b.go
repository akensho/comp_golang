package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	ns = bufio.NewScanner(os.Stdin)
)

func read() string {
	ns.Scan()
	s := ns.Text()
	return s
}

func main() {
	s := read()
	a := strings.Split(s, "")
	sort.Strings(a)
	var f bool = true
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			f = false
		}
	}
	if f {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
