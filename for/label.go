package main

import "fmt"

func main() { 
	samples := []string{"hello", "apple_π!", "これは漢字文字列"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' || r=='は' {
				continue outer 
			}
		}
		fmt.Println("===")
	}
} 

// 0 104 h
// 1 101 e
// 2 108 l
// 0 97 a
// 1 112 p
// 2 112 p
// 3 108 l
// 0 12371 こ
// 3 12428 れ
// 6 12399 は