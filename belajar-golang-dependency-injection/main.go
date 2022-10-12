package main

import (
	"dhany007/belajar-golang-restful-api/helper"
	"dhany007/belajar-golang-restful-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMidleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMidleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()

	helper.PanicIfError(err)
}
