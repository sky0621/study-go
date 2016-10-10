package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		fmt.Println(i)
		switch {
		case i < 1:
			fmt.Println("i < 1")
		case i < 3:
			fmt.Println("i < 3")
		case i < 5:
			fmt.Println("i < 5")
		case i < 10:
			fmt.Println("i < 10")
		default:
			fmt.Println("i >= 10")
		}

	}
}
