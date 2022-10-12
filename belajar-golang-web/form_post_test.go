package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	first_name := r.PostForm.Get("first_name")
	last_name := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hallo %s %s", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	// membuat body nya
	requsetBody := strings.NewReader("first_name=Dhany&last_name=Aritonang")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requsetBody)

	// harus dibuat content type nya
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
