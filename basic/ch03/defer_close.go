package main

import "os"
import "fmt"

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Abnormal End!")
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("aiueo\n")
}
