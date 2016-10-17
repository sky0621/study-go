package main

import (
	"crypto/sha256"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	tmp, _ := ioutil.TempFile(os.TempDir(), "tmp")
	defer tmp.Close()

	hash := sha256.New()

	w := io.MultiWriter(tmp, hash)

	written, _ := io.Copy(w, os.Stdin)
	log.Printf("Wrote %d bytes to %s\nSha256: %x\n",
		written,
		tmp.Name(),
		hash.Sum(nil))
}
