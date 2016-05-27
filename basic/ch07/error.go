package main

import (
	"fmt"
	"strconv"
)

type MyError struct {
	message string
}

func (err MyError) Error() string {
	return err.message
}

func main() {
	val, err := strconv.Atoi("1")
	fmt.Println(val, err)

	val, err = strconv.Atoi("abcde")
	fmt.Println(val, err)

	val, err = strconv.Atoi("z")
	fmt.Println(val, err)
}
