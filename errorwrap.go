package main

import (
	"fmt"
	"log"
	"net/http"
)

func errorReadingJsonBody(tag string, w http.ResponseWriter, err error) {
	msg := fmt.Sprintf("%s: Error reading JSON body - %v", tag, err)
	http.Error(w, msg, http.StatusBadRequest)
	log.Println(msg)
}

func errorClosingRequestBody(tag string, w http.ResponseWriter, err error) {
	msg := fmt.Sprintf("%s: Error closing body - %v\n", tag, err)
	http.Error(w, msg, http.StatusInternalServerError)
	log.Println(msg)
}

func errorUnmarshalJson(tag string, w http.ResponseWriter, err error) {
	msg := fmt.Sprintf("%s: Could not unmarshal JSON body - %v", tag, err)
	http.Error(w, msg, http.StatusUnprocessableEntity)
	log.Println(msg)
}

func errorNoURLSupplied(tag string, w http.ResponseWriter) {
	msg := "No URL supplied."
	http.Error(w, msg, http.StatusBadRequest)
	log.Println(msg)
}

func errorCouldNotBrowseToURL(tag string, w http.ResponseWriter, err error) {
	msg := fmt.Sprintf("%s: Could not start url - %v", tag, err)
	http.Error(w, msg, http.StatusInternalServerError)
	log.Println(msg)
}

func errorCouldNotStartProgram(tag string, w http.ResponseWriter, err error) {
	msg := fmt.Sprintf("%s: Could not start program - %v", tag, err)
	http.Error(w, msg, http.StatusInternalServerError)
	log.Println(msg)
}

func errorBadKonsoleParam(tag, param string, w http.ResponseWriter) {
	msg := fmt.Sprintf("%s: Bad param '%s'.", tag, param)
	http.Error(w, msg, http.StatusBadRequest)
	log.Println(msg)
}

func respondStatusOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
