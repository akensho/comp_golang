package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

var (
	in   = bufio.NewReader(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
	INF  = (1 << 32) - 1
)

func main() {
	s := intv()
	base := []int{357, 375, 537, 573, 735, 753}
	save := make([][]int, 0)
	save = append(save, base)
	for i := 1000; i<=1000000000; i=i*10{
		l := save[len(save)-1:][0]
		tmp  := f(l, i, 3)
		tmp2 := f(l, i, 5)
		tmp3 := f(l, i, 7)
		list := append(tmp, tmp2...)
		list = append(list, tmp3...)
		save = append(save, list)
	}
	list := make([]int, 0)
	for _, li := range save {
		for _, val := range li {
			list = append(list, val)
		}
	}
	sort.Ints(list)
	ans := 0
	for _, val := range list {
		if val <= s {
			ans++
			fmt.Println(val)
		}
	}
	fmt.Println(ans)
}

func f(a []int, x ,y int)[]int{
	tmp := make([]int, 0)
	for _, val := range a {
		t := x*y+val
		tmp = append(tmp, t)
	}
	return tmp
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
