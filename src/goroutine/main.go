package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	msg := "start"
	prmsg := func (nm string, n int) {
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Microsecond)
	}
	hello := func(n int) {
		const nm string = "hello"
		for i := 0; i < 10; i++ {
			msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}
	main := func(n int) {
		const nm string ="*main"
		for i := 0; i < 5; i++ {
			msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
	}
	go hello(60)
	main(100)
}



//--------------------------
// func hello(s string, n int) {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("<%d %s>", i, s)
// 		time.Sleep(time.Duration(n) * time.Millisecond)
// 	}
// }

// func main() {
// 	go hello("hello", 50)
// 	hello("bye!", 100)
// }