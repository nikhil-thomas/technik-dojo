package main

import (
	"flag"
	"log"

	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/platform/dblayer"

	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/cmd/myevents/config"
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/cmd/myevents/restapi/events"
)

func main() {
	confPath := flag.String("conf", "", "path to config file")
	flag.Parse()

	config, err := config.ExtractConfiguration(*confPath)
	if err != nil {
		log.Fatalln("error in configuraing app:", err)
	}
	store, err := dblayer.NewStore(config.DatabaseType, config.DBConnectionAddr)
	if err != nil {
		log.Fatalln("error in creating store:", err)
	}
	log.Println("store created successfully")
	httpErrChan, httpsErrChan := events.ServeAPI(config.RestfulEndPoint, config.RestfulTLSEndPoint, store)
	select {
	case err := <-httpErrChan:
		log.Fatalln("http error:", err)
	case err := <-httpsErrChan:
		log.Fatalln("https error:", err)
	}
}
