package belajargolangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)

	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(32) // mengubah besar file, default 20
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources_upload/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

//go:embed resources_upload/tes.jpg
var uploadFileTest []byte

func TestUploadForm(t *testing.T) {
	body := new(bytes.Buffer)                                        // membuat body ke bytes buffer
	writer := multipart.NewWriter(body)                              // bantuan multipart
	writer.WriteField("name", "Dhany")                               // meamasukkan value name
	file, _ := writer.CreateFormFile("file", "contohfileupload.jpg") // memasukkan value file
	file.Write(uploadFileTest)                                       // memasukkan value file
	writer.Close()                                                   // tutup writer nya

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType()) // ngeset multipart itu

	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	bodyResponse, _ := io.ReadAll(response.Body)

	bodyString := string(bodyResponse)
	fmt.Println(bodyString)
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources_upload"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
