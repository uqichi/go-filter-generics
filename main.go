package main

import "fmt"

func main() {
	list := []int{1,2,3,4,5,4,3,2,1}
	fmt.Println(filter(list, func(i int) bool {
		if i == 3 {
			return true
		}
		return false
	}))
}

func filter(ls []int, f func(i int) bool) []int {
	var ret []int
	for _, v := range ls {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
