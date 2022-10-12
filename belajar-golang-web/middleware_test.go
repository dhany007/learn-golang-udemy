package belajargolangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r) // ini middleware nya
	fmt.Println("After Execute Handler")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler executed")
	fmt.Fprint(w, "Hallo Middleware")
}

func HelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Name executed")
	fmt.Fprint(w, "Hallo name")
}

func PanicTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Panic executed")
	panic("Opps panic")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")

	defer func() {
		err := recover()
		fmt.Println("Recover : ", err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Err : %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(w, r) // ini middleware nya
	fmt.Println("After Execute Handler")
}

func TestMiddlewareServer(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloWorld)
	mux.HandleFunc("/name", HelloName)
	// disini akan berhenti programnya
	// tapi kalau pake error handler, bakalan tidak berhenti
	mux.HandleFunc("/panic", PanicTest)

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
