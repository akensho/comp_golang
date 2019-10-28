package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

var (
	sc  scanner
	MOD = 1000000007
)

type restaurant struct {
	city       string
	score, idx int
}

type restaurants []restaurant

func (r restaurants) Len() int {
	return len(r)
}

func (r restaurants) Less(i, j int) bool {
	if r[i].city == r[j].city {
		return r[i].score > r[j].score
	}
	return r[i].city < r[j].city
}

func (r restaurants) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func main() {
	n := sc.nextInt()
	rs := make([]restaurant, n)
	for i := 0; i < n; i++ {
		rs[i].city = sc.next()
		rs[i].score = sc.nextInt()
		rs[i].idx = i + 1
	}
	sort.Sort(restaurants(rs))
	for _, r := range rs {
		fmt.Println(r.idx)
	}
}

/* template functions */

func init() {
	sc = scanner{
		buf: make([]string, 0, 0),
		cur: 0,
		r:   bufio.NewReader(os.Stdin),
	}
}

type scanner struct {
	buf []string
	cur int
	r   *bufio.Reader
}

func (s *scanner) readln() {
	rbuf := make([]byte, 0, 0)
	for {
		line, prefix, err := s.r.ReadLine()
		if err != nil {
			panic(err)
		}
		rbuf = append(rbuf, line...)
		if prefix == false {
			break
		}
	}
	s.cur = 0
	s.buf = strings.Split(*(*string)(unsafe.Pointer(&rbuf)), " ")
}

func (s *scanner) isFull() bool {
	if s.cur == len(s.buf) {
		return true
	}
	return false
}

func (s *scanner) resetCur() {
	s.cur = 0
}

func (s *scanner) next() string {
	if s.cur == 0 {
		s.readln()
	}
	res := s.buf[s.cur]
	s.cur++
	if s.isFull() {
		s.resetCur()
	}
	return res
}

func (s *scanner) nexts() []string {
	s.readln()
	s.resetCur()
	return s.buf
}

func (s *scanner) nextInt() int {
	if s.cur == 0 {
		s.readln()
	}
	res, _ := strconv.Atoi(s.buf[s.cur])
	s.cur++
	if s.isFull() {
		s.resetCur()
	}
	return res
}

func (s *scanner) nextInts() []int {
	s.readln()
	res := make([]int, len(s.buf))
	for i, _ := range s.buf {
		res[i], _ = strconv.Atoi(s.buf[i])
	}
	s.resetCur()
	return res
}

func (s *scanner) nextFloat() float64 {
	if s.cur == 0 {
		s.readln()
	}
	res, _ := strconv.ParseFloat(s.buf[s.cur], 64)
	s.cur++
	if s.isFull() {
		s.resetCur()
	}
	return res
}

func (s *scanner) nextFloats() []float64 {
	s.readln()
	res := make([]float64, len(s.buf))
	for i, _ := range s.buf {
		res[i], _ = strconv.ParseFloat(s.buf[i], 64)
	}
	s.resetCur()
	return res
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func var_dump(value ...interface{}) {
	for _, v := range value {
		fmt.Fprintf(os.Stderr, "%#v\n", v)
	}
}
