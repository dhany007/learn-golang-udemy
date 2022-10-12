package belajargolangweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templatesCaching embed.FS

// cache nya di deklarasikan secara global
var myTemplates = template.Must(template.ParseFS(templatesCaching, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {

	myTemplates.ExecuteTemplate(w, "simple.gohtml", "HELLO TEMPLATE CACHING")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
