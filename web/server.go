package main

import (
	"encoding/json"
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
		fmt.Println("ListenAndServe :80")
		err := http.ListenAndServe(":80", s.router)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	// Kick HTTPS Server
	fmt.Println("ListenAndServeTLS :443")
	err := http.ListenAndServeTLS(":443", "ssl/server.crt", "ssl/server.key", s.router)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

func (s *server) routes() {
	s.router.HandleFunc("/greet", s.handleHello())
	s.router.HandleFunc("/cities", s.handleCities())
	//s.router.HandleFunc("/api/flyer", s.handleFlyer())
	s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handlerHello")
		fmt.Fprintf(w, "Hello, World")
	}
}

type jsmap map[string]interface{}

func (s *server) handleCities() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handlerCities")
		w.Header().Set("Content-Type", "application/json")
		list := []jsmap{
			{"name": "Setagaya", "population": 925226},
			{"name": "Nerima", "population": 733150},
			{"name": "Ota", "population": 731273},
			{"name": "Edogawa", "population": 691417},
			{"name": "Adachi", "population": 678686},
			{"name": "Suginami", "population": 577903},
		}
		m := jsmap{"list": list}
		j, err := json.Marshal(m)
		if err != nil {
			fmt.Printf("json.Marshall error:%v\n", err)
		}
		fmt.Fprintf(w, string(j))
	}
}
