package main

import (
	_"fmt"
	"io"
	"log"
	"os"
)

// func 関数名　(引数)　戻り値 {
// 	--処理--
// }

func file_open() {
	if len(os.Args) < 2 {
		log.Fatal("fail指定されてない")
	}
	f, err := os.Open(os.Args[1])
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func main()  {
	file_open()
}




