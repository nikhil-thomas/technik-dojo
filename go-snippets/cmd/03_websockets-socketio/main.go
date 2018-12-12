// Author : Nikhil Thomas
// Credit : https://tutorialedge.net/golang/golang-websockets-tutorial/
// Date   : September 05 2018

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatalln(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.Join("chat")

		// on "chat message event reply with same message and time"
		so.On("chat message", func(msg string) {
			log.Println("emit:", msg)
			now := time.Now()

			err := so.Emit("chat message", now.UTC().String()+": "+msg)
			if err != nil {
				log.Println(err)
			}

			// also broadcast same message and time to chat room"
			err = so.BroadcastTo("chat", "chat message", now.UTC().String()+": "+msg)
			if err != nil {
				log.Println(err)
			}
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	// negroni classic handles staic file serving
	//fs := http.FileServer(http.Dir("./public"))
	//http.Handle("/", fs)

	http.Handle("/socket.io/", server)

	// negroni classic provides
	// logging middleware
	// static file serving for dir ./public middleware
	// revcovery middleware
	n := negroni.Classic()
	n.UseHandler(http.DefaultServeMux)

	log.Println("listening on :3000")
	log.Fatalln(http.ListenAndServe(":3000", n))
}
