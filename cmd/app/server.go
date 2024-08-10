package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	funcMap = template.FuncMap{
		"iterate": func(start, end int) []int {
			var result []int
			for i := start; i < end; i++ {
				result = append(result, i)
			}
			return result
		},
	}
)

func (app *application) RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getDashboard)

	server := http.Server{
		Addr:    ":3030",
		Handler: mux,
	}

	err := server.ListenAndServe()
	log.Println("Server started at: ", &server.Addr)

	if err != nil {
		log.Println("Error while running the server: ", err.Error())
	}

}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("dashboard.tmpl").Funcs(funcMap).ParseGlob("./templates/dashboard.tmpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
