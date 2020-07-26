package main

import (
	"log"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.UserAgent())
	writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	http.ServeFile(writer, request, "./index.html")
}

func main() {
	server := &http.Server{
		Handler: http.HandlerFunc(Index),
		Addr:    "127.0.0.1:9000",
	}
	log.Fatal(server.ListenAndServe())
}
