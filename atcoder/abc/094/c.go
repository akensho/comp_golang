package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n := intv()
	mid := (n - 1) / 2
	x := ints()
	o := make([]orignal, n)
	for i, _ := range x {
		o[i].key = i
		o[i].value = x[i]
	}
	sort.Sort(orignals(o))
	idxMap := make(map[int]int)
	for old, _ := range o {
		idxMap[old] = o[old].key
	}

	sort.Ints(x)
	//	fmt.Println(o)
	//	fmt.Println(x)
	for i := 0; i < n; i++ {
		idx := idxMap[i]
		before := []int{}
		after := []int{}
		array := make([]int, n-1)
		if idx == 0 {
			after = x[1:]
			fmt.Println(after[mid])
			continue
		} else if idx == n-1 {
			before = x[:idx]
			fmt.Println(before[mid])
			continue
		} else {
			before = x[:idx]
			after = x[idx:]
			for i := 0; i < len(before); i++ {
				array[i] = before[i]
			}
			for i := 0; i < len(after)-1; i++ {
				array[len(before)+i] = after[i+1]
			}
			fmt.Println(array[mid])
		}
		//		fmt.Printf("i = %d, idx = %d\n", i, idx)
		//		fmt.Printf("x      => %#v\n", x)
		//		fmt.Printf("before => %#v\n", before)
		//		fmt.Printf("after  => %#v\n", after)
		//		fmt.Println(array)
	}
}

type orignal struct {
	value int
	key   int
}

type orignals []orignal

func (o orignals) Len() int {
	return len(o)
}

func (o orignals) Less(i, j int) bool {
	return o[i].value < o[j].value
}

func (o orignals) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
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
