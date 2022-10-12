package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(rw http.ResponseWriter, _ *http.Request) {
		// logic web
		fmt.Fprint(rw, "Hello World")
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestHandlerServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(rw, "Hello World")
	})

	mux.HandleFunc("/hi", func(rw http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(rw, "hi...")
	})

	// masuk ke prioritas paling panjang
	mux.HandleFunc("/hi/name", func(rw http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(rw, "hi nama")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(rw, r.RequestURI)
		fmt.Fprint(rw, r.Method)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
