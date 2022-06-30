package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func cleanup() {
	// Nothing to do for now
}

var (
	port int
)

func handleCommandLine() {
	flag.IntVar(&port, "port", 8027, "port to use (default: 8027)")
	flag.Parse()
}

func main() {
	handleCommandLine()

	portStr := fmt.Sprintf(":%d", port)
	fmt.Println("sdremote - v0.0.1")

	router := NewRouter()

	// Handle CTRL-C as method for terminating the server nicely
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		fmt.Println("\n\nCTRL-C pressed; stopping now. bye!")
		os.Exit(0)
	}()

	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(portStr, router))
}
