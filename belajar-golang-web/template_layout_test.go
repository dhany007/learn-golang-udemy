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

var templatesLayout embed.FS

func TemplateLayoutAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templatesLayout, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title": "Template Data Layout Action",
		"Name":  "Dhany",
	})
}

func TestTemplateLayoutAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutAction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TemplateLayoutNameAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templatesLayout, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Data Layout Action",
		"Name":  "Dhany",
	})
}

func TestTemplateLayoutNameAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutNameAction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
