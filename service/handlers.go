package service

import (
	"net/http"
	"net/url"

	"github.com/lab/go-proxy/mapper"

	"github.com/lab/go-proxy/reverseproxy"
)

func ReverseToFlask(w http.ResponseWriter, r *http.Request) {
	path, err := url.Parse("http://0.0.0.0:5000/api")
	if err != nil {
		panic(err)
		return
	}
	proxy := reverseproxy.NewReverseProxy(path)
	proxy.ServeHTTP(w, r)
}

func ReverseToFlaskSvc0(w http.ResponseWriter, r *http.Request) {
	path, err := url.Parse("http://0.0.0.0:9000/flask-svc-0")
	if err != nil {
		panic(err)
		return
	}
	proxy := reverseproxy.NewReverseProxy(path)
	proxy.ServeHTTP(w, r)
}

func ReverseToFlaskSvc1Get(w http.ResponseWriter, r *http.Request) {
	path, err := url.Parse("http://0.0.0.0:9001/flask-svc-1-get")
	if err != nil {
		panic(err)
		return
	}
	proxy := reverseproxy.NewReverseProxy(path)
	proxy.ServeHTTP(w, r)
}

func ReverseToFlaskSvc1Post(w http.ResponseWriter, r *http.Request) {
	path, err := url.Parse(mapper.Mapper["reverse-to-flask-svc-1-post"])
	if err != nil {
		panic(err)
		return
	}
	proxy := reverseproxy.NewReverseProxy(path)

	proxy.ServeHTTP(w, r)
}
