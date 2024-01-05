package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// jsonファイルを読み込む
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	// Goの構造体にデコード
	var persons []Person
	if err := json.Unmarshal(data, &persons); err != nil {
		panic(err)
	}

	for i, person := range persons {
		if person.Name == "Taro" {
			persons[i].Age = 100
		}
	}

	// Goの構造体をjsonにエンコード
	jsonData, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
