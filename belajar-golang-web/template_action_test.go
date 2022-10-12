package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type PageIf struct {
	Name  string
	Title string
}

type Value struct {
	Title      string
	FinalValue int
}

type HobbiesData struct {
	Title   string
	Hobbies []string
}

func TemplateIfAction(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	// Name bisa dihilangkan
	t.ExecuteTemplate(w, "if.gohtml", PageIf{
		Title: "Template Data If Action",
		Name:  "Dhany",
	})
}

func TestTemplateIfAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateIfAction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TemplateComparatorAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	// Name bisa dihilangkan
	t.ExecuteTemplate(w, "comparator.gohtml", Value{
		Title:      "Template Data Comparator Action",
		FinalValue: 60,
	})
}

func TestTemplateComparatorAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateComparatorAction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TemplateRangeAction(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	// Name bisa dihilangkan
	t.ExecuteTemplate(w, "range.gohtml", HobbiesData{
		Title: "Template Data Range Action",
		Hobbies: []string{
			"Memancing", "Mencuri",
		},
	})
}

func TestTemplateRangeAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateRangeAction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func TemplateWithAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	t.ExecuteTemplate(w, "address.gohtml", map[string]interface{}{
		"Title": "Template Data With Action",
		"Name":  "Dhany",
		"AddressData": map[string]interface{}{
			"Street": "Jalan Kebangsaan",
			"City":   "Jakarta Selatan",
		},
	})
}

func TestTemplateWithAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateWithAction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
