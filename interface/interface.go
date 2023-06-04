package main

import "fmt"

func main() {
	main1()
	main2()
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
