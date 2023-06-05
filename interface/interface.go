package main

import (
	"fmt"
	"math"
	"io"
	"strings"
)

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
}


type LogicProvider struct {}

func (lp LogicProvider) Process(data string) string {
	fmt.Println("Processing data")
	return data
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) DoSomething() {
	c.L.Process("data")
}

func main1() {
	c := Client{L: LogicProvider{}}
	c.DoSomething()
}

// duck typing
// Python
// class Logic:
// def process(self, data):
//     # Process the data
//     # Return the result

// def program(logic):
//     # Create a program
//     # Return the program

// logToUse = Logic()
// program(logToUse)


func main2() {
	dog := Dog{}
	cat := Cat{}
	DoSpeak(&dog)
	DoSpeak(&cat)
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

func main3() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f
	a = &v

	// a = v
	fmt.Println(f.Abs(), "f.Abs()")
	fmt.Println(v.Abs(), "v.Abs()")

	fmt.Println(a.Abs(), "a.Abs(), a = &v")
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}else {
		return float64(f)
	}
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}	


func main4() {
	a := &Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, "a")
	fmt.Println(z, "z")
	
}

type Person struct {
	Name string
	Age int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


func main5() {
	r := strings.NewReader("Hello, Reader!")
	// func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
	fmt.Println(r, "r") // &{Hello, Reader! 0 -1} r
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}


