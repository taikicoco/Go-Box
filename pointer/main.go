package main

import (
	"fmt"
)

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func main() {
	x := 10
	failedUpdate(&x)//10
	fmt.Println(x)
	update(&x)
	fmt.Println(x) //20
}