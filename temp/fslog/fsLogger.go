package fslog

import (
	"io"
	"log"
	"os"
)

type FsLogger struct {
	FilePathStr string
}

func (l *FsLogger) Println(str ...interface{}) {
	logfile, err := os.OpenFile(l.FilePathStr, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannot open " + l.FilePathStr + ": " + err.Error())
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(str)
}
