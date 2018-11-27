package main

import (
	"os"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/09_logrus/num"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	log.WithFields(log.Fields{
		"key1": "val1",
		"key2": "val2",
	}).Warn("this is a warning")
	num.Print()
}
