package belajargolangweb

import (
	"embed"
	"fmt"
	"net/http"
	"testing"
)

func FileServe(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/not_found.html")
	}
}

func TestFileServe(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(FileServe),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk embed.FS

//go:embed resources/not_found.html
var resourceNotFound embed.FS

func FileServeEmbed(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestFileServeEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(FileServe),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
