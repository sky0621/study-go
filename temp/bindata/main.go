package main

import "log"

//go:generate go-bindata data/

func main() {
	_, err := Asset("data/logo.png")
	if err != nil {
		log.Fatal(err)
	}

}
