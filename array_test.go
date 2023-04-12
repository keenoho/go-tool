package tool

import (
	"fmt"
	"testing"
)

func Test_Array(t *testing.T) {
	array := Array{}
	array2 := Array{}
	// array.Append(1)
	// fmt.Println(array)
	array.Append(1, 2, 3)
	fmt.Println(array)
	array.Prepend(4, 5, 6)
	fmt.Println(array)
	array.Pop()
	fmt.Println(array)
	array.Shift()
	fmt.Println(array)
	array.Insert(1, 666)
	fmt.Println(array)
	res := array.Slice(0, 2)
	fmt.Println(res)

	array2.Append(9, 99, 999)
	array.Concat(array2)

	fmt.Println(array)

	t.Log("ok")
}
