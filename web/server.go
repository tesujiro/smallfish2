package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

func newServer() *server {
	return &server{
		router: http.NewServeMux(),
	}
}

func main() {
	s := newServer()
	s.routes()

	// Kick HTTP Server
	go func() {
		err := http.ListenAndServe(":80", s.router)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	// Kick HTTPS Server
	err := http.ListenAndServeTLS(":443", "ssl/server.crt", "ssl/server.key", s.router)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

func (s *server) routes() {
	s.router.HandleFunc("/greet", s.handleHello())
	//s.router.HandleFunc("/api/flyer", s.handleFlyer())
	s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	}
}
