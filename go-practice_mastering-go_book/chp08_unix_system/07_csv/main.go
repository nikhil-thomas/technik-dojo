package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Arafatk/glot"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("spcify a data file")
		return
	}
	file := os.Args[1]
	_, err := os.Stat(file)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	allRecords, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	xP := []float64{}
	yP := []float64{}

	for _, rec := range allRecords {
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		yP = append(yP, y)
	}
	points := [][]float64{}
	points = append(points, xP)
	points = append(points, yP)
	fmt.Println(points)

	dimensions := 2
	persist := true
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	plot.SetTitle("Using Glot with csv data")
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("YAxis")
	style := "circle"

	plot.AddPointGroup("Circle", style, points)
	plot.SavePlot("output_plot.png")

}
