package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello, playground")

	now := time.Now()
	sec := now.Unix()

	fmt.Println(sec)
	fmt.Println(strconv.FormatInt(sec, 10))
}
