package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	go_marshal()
	go_new_encoder()
	go_file_open()
}

type go_sample_stract []struct {
	ID int `json:"id"`
	Answer string `json:"answer"`
}

func go_file_open() {

	f, err := os.Open("sample.json")
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var data go_sample_stract

	err = json.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

type GoStruct struct {
	A int
	B string
}

func go_marshal() {
	
	stcData := GoStruct{A: 1, B: "bbb"}

	// Marshal関数でjsonエンコード
	// ->返り値jsonDataにはエンコード結果が[]byteの形で格納される
	jsonData, err := json.Marshal(stcData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)
	fmt.Println(stcData)
}

func go_new_encoder() {
	stcData := GoStruct{A: 1, B: "bbb"}

	// 標準出力にjsonエンコード結果を出す
	err := json.NewEncoder(os.Stdout).Encode(stcData)
	if err != nil {
		fmt.Println(err)
	}
}

