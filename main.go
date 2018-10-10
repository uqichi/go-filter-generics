package main

import (
	"fmt"
	"reflect"
)

func main() {
	li := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	ls := []string{"a", "b", "c", "d", "e", "d", "c", "b", "a"}

	// expt. 1: int filter plain
	{
		res := filterInt(li, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		})
		fmt.Printf("res1-a: %+v\n", res)
	}

	// expt. 2: string filter plain
	{
		res := filterString(ls, func(s string) bool {
			if s == "c" {
				return true
			}
			return false
		})
		fmt.Printf("res1-b: %+v\n", res)
	}

	// expt. 2-a: int filter generics-like using interfaces
	{
		res := filterWithIface(li, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		})
		resInt := res.([]int)
		fmt.Printf("res2-a: %+v\n", resInt)
	}

	// expt. 2-b: string filter generics-like using interfaces
	{
		res := filterWithIface(ls, func(s string) bool {
			if s == "c" {
				return true
			}
			return false
		})
		resString := res.([]string)
		fmt.Printf("res2-b: %+v\n", resString)
	}
}

func filterInt(ls []int, f func(i int) bool) []int {
	ret := make([]int, 0)

	for _, v := range ls {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func filterString(ls []string, f func(s string) bool) []string {
	ret := make([]string, 0)

	for _, v := range ls {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func filterWithIface(ls interface{}, f interface{}) interface{} {
	lsVal := reflect.ValueOf(ls)
	fVal := reflect.ValueOf(f)

	ret := reflect.MakeSlice(reflect.TypeOf(ls), 0, lsVal.Len())

	for i := 0; i < lsVal.Len(); i++ {
		b := fVal.Call([]reflect.Value{lsVal.Index(i)})[0]
		if b.Bool() {
			ret = reflect.Append(ret, lsVal.Index(i))
		}
	}

	return ret.Interface()
}
