package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main(){
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	w := csv.NewWriter(file)
	w.Write([]string{"No.", "Item", "Amount", "Note"})
	w.Write([]string{"1", "Carot", "150", "fresh!"})
	w.Write([]string{"2", "Onion", "100", "taste good."})
	w.Write([]string{"3", "Tomato", "99", "cheap"})

	w.Flush()
}
