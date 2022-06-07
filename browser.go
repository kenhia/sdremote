package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type BrowseRequest struct {
	Url     string `json:"url"`
	Browser string `json:"browser"`
}

func startUrlInBrowser(browserPath, url string) {
	cmd := exec.Command(browserPath, url)
	err := cmd.Start()
	if err != nil {
		log.Printf("Browse: ERROR - could not start '%s %s'\n", browserPath, url)
	} else {
		log.Printf("Browse: %s %s\n", browserPath, url)
	}
}

func handleBrowse(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleBrowseGet(w, r)
	case "POST":
		fmt.Println("in Browse POST")
		handleBrowsePost(w, r)
	}
}

func handleBrowsePost(w http.ResponseWriter, r *http.Request) {
	var browseRequest BrowseRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Printf("handleBrowsePost: Error reading body - %v\n", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Printf("handleBrowsePost: Error closing body - %v\n", err)
		return
	}
	if err := json.Unmarshal(body, &browseRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf("handleBrowsePost: Error return error (yep!) - %v\n", err)
		}
	}
	exePath := browserNameToExe(browseRequest.Browser)
	startUrlInBrowser(exePath, browseRequest.Url)
}

func browserNameToExe(browser string) (exe string) {
	switch browser {
	case "chrome":
		exe = "/usr/bin/google-chrome"
	default:
		exe = "/usr/bin/microsoft-edge-dev"
	}
	return
}

func handleBrowseGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		log.Println("Browse: ERROR - no url supplied")
		http.Error(w, "400 Bad Client Request; no url supplied", http.StatusBadRequest)
		return
	}
	browser := query.Get("browser")
	exe := browserNameToExe(browser)

	startUrlInBrowser(exe, url)
}

func handleEdge(w http.ResponseWriter, r *http.Request) {
	exe := "/usr/bin/microsoft-edge-dev"
	query := r.URL.Query()
	url := query.Get("url")

	startUrlInBrowser(exe, url)
}

func handleChrome(w http.ResponseWriter, r *http.Request) {
	exe := "/usr/bin/google-chrome"
	query := r.URL.Query()
	url := query.Get("url")

	startUrlInBrowser(exe, url)
}
