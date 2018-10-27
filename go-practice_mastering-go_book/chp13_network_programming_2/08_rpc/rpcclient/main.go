package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/nikhil-thomas/technik-dojo/go-practice_mastering-go_book/chp13_network_programming_2/08_rpc/sharedrpc"
)

func main() {
	connectTo := ":8000"

	conn, err := rpc.Dial("tcp", connectTo)
	if err != nil {
		log.Fatalln(err)
	}
	args := sharedrpc.CalcInput{16, 0.5}
	var reply float64

	err = conn.Call("MathExecer.Multiply", args, &reply)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("reply (multiply): %f\n", reply)

	err = conn.Call("MathExecer.Power", args, &reply)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("reply (power): %f\n", reply)
}
