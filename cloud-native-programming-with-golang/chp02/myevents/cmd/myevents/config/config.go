package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/platform/dblayer"
)

var (
	// DBTypeDefault specifies the default db choice
	DBTypeDefault = dblayer.MONGODB

	// DBConnectionDefault specifies default db address
	DBConnectionDefault = "mongodb://127.0.0.1"

	// RestfulEPDefault specifies the default rest api end point
	RestfulEPDefault = ":8080"
)

// ServiceConfig describes service config
type ServiceConfig struct {
	DatabaseType     dblayer.DBTYPE `json:"databasetype"`
	DBConnectionAddr string         `json:"dbconnectionaddr"`
	RestfulEndPoint  string         `json:"restfulapi_endpoint"`
}

// ExtractConfiguration returns config
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
	}
	log.Println("filename:", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error in opening config file,", err)
		fmt.Println("using default configuration")
		return conf, nil
	}
	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
