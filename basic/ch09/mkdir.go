package main

import "os"

func main(){
	os.MkdirAll("a/b/c", 0777)
}
