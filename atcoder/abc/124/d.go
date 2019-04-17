package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	INF = (1 << 32) - 1
)

func main() {
	row := ints()
	n, k := row[0], row[1]
	s := strv()
	list := func() []int {
		res := make([]int, len(s))
		for i, val := range s {
			if val == '0' {
				res[i] = 0
			} else {
				res[i] = 1
			}
		}
		return res
	}()
	list = comvert(list)
	fmt.Println(list)
	cnt, max := 0, 0
	for {
		if cnt == k {
			fmt.Println(max)
			return
		}
		if max == n {
			fmt.Println(n)
			return
		}
		set := make([][]int, 0)
		for i := 0; i < len(list); i++ {
			if list[i] != 0 {
				continue
			}
			tmp := compress(list, i)
			set = append(set, tmp)
		}
		fmt.Println(set)
		cnt++
	}

}

func compress(a []int, x int) []int {
	s, e := x, x
	for i := x; i < len(a); i++ {
		if a[i] == 0 {
			a[i] = 1
			continue
		}
		e = i
		break
	}
	mid := []int{func() (sum int) {
		for _, val := range a[s:e] {
			sum += val
		}
		return sum
	}()}
	fmt.Printf("s = %d, e = %d, pre = %v, mid = %v, post = %v\n", s, e, a[:s], mid, a[e:])
	res := make([]int, 0)
	if s == 0 {
		res = append(mid, a[e:]...)
	} else if e == len(a)-1 {
		res = append(a[:s], mid...)
	} else {
		res = append(a[:s], mid...)
		res = append(res, a[e:]...)
	}
	return res
}

func comvert(a []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			res = append(res, 0)
			continue
		}
		cnt := a[i]
		for j := i + 1; j < len(a); j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if a[j] != 0 {
				cnt += a[j]
				fmt.Printf("cnt = %d\n", cnt)
				continue
			}
			res = append(res, cnt)
			res = append(res, 0)
			i = j
			break
		}
	}
	return res
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

func strv() string {
	return strs()[0]
}

func strs() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intv() int {
	return ints()[0]
}

func ints() []int {
	line := strs()
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

func floatv() float64 {
	return floats()[0]
}

func floats() []float64 {
	line := strs()
	slice := make([]float64, 0)
	for _, tmp := range line {
		val, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			panic(err)
		}
		slice = append(slice, val)
	}
	return slice
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func var_dump(value ...interface{}) {
	for _, v := range value {
		fmt.Printf("%#v\n", v)
	}
}
