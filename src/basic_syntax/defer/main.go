package main

import (
	"fmt"
)

func main() {
	one()
	fmt.Println("====================")
	two()
	fmt.Println("====================")
	three()
}

func one() {
	defer fmt.Print(" World")
	fmt.Print("Hello")
}

func two() {
	defer fmt.Print(" World")
	defer fmt.Print(" Go")
	fmt.Print("Hello")
}

// deferはLIFO

func three() {
	defer func() {
		fmt.Print(" Go")
		fmt.Print(" World")
	}()

	fmt.Print("Hello")
}
