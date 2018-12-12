package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users", allUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{name}", deleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/users/{name}/{email}", updateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{name}/{email}", newUser).Methods(http.MethodPost)

	n := negroni.New(negroni.NewLogger())
	n.UseHandler(r)
	log.Println("Listening at :3000")
	log.Fatalln(http.ListenAndServe(":3000", n))
}

func main() {
	initialMigration()
	handleRequests()
}
