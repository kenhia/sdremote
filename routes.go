package main

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func api_v1(subRoute string) string {
	return fmt.Sprintf("/_api/v1/%s", subRoute)
}

var routes = Routes{
	Route{"Index", "GET", "/", handleRequest},
	Route{"Konsole", "GET", api_v1("konsole"), handleKonsoleGet},
	Route{"Konsole", "POST", api_v1("konsole"), handleKonsolePost},
	Route{"Edge", "GET", api_v1("edge"), handleEdge},
	Route{"Chrome", "GET", api_v1("chrome"), handleChrome},
	Route{"Browse", "GET", api_v1("browse"), handleBrowseGet},
	Route{"Browse", "POST", api_v1("browse"), handleBrowsePost},
}
