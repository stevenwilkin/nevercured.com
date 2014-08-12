package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
)

var (
	port int
	tmpl *template.Template
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()

	pwd, _ := os.Getwd()
	filename := path.Join(pwd, "templates", "index.tmpl")
	tmpl, _ = template.ParseFiles(filename)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, struct{}{})
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting!")
	}
}
