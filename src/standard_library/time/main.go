package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	fmt.Println(startOfMonth)
}
