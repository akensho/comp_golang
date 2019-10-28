package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

var (
	sc        scanner
	factorial []uint64
	inverse   []uint64
	MOD       = uint64(1e9 + 7)
)

func main() {
	n, k := sc.nextInt(), sc.nextInt()
	combination(uint64(n))
	for i := 1; i <= k; i++ {
		fmt.Println(((nCm(uint64(n-k+1), uint64(i)) % MOD) * (nCm(uint64(k-1), uint64(i-1)) % MOD)) % MOD)
	}
}

func combination(n uint64) {
	factorial = make([]uint64, n+1)
	inverse = make([]uint64, n+1)
	factorial[0] = 1
	inverse[0] = 1
	for i := 1; i < len(factorial); i++ {
		factorial[i] = (factorial[i-1] * uint64(i)) % uint64(1e9+7)
		inverse[i] = pow(factorial[i], uint64(uint64(1e9+7)-2)) % uint64(1e9+7)
	}
}

// nCm
func nCm(n, m uint64) uint64 {
	if n < m {
		return 0
	}
	return factorial[n] * inverse[m] % uint64(1e9+7) * inverse[n-m] % uint64(1e9+7)
}

// x^y
func pow(x, y uint64) (res uint64) {
	res = 1
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % uint64(1e9+7)
		}
		x = (x * x) % uint64(1e9+7)
		y = y >> 1
	}
	return res
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
