package main

import (
	"fmt"
	"reflect"
)

func main() {
	ls := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}

	f := func(i int) bool {
		if i == 3 {
			return true
		}
		return false
	}

	// expt. 1 filter plain
	{
		res := filter(ls, f)
		fmt.Printf("res1: %+v\n", res)
	}

	// expt. 2 filter generics-like using interfaces
	{
		res := filterWithIface(ls, f)
		resInt := res.([]int)
		fmt.Printf("res2: %+v\n", resInt)
	}
}

func filter(ls []int, f func(i int) bool) []int {
	ret := make([]int, 0)

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
