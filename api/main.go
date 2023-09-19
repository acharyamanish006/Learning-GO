package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/hey", basicHandler)
	router.Post("/post", postFunc)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server err", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey"))
	fmt.Println("hey")
}

func postFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "ok")
	// fmt.Println("post")

	// res.Write([]byte("post"))
	// res.Write([]byte("post"))
}
