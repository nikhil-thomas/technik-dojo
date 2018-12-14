package main

import (
    "fmt"
    "os"
//    "github.com/spf13/viper"
)

func main() {
    fmt.Println("viper test")
    fmt.Println("V_VAL_1:", os.Getenv("V_VAL_1"))
    //viper.SetConfigName("config")
    //viper.AddConfigPath(".")
    //err := viper.ReadInConfig()
    //if err != nil { // Handle errors reading the config file
//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
  //  }
    //fmt.Println("V_VAL_2:", viper.GetString("V_VAL_2"))
}

