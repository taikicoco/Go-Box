package main

import "fmt"

//PrintData is println all data
func (md Mydata) PrintData() {
	fmt.Println("** Mydata **")
	fmt.Println("name", md.Name)
	fmt.Println("data", md.Data)
	fmt.Println("** end **")
}

func main() {
	taiki := Mydata{
		"taiki",[]int{2,4,5},
	}
	taiki.PrintData()
}
// //newで構造体を作成する
// func main() {
// 	taiki := new(Mydata)
// 	fmt.Println(taiki)
// 	taiki.Name = "taiki"
// 	taiki.Data = make([]int,5,5)
// 	fmt.Println(taiki)
// }

// //構造体と参照渡し
// func main() {
// 	taiki := Mydata{
// 		"taiki",
// 		[]int{1,1,3},
// 	}
// 	fmt.Println(taiki)
// 	rev(&taiki)
// 	fmt.Println(taiki)
// }
// func rev(md *Mydata) {
// 	od := (*md).Data
// 	nd := []int{}
// 	for i := len(od) -1; i >= 0; i-- {
// 		nd = append(nd, od[i])
// 	}
// 	md.Data = nd
// }


//tyepで型として定義する
//Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

// func main(){
// 	taiki := Mydata{"taiki",[]int{1,2,3}}//型として定義できる
// 	fmt.Println(taiki)
// }