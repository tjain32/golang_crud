package main

import (
	"log"
	"net/http"

	"crud/config"

	"github.com/julienschmidt/httprouter"
)

func main() {
	loadRoutes()
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
