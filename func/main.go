package main

import "fmt"

// func 関数名　(引数)　戻り値 {
// 	--処理--
// }

//関数とクロージャ
func main() {
	data := "*新しい値*"
	m1 := modify(data)
	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())
}

func modify(d string) func() []string {
	m := []string{
		"1st","2nd",
	}
	return func() []string {
		return append(m, d)
	}
}

//高度関数
// func main() {
// 	modify := func(a []string, f func([]string) []string) []string {
// 		return f(a)
// 	}

// 	m := []string{
// 		"1st","2nd","3rd",
// 	}
// 	fmt.Println(m)

// 	m1 := modify(m, func([]string) []string {
// 		return append(m, m...)
// 	})
// 	fmt.Println(m1)

// 	m2 := modify(m, func([]string) []string {
// 		return m[:len(m)-1]
// 	})
// 	fmt.Println(m2)

// 	m3 := modify(m, func([]string) []string {
// 		return m[1:]
// 	})
// 	fmt.Print(m3)
// }

//無名関数
// func main() {
// 	f := func(a []string)([]string, string) {
// 		return a[1:],a[0]
// 	}
// 	m := []string{
// 		"one",
// 		"two",
// 		"three",
// 	}
// 	s := ""
// 	fmt.Println(m)
// 	for len(m) > 0 {
// 		m, s = f(m)
// 		fmt.Println(s + "->",m)
// 	}
// }

//可変長変数
// func main() {
// 	m := []string{
// 		"one","two","three",
// 	}
// 	fmt.Println(m)
// 	m = push(m, "1","2","3")
// 	fmt.Println(m)
// }

// func push(a []string, v ...string) (s []string) {
// 	s = append(a, v...)
// 	return
// }

//名前付き戻り値
// func main(){
// 	m := []string{
// 		"one","two","three",
// 	}
// 	fmt.Println(m)
// 	m = insert(m, "*",2)
// 	m = insert(m, "*",1)
// 	fmt.Println(m)
// }

// func insert(a []string, v string, p int) (s []string) {
// 	s = append(a, "")
// 	s = append(s[:p+1], s[p:len(s)-1]...)
// 	s[p] = v
// 	return
// }


//複数の戻り値
// func push(a []string, v string) ([]string, int) {
// 	return append(a, v), len(a)
// }

// func pop(a []string) ([]string, string) {
// 	return a[:len(a)-1], a[len(a)-1]
// }

// func main() {
// 	m := []string{}
// 	m, _= push(m, "apple")
// 	m, _= push(m, "banana")
// 	m, a := push(m, "orange")
// 	fmt.Println(a)
// 	fmt.Println(m)
// 	m, v := pop(m)
// 	fmt.Println("get " + v + "->", m)
// }


//関数
// func push(a []int, v int) []int {
// 	return append(a, v)
// }

// func pop(a []int) []int {
// 	return a[:len(a)-1]
// }

// func main() {
// 	a := []int{10,20,30}
// 	fmt.Println(a)
// 	a = push(a, 1000)
// 	fmt.Println(a)
// 	a = pop(a)
// 	fmt.Println(a)
// }

