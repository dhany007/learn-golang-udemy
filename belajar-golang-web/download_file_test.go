package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")

	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BAD REQUEST")
		return
	}

	// menjadikan dia attacment, tidak dirender oleh browser
	w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	http.ServeFile(w, r, "./resources_upload/"+filename)
}

func TestDownloadFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
