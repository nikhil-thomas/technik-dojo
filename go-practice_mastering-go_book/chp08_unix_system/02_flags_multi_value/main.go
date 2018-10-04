package main

import (
	"flag"
	"fmt"
	"strings"
)

type namesFlag struct {
	Names []string
}

// namesFlag should to implement flag.Value interface

func (s *namesFlag) GetNames() []string {
	return s.Names
}

func (s *namesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *namesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames namesFlag

	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "John", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)

	for i, name := range manyNames.GetNames() {
		fmt.Println(i, name)
	}

	fmt.Println("Remaining arguments")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
