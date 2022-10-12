package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()

	directory, _ := fs.Sub(resources, "resources") // masuk ke folder resources

	// filepath nya memang udah ada, coba klik aja servefiles itu
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	assert.Equal(t, "Hallo", string(bytes))
}

func TestServeFileGoodbye(t *testing.T) {
	router := httprouter.New()

	directory, _ := fs.Sub(resources, "resources")

	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	assert.Equal(t, "goodbye boy", string(bytes))
}
