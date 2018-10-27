package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"net/rpc"

	"github.com/nikhil-thomas/technik-dojo/go-practice_mastering-go_book/chp13_network_programming_2/08_rpc/sharedrpc"
)

// MathExecer implements sharedrpc.CalcExecer interface
type MathExecer struct{}

// Power computes x^y
func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

// Multiply computes product of inputs
func (c *MathExecer) Multiply(args *sharedrpc.CalcInput, result *float64) error {
	*result = args.A1 * args.A2
	return nil
}

// Power computes power of inputs
func (c *MathExecer) Power(args *sharedrpc.CalcInput, result *float64) error {
	*result = Power(args.A1, args.A2)
	return nil
}

func main() {
	listenOn := ":8000"

	calcHandle := new(MathExecer)

	rpc.Register(calcHandle)

	t, err := net.ResolveTCPAddr("tcp4", listenOn)
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.ListenTCP("tcp4", t)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("remote addr:", conn.RemoteAddr())
		rpc.ServeConn(conn)
	}
}
