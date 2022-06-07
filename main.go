package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func handleKonsole(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here!")
	exe := "/usr/bin/konsole"
	query := r.URL.Query()
	workdir := query.Get("workdir")
	if len(workdir) == 0 {
		workdir = os.Getenv("HOME")
	}
	cmd := exec.Command(exe, "--workdir", workdir)
	err := cmd.Start()
	if err != nil {
		log.Printf("Could not start '%s', Error: %q", exe, err)
	} else {
		log.Printf("Konsole: %s", workdir)
	}
}

func main() {
	port := 8027
	portStr := fmt.Sprintf(":%d", port)
	fmt.Println("sdremote - v0.0.1")
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/_api/v1/edge", handleEdge)
	http.HandleFunc("/_api/v1/chrome", handleChrome)
	http.HandleFunc("/_api/v1/konsole", handleKonsole)
	http.HandleFunc("/_api/v1/browse", handleBrowse)

	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(portStr, nil))
}
