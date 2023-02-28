package main

import (
	"fmt"
)

func f1() {
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println(pointerX, pointerY, pointerZ)
}

func f2() {
	// 「*」は間接参照
	// ポインタ型変数の前に付けると、そのポインタが参照するアドレスに保存されている値を返す
	// これをデリファレンスと呼ぶ
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX) // アドレスを表示
	fmt.Println(*pointerToX) // 10 // デリファレンスする

	var y *int
	fmt.Println(y == nil) // true
	fmt.Println(*y) // パニックになる
}

// 値渡しと参照渡し
func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func f3() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // 10
	update(&x)
	fmt.Println(x) // 20
}

func main() {
	f1()
	f2()
	f3()
}