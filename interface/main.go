package main
import (
	"fmt"
	"strings"
	"strconv"
)

//Data is interface for Mydata
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// Mydata is structure
type Mydata struct {
	Name string
	Data []int
}
//SetValue is Mydata method 
func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

//PrintData is Mydata method
func (md *Mydata) PrintData(){
	fmt.Println("Name:", md.Name)
	fmt.Println("Data:", md.Data)
}

//Yourdata is structure
type Yourdata struct {
	Name string
	Mail string
	Age int
}

//SetValu is Yourdata method
func (md *Yourdata) SetValue(vals map[string]string){
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

//PrintData is Yourdata method
func (md *Yourdata)PrintData() {
	fmt.Printf("I'm %s. (%d).\n",md.Name,md.Age)
	fmt.Printf("mail:%s.\n",md.Mail)
}

func main() {
	ob := [2]Data{}

	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name":"taiki",
		"data":"23,44,12",
	})

	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name":"taro",
		"mail":"taro@gmail.wo",
		"aga": "22",
	})

	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}


// //Data is interface
// type Data interface {
// 	Initial(name string, data []int)
// 	PrintData()
// }

// //Mydata is Struct
// type Mydata struct {
// 	Name string
// 	Data []int
// }

// //Initial is init method
// func (md *Mydata) Initial(name string, data []int) {
// 	md.Name = name
// 	md.Data = data
// }

// //PfintData is printIn all data
// func (md *Mydata) PrintData() {
// 	fmt.Println("Name:", md.Name)
// 	fmt.Println("Data:", md.Data)
// }

// func main() {
// 	var ob Data = new(Mydata)
// 	ob.Initial("taiki", []int{3,4,5})
// 	ob.PrintData()
// }

// // func main () {
// // 	var ob Mydata = Mydata{}
// // 	ob.Initial("Taiki", []int{11,22,33})
// // 	ob.PrintData()
// // }
