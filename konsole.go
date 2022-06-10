package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func isBlessedKonsoleParam(parameter string) bool {
	blessedSwitches := []string{
		"workdir",
		"hold",
		"no-close",
		"new-tab",
		"tabs-from-file",
		"background-mode",
		"show-menubar",
		"hide-menubar",
		"show-tabbar",
		"hide-tabbar",
		"fullscreen",
		"notransparency",
	}
	for _, item := range blessedSwitches {
		if item == parameter {
			return true
		}
	}
	return false
}

func launchKonsole(tag string, w http.ResponseWriter, args []string) {
	exe := "/usr/bin/konsole"

	cmd := exec.Command(exe)
	for _, v := range args {
		cmd.Args = append(cmd.Args, v)
	}
	if err := cmd.Start(); err != nil {
		errorCouldNotStartProgram(tag, w, err)
		return
	}
	respondStatusOK(w)

}

func handleKonsoleGet(w http.ResponseWriter, r *http.Request) {
	tag := "Konsole GET"
	query := r.URL.Query()
	workdir := query.Get("workdir")
	if len(workdir) == 0 {
		workdir = os.Getenv("HOME")
	}
	args := []string{"--workdir", workdir}

	launchKonsole(tag, w, args)
}

func handleKonsolePost(w http.ResponseWriter, r *http.Request) {
	tag := "Konsole POST"
	var konsoleArgs map[string]string

	var command_args []string

	body, err := ioutil.ReadAll((io.LimitReader(r.Body, 1048576)))
	if err != nil {
		msg := fmt.Sprintf("handleKonsolePost: Error reading body - %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		log.Println(msg)
		return
	}
	if err := r.Body.Close(); err != nil {
		errorReadingJsonBody(tag, w, err)
		return
	}
	if err := json.Unmarshal(body, &konsoleArgs); err != nil {
		errorUnmarshalJson(tag, w, err)
		return
	}
	fmt.Printf("- konsoleArgs - %v\n", konsoleArgs)
	for param, arg := range konsoleArgs {
		if !isBlessedKonsoleParam(param) {
			errorBadKonsoleParam(tag, param, w)
			return
		}
		command_args = append(command_args, fmt.Sprintf("--%s", param))
		if arg != "" {
			command_args = append(command_args, arg)
		}
	}
	launchKonsole("POST", w, command_args)
}
