package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read() (a, b int) {
	sc := initScanner()
	sc.readln()
	a, _ = strconv.Atoi(string(sc.buf[0]))
	b, _ = strconv.Atoi(string(sc.buf[1]))
	return a, b
}

func main() {
	a, b := read()
	fmt.Println(a, b)
}

type Scanner struct {
	r   *bufio.Reader
	buf []byte
}

func initScanner() *Scanner {
	i := bufio.NewReaderSize(os.Stdin, 1000000)
	b := make([]byte, 1000000)
	return &Scanner{r: i, buf: b}
}

func (s *Scanner) readln() {
	for {
		l, p, e := s.r.ReadLine()
		if e != nil {
			panic(e)
		}
		s.buf = append(s.buf, l...)
		if !p {
			break
		}
	}
}
