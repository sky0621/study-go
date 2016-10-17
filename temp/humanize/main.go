package main

import (
	"log"
	"os"
	"time"

	humanize "github.com/dustin/go-humanize"
)

func main() {
	name := os.Args[1]
	s, _ := os.Stat(name)
	log.Printf("%s: %s\n",
		humanize.Time(time.Now()),
		humanize.Bytes(uint64(s.Size())))
}
