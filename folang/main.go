package main

import (
	"fmt"
)

func main() {
	li := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	ls := []string{"a", "b", "c", "d", "e", "d", "c", "b", "a"}

	{
		res := filter__int(li, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		})
		fmt.Printf("res3-a: %+v\n", res)
	}

	{
		res := filter__string(ls, func(s string) bool {
			if s == "c" {
				return true
			}
			return false
		})
		fmt.Printf("res3-b: %+v\n", res)
	}
}

func filter__int(ls []int, f func(int) bool) []int {
	ret := make([]int, 0)

	for _, v := range ls {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func filter__string(ls []string, f func(string) bool) []string {
	ret := make([]string, 0)

	for _, v := range ls {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
