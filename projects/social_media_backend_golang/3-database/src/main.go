package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", testHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
