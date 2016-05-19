package main

import "fmt"

func main(){

	capitals := make(map[string]string)

	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントン"
	capitals["中国"] = "北京"

	fmt.Println(capitals)
	fmt.Println()

	fmt.Println(capitals["日本"])
	fmt.Println()

	delete(capitals, "中国")
	fmt.Println(capitals)

}
