package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var ns = bufio.NewScanner(os.Stdin)

func read() (o, e []string) {
	ns.Scan()
	s := ns.Text()
	for _, v := range s {
		o = append(o, string(v))
	}
	ns.Scan()
	s = ns.Text()
	for _, v := range s {
		e = append(e, string(v))
	}
	return o, e
}

func main() {
	o, e := read()
	var s []string
	min := math.Min(float64(len(o)), float64(len(e)))
	for i := 0; i < int(min); i++ {
		s = append(s, o[i])
		s = append(s, e[i])
	}
	if len(o) > len(e) {
		s = append(s, o[len(o)-1])
	} else if len(e) > len(o) {
		s = append(s, e[len(e)-1])
	}
	fmt.Println(strings.Join(s, ""))
}
