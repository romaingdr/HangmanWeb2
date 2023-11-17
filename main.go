package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func main() {

	temp, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur => %s", err.Error()))
	}

	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "accueil", nil)
	})

	http.HandleFunc("/rules", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "rules", nil)
	})

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "game", nil)
	})

	// Gestion des fichiers dans assets
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	// Serveur
	fmt.Println("Serveur lancé sur : http://localhost:8080/accueil")
	http.ListenAndServe("localhost:8080", nil)
}
