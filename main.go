package main

import (
	"gopher/config"
	"gopher/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":" + config.PORT

	app := routes.Init()

	router := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.ServeHTTP(w, r)
	})

	s := &http.Server{
		Handler: router,
		Addr:    addr,
	}

	log.Println("INFO Starting server on", os.Getenv("PORT"))

	log.Fatalf("ERROR %+v", s.ListenAndServe())
}
