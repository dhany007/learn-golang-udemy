package belajargolangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// stripprefix => untuk menghilangkan /static pada url nya
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	// nanti akan menjadi /static/resources/index.js
	// resource itu sub directory, harus dihilangkan
	// kita masuk kedalam subdirectory nya
	directory, _ := fs.Sub(resources, "resources")
	// otomatis file akan di embed seperti course embed
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
