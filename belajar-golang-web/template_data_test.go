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
var templatesData embed.FS

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templatesData, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Dhany",
		"Address": map[string]interface{}{
			"Street": "Belom ada jalan map",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

type Address struct {
	Street string
}

type Page struct {
	Name    string
	Title   string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templatesData, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Dhany",
		Address: Address{
			Street: "Jalan Belom ada",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
