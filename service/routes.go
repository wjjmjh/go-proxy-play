package service

import "github.com/lab/go-proxy/models"

var ReverseToFlaskRoute = models.Route{
	Name:         "ReverseToGitHub",   // Name
	Method:       "GET",               // HTTP method
	Pattern:      "/reverse-to-flask", // Route pattern
	HandlerFunc:  ReverseToFlask,
	Authenticate: true,
}

var ReverseToFlaskSvc0Route = models.Route{
	Name:         "ReverseToGitHub",         // Name
	Method:       "GET",                     // HTTP method
	Pattern:      "/reverse-to-flask-svc-0", // Route pattern
	HandlerFunc:  ReverseToFlaskSvc0,
	Authenticate: false,
}

var ReverseToFlaskSvc1GetRoute = models.Route{
	Name:         "ReverseToGitHub",             // Name
	Method:       "GET",                         // HTTP method
	Pattern:      "/reverse-to-flask-svc-1-get", // Route pattern
	HandlerFunc:  ReverseToFlaskSvc1Get,
	Authenticate: false,
}

var ReverseToFlaskSvc1PostRoute = models.Route{
	Name:         "ReverseToGitHub",              // Name
	Method:       "POST,OPTIONS",                 // HTTP method
	Pattern:      "/reverse-to-flask-svc-1-post", // Route pattern
	HandlerFunc:  ReverseToFlaskSvc1Post,
	Authenticate: false,
}
