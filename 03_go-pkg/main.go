package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "log"
)

type T struct {
    F int "a,omitempty"
    B int
}

func main() {
	t := &T{F: 1}
	out, err := yaml.Marshal(t)
	if err != nil {
            log.Fatalf("cannot marshal %#v: %v", t, err)
	}
	fmt.Printf("we got %q\n", out)
}

