package main

import "fmt"

func main() {
	pgs := []string{"Go", "Java", "Scala", "Ruby", "PHP", "Python", "JavaScript"}
	for _, pg := range pgs {
		fmt.Println(pg)
	}
	fmt.Println("############################")
	fmt.Print("pgs[1]: ")
	fmt.Println(pgs[1])
	fmt.Print("pgs[2:]: ")
	fmt.Println(pgs[2:])
	fmt.Print("pgs[3:5]: ")
	fmt.Println(pgs[3:5])
	fmt.Print("pgs[:4]: ")
	fmt.Println(pgs[:4])
}
