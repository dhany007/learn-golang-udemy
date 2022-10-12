package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		fmt.Fprintf(w, "Hallo %s", name)
	}
}

func TestResponseCodeValid(t *testing.T) {
	// membuat body nya
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/?name=dhany", nil)

	// harus dibuat content type nya
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TestResponseCodeInvalid(t *testing.T) {
	// membuat body nya
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/?name", nil)

	// harus dibuat content type nya
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	bodyString := string(body)
	fmt.Println(bodyString)
}
