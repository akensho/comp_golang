package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var (
	ns = bufio.NewScanner(os.Stdin)
)

func read() (n int, s string) {
	ns.Scan()
	n, _ = strconv.Atoi(ns.Text())
	ns.Scan()
	s = ns.Text()
	return n, s
}

func main() {
	_, s := read()
	stack := list.New()
	for _, v := range s {
		stack.PushFront(v)
	}
	var res string
	for stack.Len() > 0 {
		now := stack.Front()
		stack.Remove(stack.Front())
		if stack.Len() == 0 {
			res += now.Value
			break
		}
		top := stack.Front()
		if now+top == "()" || top+now == "()" {
			stack.Remove(stack.Front())
		} else {
			res += now
		}
	}
	var ans string
	for _, v := range res {
		if v == "(" {
			ans += ")"
		} else {
			ans += "("
		}
	}
	fmt.Println(ans + s)
}
