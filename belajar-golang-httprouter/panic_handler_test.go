package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// by default, panic akan menghentikan aplikasi kita
// maka dibuat panic handler dari httproute supaya ketika panic, tidak menghentikan aplikasinya
func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// ini adalah panic handler nya
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic ", i)
	}

	// disini kita akan mereturn kan panic
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	assert.Equal(t, "Panic Ups", string(bytes))
}
