package main

import (
	"fmt"
)

func main() {
	assignVar()
	useIota()
	useConditionalBranching()
	useArraySlice()

	dog := &Dog{}
	cat := &Cat{}
	DoSpeak(dog)
	DoSpeak(cat)
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

	const _const int = 100
	const _const2 int = 200 //定数は使用しなくてもコンパイルエラーにならない
	fmt.Println("_const =", _const)

	// 定数でも使用時に方が決まる
	const x = 1
	y := 1 + x
	z := 2.1 + x
	fmt.Println("y =", y)
	fmt.Println("z =", z)
}
func useIota() {
	const (
		_ = iota
		Red
		Blue
		Green
		Black
	)
	fmt.Println("Red =", Red)
	fmt.Println("Blue =", Blue)
	fmt.Println("Green =", Green)
	fmt.Println("Black =", Black)

	const (
		apple = iota
		orange
		banana
		peche
	)
	fmt.Println("apple =", apple)
	fmt.Println("orange =", orange)
	fmt.Println("banana =", banana)
	fmt.Println("peche =", peche)
}
func useConditionalBranching(){
	x := 2
	y := 3

	if x == 2 && y == 3 {
		fmt.Println("x == 2 && y == 3", x, y)
	}

	// breakを書かないで処理が止まる
	z := 2
	switch z {
	case 1:
		fmt.Println("x == 1")
	case 2:
		fmt.Println("x == 2")
		fallthrough // breakしない "x == 2"と"x == 3"が出力される
	case 3, 4:
		fmt.Println("x == 3 or 4")
	default:
		fmt.Println("x == other")
	}
}
func useArraySlice() {
	var a [5]int
	a[2] = 3
	fmt.Println(a)

	// var b []int
	// 要素にアクセスするとpanicになる
	// b[2] = 3
	// fmt.Println(b)

	b := make([]int, 5)
	b[2] = 3
	fmt.Println(b)

	var c []int
	c = append(c, 1)
	c = append(c, 2)
	c = append(c, 3, 4, 5)
	fmt.Println(c)

	d := []int{1, 2, 3, 4, 5}
	fmt.Println(d)
	d = append(d, 6)
	fmt.Println(d)
}


// interface
type Speaker interface {
	Speak() error
}

type Dog struct {}
func (d *Dog) Speak() error {
	fmt.Println("Bow Wow")
	return nil
}

type Cat struct {}
func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}
func DoSpeak(speaker Speaker) error {
	return speaker.Speak()
}

