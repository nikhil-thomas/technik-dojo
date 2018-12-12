package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type entry struct {
	Number int
	Double int
	Square int
}

var data []entry
var tmplF string

func hndlr(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s, Path: %s\n", r.Host, r.URL.Path)
	tmpl := template.Must(template.ParseGlob(tmplF))
	tmpl.Execute(w, data)
}

func main() {
	args := os.Args
	fmt.Println("start app")
	if len(args) != 3 {
		fmt.Printf("Usage\n%s db_filename template_filename\n\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	dbFile := args[1]
	tmplF = args[2]

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Emptying database table")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Populating", dbFile)
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values(?,?,?)")
	for i := 21; i <= 50; i++ {
		stmt.Exec(i, i*2, i*i)
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		log.Fatalln(err)
	}

	n, d, s := 0, 0, 0

	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		temp := entry{
			Number: n,
			Double: d,
			Square: s,
		}
		data = append(data, temp)
	}

	http.HandleFunc("/", hndlr)
	fmt.Println("Listening at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
