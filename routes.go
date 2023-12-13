package main

import (
	"crud/books"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO --> make it all /books and take methods as per called

func loadRoutes() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/books", books.Index)
	mux.GET("/books/show", books.Show)
	mux.GET("/books/create", books.Create)
	mux.POST("/books/create/process", books.CreateProcess)
	mux.GET("/books/update", books.Update)
	mux.POST("/books/update/process", books.UpdateProcess)
	mux.GET("/books/delete/process", books.DeleteProcess)
	http.ListenAndServe(":8080", mux)

}
