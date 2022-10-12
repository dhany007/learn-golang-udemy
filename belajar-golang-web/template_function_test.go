package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", my name is " + myPage.Name
}

func TemplateFunctionString(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Kalai"}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Dhany",
	})
}

func TestTemplateFunctionString(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionString(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Kalai"}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Dhany",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.ParseFS(templatesFunction, "templates/*.gohtml"))
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Dhany",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

// menambah function global
func TemplateFunctionMap(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Dhany",
	})
}
func TestTemplateFunctionMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

// pipelines => hasil function dikirim ke function berikutnya
// memakai |
func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Dhany Putra Aritonang",
	})
}
func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
