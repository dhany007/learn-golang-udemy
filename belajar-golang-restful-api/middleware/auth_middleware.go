package middleware

import (
	"dhany007/belajar-golang-restful-api/helper"
	"dhany007/belajar-golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}

}

// memakai middleware ini, maka semua request yang masuk akan di-check
// memakai pointer
func (authMiddleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// cek token
	token := "tokenrahasia"
	if token == r.Header.Get("api_key") {
		// ok
		authMiddleware.Handler.ServeHTTP(w, r)
	} else {
		// err
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
