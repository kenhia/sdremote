package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
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

func startUrlInBrowser(browserPath, url, tag string, w http.ResponseWriter) {
	cmd := exec.Command(browserPath, url)
	err := cmd.Start()
	if err != nil {
		errorCouldNotBrowseToURL(tag, w, err)
		return
	}
	respondStatusOK(w)
}

func getUrlAndBrowserExeFromQuery(r *http.Request, browserDefault string) (url string, browserExe string) {
	query := r.URL.Query()
	url = query.Get("url")
	browser := query.Get("browser")
	if browser == "" {
		browser = browserDefault
	}
	browserExe = browserNameToExe(browser)
	return
}

func handleBrowseGet(w http.ResponseWriter, r *http.Request) {
	tag := "Browse GET"
	url, exe := getUrlAndBrowserExeFromQuery(r, "")
	if url == "" {
		errorNoURLSupplied(tag, w)
		return
	}
	startUrlInBrowser(exe, url, tag, w)
}

func handleEdge(w http.ResponseWriter, r *http.Request) {
	tag := "Edge GET"
	url, exe := getUrlAndBrowserExeFromQuery(r, "edge")
	if url == "" {
		errorNoURLSupplied(tag, w)
		return
	}
	startUrlInBrowser(exe, url, tag, w)
}

func handleChrome(w http.ResponseWriter, r *http.Request) {
	tag := "Chrome GET"
	url, exe := getUrlAndBrowserExeFromQuery(r, "chrome")
	if url == "" {
		errorNoURLSupplied(tag, w)
		return
	}
	startUrlInBrowser(exe, url, tag, w)
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

	exe := browserNameToExe(browseRequest.Browser)
	startUrlInBrowser(exe, browseRequest.Url, tag, w)
}
