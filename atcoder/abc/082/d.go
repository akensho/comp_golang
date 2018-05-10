package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in        = bufio.NewReader(os.Stdin)
	out       = bufio.NewWriter(os.Stdout)
	MOD       = 1e9 + 7
	UMOD      = uint64(1e9 + 7)
	factorial []uint64
	inverse   []uint64
)

func main() {
	s := strings.Split(strValue(), "")
	row := intSlice()
	x, y := row[0], row[1]
	var p point
	p.x, p.y, p.d = 0, 0, "r"
	q := make([]point, 1)
	q = append(q, p)
	for _, v := range s {
		if v == "T" {
			np := q[0]
			rp := np.t("r")
			//			dp := np.t("d")
			//			lp := np.t("l")
			//			up := np.t("u")
			q = append(q, rp)
			//			q = append(q, dp)
			//			q = append(q, lp)
			//			q = append(q, up)
		}
		if v == "F" {
			np := p.f()
			q = append(q, np)
		}
	}
	for _, qq := range q {
		if check(qq, x, y) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func check(p point, x int, y int) bool {
	if p.x == x && p.y == y {
		return true
	}
	return false
}

type point struct {
	x int
	y int
	d string
}

func (p *point) f() point {
	if p.d == "r" {
		p.x++
	} else if p.d == "d" {
		p.y--
	} else if p.d == "l" {
		p.x--
	} else if p.d == "u" {
		p.y++
	} else {
		panic("")
	}
	return *p
}

func (p *point) t(dir string) point {
	if dir == "r" {
		p.d = "d"
	} else if dir == "d" {
		p.d = "l"
	} else if dir == "l" {
		p.d = "u"
	} else if dir == "u" {
		p.d = "l"
	} else {
		panic("")
	}
	return *p

}

/* template functions */

func readln() string {
	buf := make([]byte, 0)
	for {
		line, prefix, err := in.ReadLine()
		if err != nil {
			panic(err)
		}
		buf = append(buf, line...)
		if prefix == false {
			break
		}
	}
	return string(buf)
}

func strValue() string {
	return strSlice()[0]
}

func strSlice() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intValue() int {
	return intSlice()[0]
}

func intSlice() []int {
	line := strSlice()
	slice := make([]int, 0)
	for _, tmp := range line {
		val, err := strconv.Atoi(tmp)
		if err != nil {
			panic(err)
		}
		slice = append(slice, val)
	}
	return slice
}

func IntMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func IntMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
