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

func startUrlInBrowser(browserPath, url string) error {
	cmd := exec.Command(browserPath, url)
	err := cmd.Start()
	return err
}

func handleEdge(w http.ResponseWriter, r *http.Request) {
	exe := "/usr/bin/microsoft-edge-dev"
	query := r.URL.Query()
	url := query.Get("url")
	cmd_str := fmt.Sprintf("%s %s", exe, url)

	err := startUrlInBrowser(exe, url)
	if err != nil {
		log.Printf("Could not start '%s', Error: %q", cmd_str, err)
	} else {
		log.Printf("Edge: %s", url)
	}
}

func handleChrome(w http.ResponseWriter, r *http.Request) {
	exe := "/usr/bin/google-chrome"
	query := r.URL.Query()
	url := query.Get("url")
	cmd_str := fmt.Sprintf("%s %s", exe, url)

	err := startUrlInBrowser(exe, url)
	if err != nil {
		log.Printf("Could not start '%s', Error: %q", cmd_str, err)
	} else {
		log.Printf("Chrome: %s", url)
	}
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

	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(portStr, nil))
}
