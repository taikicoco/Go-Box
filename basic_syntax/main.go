package main

import (
	"fmt"
)

func main() {
	assignVar()
}

// 変数の宣言と代入
func assignVar(){
	var _int int = 10
	fmt.Println("_int =", _int)
	_int += 20
	fmt.Println("_int =", _int)

	var _float float32 = 3.14
	fmt.Println("_float =", _float)

	var _string string = "Hello World"
	fmt.Println("_string =", _string)
	_string = "Hello," + "World"
	fmt.Println("__string =", _string)

	var _bool bool = true
	fmt.Println("_bool =", _bool)

	var _array [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("_array =", _array)

	var _slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println("_slice =", _slice)

	var _map map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("_map =", _map)
}