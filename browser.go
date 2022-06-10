package main

import (
	"encoding/json"
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

func browserNameToExe(browser string) (exe string) {
	switch browser {
	case "chrome":
		exe = "/usr/bin/google-chrome"
	default:
		exe = "/usr/bin/microsoft-edge-dev"
	}
	return
}

func startUrlInBrowser(browserPath, url string) (err error) {
	cmd := exec.Command(browserPath, url)
	err = cmd.Start()
	if err != nil {
		log.Printf("Browse: ERROR - could not start '%s %s'\n", browserPath, url)
	} else {
		log.Printf("Browse: %s %s\n", browserPath, url)
	}
	return
}

func handleBrowsePost(w http.ResponseWriter, r *http.Request) {
	tag := "Browse POST"
	var browseRequest BrowseRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		errorReadingJsonBody(tag, w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		errorClosingRequestBody(tag, w, err)
		return
	}
	if err := json.Unmarshal(body, &browseRequest); err != nil {
		errorUnmarshalJson(tag, w, err)
		return
	}

	exePath := browserNameToExe(browseRequest.Browser)
	if err = startUrlInBrowser(exePath, browseRequest.Url); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	respondStatusOK(w)
}

func handleBrowseGet(w http.ResponseWriter, r *http.Request) {
	tag := "Browse GET"
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		errorNoURLSupplied(tag, w)
		return
	}
	browser := query.Get("browser")
	exe := browserNameToExe(browser)

	if err := startUrlInBrowser(exe, url); err != nil {
		errorCouldNotBrowseToURL(tag, w, err)
		return
	}
	respondStatusOK(w)
}

func handleEdge(w http.ResponseWriter, r *http.Request) {
	tag := "Edge GET"
	exe := browserNameToExe("edge")
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		errorNoURLSupplied(tag, w)
		return
	}

	if err := startUrlInBrowser(exe, url); err != nil {
		errorCouldNotBrowseToURL(tag, w, err)
		return
	}
	respondStatusOK(w)
}

func handleChrome(w http.ResponseWriter, r *http.Request) {
	tag := "Chrome GET"
	exe := browserNameToExe("chrome")
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		errorNoURLSupplied(tag, w)
		return
	}

	if err := startUrlInBrowser(exe, url); err != nil {
		errorCouldNotBrowseToURL(tag, w, err)
		return
	}
	respondStatusOK(w)
}
