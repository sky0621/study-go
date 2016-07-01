package main

import (
	"fmt"
	"strconv"
)

func main() {
	i64, err := strconv.ParseInt("1234567890", 10, 0)
	fmt.Println(int(i64), err)
}
