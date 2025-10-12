package main

import (
	"log"
	"net/http"

	"github.com/cobrich/url-shortener/handler"
)

func main() {
	port := ":8080"

	handler := handler.NewHandler()

	router := http.DefaultServeMux
	router.HandleFunc("/{short_code}", handler.GetLongURLHundler)
	router.HandleFunc("/shorten", handler.CreateShortURLHundler)

	log.Print("Server running on localhost port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}
