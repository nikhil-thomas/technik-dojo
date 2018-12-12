package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, programName)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Panic(sysLog)
	fmt.Println("will you see this")
}
