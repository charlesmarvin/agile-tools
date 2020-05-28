package main

import (
	"log"
	"net/http"
	"os"

	"github.com/charlesmarvin/agile-tools/httpapi"
)

func main() {
	http.HandleFunc("/api/v1/boards", httpapi.CreateBoardHandler)
	http.HandleFunc("/api/v1/boards/", httpapi.GetBoardHandler)

	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
