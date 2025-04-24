package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	FLAG = "cyn0x{0mn1_w3b_tr4c3_r3m41n5_1n_th3_j4v4sh4d0w5_0f_th3_nu11v01d}"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/redirect.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/get-flag-e82h13h52n23", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Fprint(w, FLAG)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
