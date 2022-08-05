package main

import (
	"fmt"
)

// func main() {
// 	ar := []int{10,20,30}
// 	fmt.Println(ar)
// 	initial(&ar)
// 	fmt.Println(ar)
// }

func initial(ar *[]int){
	for i := 0; i < len(*ar); i++ {
		(*ar)[i] = 0
	}
}

// func main() {
// 	n := 123
// 	fmt.Printf("value:%d.\n",n)
// 	change1(n)
// 	fmt.Printf("value:%d.\n",n)
// 	change2(&n)
// 	fmt.Printf("value:%d.\n",n)
// }

// //値渡し
// func change1(n int){
// 	n *= 2
// }
// //参照渡し
// func change2(n *int){
// 	*n *= 2
// }

// func main(){
// 	n := 123
// 	p := &n
// 	q := &p
// 	m := 1000
// 	p2 := &m
// 	q2 := &p2
// 	fmt.Printf("p value:%d, address:%p\n",**q,q)
// 	fmt.Printf("p2 value:%d, address:%p\n", **q2, q2)
// 	pb := p
// 	p = p2
// 	p2 = pb
// 	fmt.Printf("p value:%d, address:%p\n",**q,q)
// 	fmt.Printf("p2 value:%d, address:%p\n", **q2, q2)
// }


// var 変数 *型 = &値
// ポインタ型　アドレスを代入
// 変数の前に&をつけることでアドレスの値を取り出せる
//値を取り出す時[*]をつける


// func main() {
// 	n := 123
// 	p := &n
// 	fmt.Println("number",n)
// 	fmt.Println("pointer",p)
// 	fmt.Println("value",*p)
// }

//ポインタを引数に使う

func main() {
	n := 123
	fmt.Printf("value:%d.\n",n)
	change1(n)
	fmt.Printf("value:%d.\n",n)
	change2(&n)
	fmt.Printf("value:%d.\n",n)
}

func change1(n int) {
	n *= 2
}

func change2(n *int) {
	*n *= 2
}