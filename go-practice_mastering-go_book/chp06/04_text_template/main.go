package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type entry struct {
	Number int
	Square int
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("specify a template file")
		fmt.Println("Usage:", filepath.Base(args[0]), "template_path")
		return
	}

	templF := args[1]
	data := [][]int{
		{-1, 1},
		{-2, 4},
		{-3, 9},
		{-4, 16},
	}

	var entries []entry

	for _, v := range data {
		d := entry{
			Number: v[0],
			Square: v[1],
		}
		entries = append(entries, d)
	}

	templ := template.Must(template.ParseGlob(templF))
	templ.Execute(os.Stdout, entries)
}
