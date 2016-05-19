package main

import "fmt"

func main(){

	var date [7]string
	date[0] = "日"
	date[1] = "月"
	date[2] = "火"
	date[3] = "水"
	date[4] = "木"
	date[5] = "金"
	date[6] = "土"

	for i := 0; i < len(date); i++ {
		fmt.Print(date[i], " ")
	}
	fmt.Println()

	for _, value := range date {
		fmt.Print(value, " ")
	}
	fmt.Println()

}
