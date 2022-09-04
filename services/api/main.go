package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	listenAddr := fmt.Sprintf(":%s", port)

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello!")
		})

		log.Printf("Listening a server on port: %s\n", port)
		if err := http.ListenAndServe(listenAddr, mux); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Server is ended!")
}
